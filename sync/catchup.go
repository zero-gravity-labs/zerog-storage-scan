package sync

import (
	"context"
	"strings"
	"time"

	nhContract "github.com/0glabs/0g-storage-scan/contract"
	"github.com/0glabs/0g-storage-scan/rpc"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/health"
	viperUtil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type CatchupSyncer struct {
	conf           *SyncConfig
	sdk            *web3go.Client
	db             *store.MysqlStore
	currentBlock   uint64
	finalizedBlock uint64

	flowAddr      string
	flowSubmitSig string
	rewardAddr    string
	rewardSig     string

	daEntranceAddr    string
	dataUploadSig     string
	commitVerifiedSig string
	daRewardSig       string
	daSignersAddr     string
	newSignerSig      string
	socketUpdatedSig  string

	addresses     []common.Address
	topics        [][]common.Hash
	alertChannel  string
	healthReport  health.TimedCounterConfig
	nodeRpcHealth health.TimedCounter
}

func MustNewCatchupSyncer(sdk *web3go.Client, db *store.MysqlStore, conf SyncConfig, alertChannel string,
	healthReport health.TimedCounterConfig) *CatchupSyncer {
	var flow struct {
		Address              string
		SubmitEventSignature string
	}
	viperUtil.MustUnmarshalKey("flow", &flow)

	var reward struct {
		Address              string
		RewardEventSignature string
	}
	viperUtil.MustUnmarshalKey("reward", &reward)

	var daEntrance struct {
		Address                            string
		DataUploadSignature                string
		ErasureCommitmentVerifiedSignature string
		DARewardSignature                  string
	}
	viperUtil.MustUnmarshalKey("daEntrance", &daEntrance)

	var daSigners struct {
		Address                string
		NewSignerSignature     string
		SocketUpdatedSignature string
	}
	viperUtil.MustUnmarshalKey("daSigners", &daSigners)

	return &CatchupSyncer{
		conf: &conf,
		sdk:  sdk,
		db:   db,

		flowAddr:      flow.Address,
		flowSubmitSig: flow.SubmitEventSignature,
		rewardAddr:    reward.Address,
		rewardSig:     reward.RewardEventSignature,

		daEntranceAddr:    daEntrance.Address,
		dataUploadSig:     daEntrance.DataUploadSignature,
		commitVerifiedSig: daEntrance.ErasureCommitmentVerifiedSignature,
		daRewardSig:       daEntrance.DARewardSignature,
		daSignersAddr:     daSigners.Address,
		newSignerSig:      daSigners.NewSignerSignature,
		socketUpdatedSig:  daSigners.SocketUpdatedSignature,

		addresses: []common.Address{
			common.HexToAddress(flow.Address),
			common.HexToAddress(reward.Address),
			common.HexToAddress(daEntrance.Address),
			common.HexToAddress(daSigners.Address),
		},
		topics: [][]common.Hash{{
			common.HexToHash(flow.SubmitEventSignature),
			common.HexToHash(reward.RewardEventSignature),
			common.HexToHash(daEntrance.DataUploadSignature),
			common.HexToHash(daEntrance.ErasureCommitmentVerifiedSignature),
			common.HexToHash(daEntrance.DARewardSignature),
			common.HexToHash(daSigners.NewSignerSignature),
			common.HexToHash(daSigners.SocketUpdatedSignature),
		}},

		alertChannel:  alertChannel,
		healthReport:  healthReport,
		nodeRpcHealth: health.TimedCounter{},
	}
}

func (s *CatchupSyncer) Sync(ctx context.Context) {
	logrus.Info("Catchup syncer starting to sync data")
	for {
		needProcess := s.tryBlockRange(ctx)
		if !needProcess || interrupted(ctx) {
			return
		}

		if err := s.syncRange(ctx, s.currentBlock, s.finalizedBlock); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"currentBlock":   s.currentBlock,
				"finalizedBlock": s.finalizedBlock,
			}).Error("Catchup syncer sync range")
			time.Sleep(time.Second)
			continue
		}
	}
}

func (s *CatchupSyncer) syncRange(ctx context.Context, rangeStart, rangeEnd uint64) error {
	start, end := s.nextSyncRange(rangeStart, rangeEnd)

	for {
		var logs []types.Log
		var err error
		logs, _, end, err = s.batchGetLogsBestEffort(s.sdk, start, end, s.addresses, s.topics)
		if err != nil {
			return err
		}

		var bn2TimeMap map[uint64]uint64
		if len(logs) > 0 {
			blockNums := make([]types.BlockNumber, 0)
			for _, log := range logs {
				blockNums = append(blockNums, types.BlockNumber(log.BlockNumber))
			}
			bn2TimeMap, err = rpc.BatchGetBlockTimes(ctx, s.sdk, blockNums, s.conf.BatchBlocksOnBatchCall)
		}
		if err != nil {
			return err
		}

		rangeEndBlock, err := s.sdk.Eth.BlockByNumber(types.BlockNumber(end), false)
		if err != nil {
			return err
		}

		block := store.NewBlock(rangeEndBlock)
		decodedLogs, err := s.convertLogs(logs, bn2TimeMap)
		if err != nil {
			return err
		}

		err = s.db.Push(block, decodedLogs)
		if err != nil {
			return err
		}
		logrus.WithField("block", block.BlockNumber).Info("Catchup")

		if end >= rangeEnd || interrupted(ctx) {
			break
		}

		start, end = s.nextSyncRange(end+1, rangeEnd)
	}

	return nil
}

func (s *CatchupSyncer) nextSyncRange(curStart, rangeEnd uint64) (uint64, uint64) {
	start, end := curStart, rangeEnd

	if s.conf.BatchBlocksOnCatchup > 0 {
		end = start + s.conf.BatchBlocksOnCatchup - 1
		if end > rangeEnd {
			end = rangeEnd
		}
	}

	return start, end
}

// queryFlowSubmitsBestEffort queries flow submits from the contract event logs between a block range of (bnFrom, bnTo).
// It returns the logs, the actual start and end block numbers of the queried range, and an error if any.
// It may not return all the data for the whole range, but only the available or feasible data.
// It uses a binary search algorithm to find the optimal range that maximizes the number of logs returned.
func (s *CatchupSyncer) batchGetLogsBestEffort(w3c *web3go.Client, bnFrom, bnTo uint64,
	addresses []common.Address, topics [][]common.Hash) ([]types.Log, uint64, uint64, error) {
	start, end := bnFrom, bnTo

	for {
		logs, err := rpc.BatchGetLogs(w3c, start, end, addresses, topics)
		if err == nil {
			return logs, start, end, nil
		}

		if strings.Contains(err.Error(), "please narrow down your filter condition") || // espace
			strings.Contains(err.Error(), "block range") || // bsc
			strings.Contains(err.Error(), "blocks distance") || // evmos
			strings.Contains(err.Error(), "returned more than") { // kava
			end = start + (end-start)/2
			continue
		}

		return nil, 0, 0, err
	}
}

func (s *CatchupSyncer) tryBlockRange(ctx context.Context) bool {
	for try := 1; ; try++ {
		err := s.updateBlockRange(ctx)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"try":      try,
				"curBlk":   s.currentBlock,
				"finalBlk": s.finalizedBlock}).Info("try block range")
			time.Sleep(time.Second)
			continue
		}

		return s.currentBlock <= s.finalizedBlock
	}
}

func (s *CatchupSyncer) updateBlockRange(ctx context.Context) error {
	maxBlock, ok, err := s.db.MaxBlock()
	if err != nil {
		return errors.WithMessage(err, "failed to get max block from db")
	}
	if ok {
		s.currentBlock = maxBlock + 1
	} else {
		s.currentBlock = s.conf.BlockWhenFlowCreated
	}

	finalizedBlock, err := s.sdk.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if s.alertChannel != "" {
		if e := rpc.AlertErr(ctx, "BlockchainRPCError", s.alertChannel, err, s.healthReport,
			&s.nodeRpcHealth); e != nil {
			return e
		}
	}
	if err != nil {
		return errors.WithMessage(err, "failed to get finalized block")
	}

	s.finalizedBlock = finalizedBlock.Number.Uint64()

	return nil
}

func (s *CatchupSyncer) convertLogs(logs []types.Log, bn2TimeMap map[uint64]uint64) (*store.DecodedLogs, error) {
	var decodedLogs store.DecodedLogs

	for _, log := range logs {
		if log.Removed {
			continue
		}

		ts := bn2TimeMap[log.BlockNumber]
		blockTime := time.Unix(int64(ts), 0)

		switch log.Topics[0].String() {
		case s.flowSubmitSig:
			submit, err := s.decodeSubmit(blockTime, log)
			if err != nil {
				return nil, err
			}
			if submit != nil {
				decodedLogs.Submits = append(decodedLogs.Submits, *submit)
			}
		case s.rewardSig:
			reward, err := s.decodeReward(blockTime, log)
			if err != nil {
				return nil, err
			}
			if reward != nil {
				decodedLogs.Rewards = append(decodedLogs.Rewards, *reward)
			}
		case s.newSignerSig:
			signer, err := s.decodeNewSigner(blockTime, log)
			if err != nil {
				return nil, err
			}
			if signer != nil {
				decodedLogs.DASigners = append(decodedLogs.DASigners, *signer)
			}
		case s.socketUpdatedSig:
			signer, err := s.decodeSocketUpdated(blockTime, log)
			if err != nil {
				return nil, err
			}
			if signer != nil {
				decodedLogs.DASignersWithSocketUpdated = append(decodedLogs.DASignersWithSocketUpdated, *signer)
			}
		case s.dataUploadSig:
			daSubmit, err := s.decodeDataUpload(blockTime, log)
			if err != nil {
				return nil, err
			}
			if daSubmit != nil {
				decodedLogs.DASubmits = append(decodedLogs.DASubmits, *daSubmit)
			}
		case s.commitVerifiedSig:
			daSubmit, err := s.decodeCommitVerified(blockTime, log)
			if err != nil {
				return nil, err
			}
			if daSubmit != nil {
				decodedLogs.DASubmitsWithVerified = append(decodedLogs.DASubmitsWithVerified, *daSubmit)
			}
		case s.daRewardSig:
			daReward, err := s.decodeDAReward(blockTime, log)
			if err != nil {
				return nil, err
			}
			if daReward != nil {
				decodedLogs.DARewards = append(decodedLogs.DARewards, *daReward)
			}
		default:
			return nil, errors.Errorf("Faild to decode log, sig %v", log.Topics[0].String())
		}
	}

	return &decodedLogs, nil
}

func (s *CatchupSyncer) decodeSubmit(blkTime time.Time, log types.Log) (*store.Submit, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.flowAddr) || sig != s.flowSubmitSig {
		return nil, nil
	}

	submit, err := store.NewSubmit(blkTime, log, nhContract.DummyFlowFilterer())
	if err != nil {
		return nil, err
	}

	senderID, err := s.db.AddressStore.Add(submit.Sender, blkTime)
	if err != nil {
		return nil, err
	}

	submit.SenderID = senderID

	return submit, nil
}

func (s *CatchupSyncer) decodeReward(blkTime time.Time, log types.Log) (*store.Reward, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.rewardAddr) || sig != s.rewardSig {
		return nil, nil
	}

	reward, err := store.NewReward(blkTime, log, nhContract.DummyRewardFilterer())
	if err != nil {
		return nil, err
	}

	minerID, err := s.db.AddressStore.Add(reward.Miner, blkTime)
	if err != nil {
		return nil, err
	}

	_, err = s.db.MinerStore.Add(minerID, blkTime, reward.Amount)
	if err != nil {
		return nil, err
	}

	reward.MinerID = minerID

	return reward, nil
}

func (s *CatchupSyncer) decodeNewSigner(blkTime time.Time, log types.Log) (*store.DASigner, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daSignersAddr) || sig != s.newSignerSig {
		return nil, nil
	}

	signer, err := store.NewDASigner(blkTime, log, nhContract.DummyDASignersFilterer())
	if err != nil {
		return nil, err
	}

	signerID, err := s.db.AddressStore.Add(signer.Address, blkTime)
	if err != nil {
		return nil, err
	}

	signer.SignerID = signerID

	return signer, nil
}

func (s *CatchupSyncer) decodeSocketUpdated(blkTime time.Time, log types.Log) (*store.DASigner, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daSignersAddr) || sig != s.socketUpdatedSig {
		return nil, nil
	}

	signer, err := store.NewDASignerSocket(log, nhContract.DummyDASignersFilterer())
	if err != nil {
		return nil, err
	}

	signerID, err := s.db.AddressStore.Add(signer.Address, blkTime)
	if err != nil {
		return nil, err
	}

	signer.SignerID = signerID

	return signer, nil
}

func (s *CatchupSyncer) decodeDataUpload(blkTime time.Time, log types.Log) (*store.DASubmit, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daEntranceAddr) || sig != s.dataUploadSig {
		return nil, nil
	}

	daSubmit, err := store.NewDASubmit(blkTime, log, nhContract.DummyDAEntranceFilterer())
	if err != nil {
		return nil, err
	}

	senderID, err := s.db.AddressStore.Add(daSubmit.Sender, blkTime)
	if err != nil {
		return nil, err
	}

	_, err = s.db.DAClientStore.Add(senderID, blkTime)
	if err != nil {
		return nil, err
	}

	daSubmit.SenderID = senderID

	return daSubmit, nil
}

func (s *CatchupSyncer) decodeCommitVerified(blkTime time.Time, log types.Log) (*store.DASubmit, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daEntranceAddr) || sig != s.commitVerifiedSig {
		return nil, nil
	}

	daSubmit, err := store.NewDASubmitVerified(blkTime, log, nhContract.DummyDAEntranceFilterer())
	if err != nil {
		return nil, err
	}

	return daSubmit, nil
}

func (s *CatchupSyncer) decodeDAReward(blkTime time.Time, log types.Log) (*store.DAReward, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daEntranceAddr) || sig != s.daRewardSig {
		return nil, nil
	}

	daReward, err := store.NewDAReward(blkTime, log, nhContract.DummyDAEntranceFilterer())
	if err != nil {
		return nil, err
	}

	minerID, err := s.db.AddressStore.Add(daReward.Miner, blkTime)
	if err != nil {
		return nil, err
	}

	daReward.MinerID = minerID

	return daReward, nil
}

func interrupted(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}
