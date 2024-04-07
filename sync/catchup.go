package sync

import (
	"context"
	"strings"
	"time"

	nhContract "github.com/0glabs/0g-storage-scan/contract"
	"github.com/0glabs/0g-storage-scan/store"
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
	flowAddr       string
	flowSubmitSig  string
	rewardAddr     string
	rewardSig      string
	addresses      []common.Address
	topics         [][]common.Hash
}

func MustNewCatchupSyncer(sdk *web3go.Client, db *store.MysqlStore, conf SyncConfig) *CatchupSyncer {
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

	return &CatchupSyncer{
		conf:          &conf,
		sdk:           sdk,
		db:            db,
		flowAddr:      flow.Address,
		flowSubmitSig: flow.SubmitEventSignature,
		rewardAddr:    reward.Address,
		rewardSig:     reward.RewardEventSignature,
		addresses:     []common.Address{common.HexToAddress(flow.Address), common.HexToAddress(reward.Address)},
		topics:        [][]common.Hash{{common.HexToHash(flow.SubmitEventSignature), common.HexToHash(reward.RewardEventSignature)}},
	}
}

func (s *CatchupSyncer) Sync(ctx context.Context) {
	logrus.Info("Catchup syncer starting to sync data")
	for {
		needProcess := s.tryBlockRange()
		if !needProcess || s.interrupted(ctx) {
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
			bn2TimeMap, err = batchGetBlockTimes(ctx, s.sdk, blockNums, s.conf.BatchBlocksOnBatchCall)
		}
		if err != nil {
			return err
		}

		rangeEndBlock, err := s.sdk.Eth.BlockByNumber(types.BlockNumber(end), false)
		if err != nil {
			return err
		}

		block := store.NewBlock(rangeEndBlock)
		submits, rewards, err := s.convertLogs(logs, bn2TimeMap)
		if err != nil {
			return err
		}

		err = s.db.Push(block, submits, rewards)
		if err != nil {
			return err
		}

		if end >= rangeEnd || s.interrupted(ctx) {
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
		logs, err := batchGetLogs(w3c, start, end, addresses, topics)
		if err == nil {
			return logs, start, end, nil
		}

		if strings.Contains(err.Error(), "please narrow down your filter condition") || // espace
			strings.Contains(err.Error(), "block range") || // bsc
			strings.Contains(err.Error(), "blocks distance") { // evmos
			end = start + (end-start)/2
			continue
		}

		return nil, 0, 0, err
	}
}

func (s *CatchupSyncer) tryBlockRange() bool {
	for try := 1; ; try++ {
		err := s.updateBlockRange()
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

func (s *CatchupSyncer) updateBlockRange() error {
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
	if err != nil {
		return errors.WithMessage(err, "failed to get finalized block")
	}
	s.finalizedBlock = finalizedBlock.Number.Uint64()

	return nil
}

func (s *CatchupSyncer) convertLogs(logs []types.Log, bn2TimeMap map[uint64]uint64) ([]*store.Submit,
	[]*store.Reward, error) {
	var submits []*store.Submit
	var rewards []*store.Reward

	for _, log := range logs {
		if log.Removed {
			continue
		}

		ts := bn2TimeMap[log.BlockNumber]
		blockTime := time.Unix(int64(ts), 0)

		submit, err := s.decodeSubmit(blockTime, log)
		if err != nil {
			return nil, nil, err
		}
		if submit != nil {
			submits = append(submits, submit)
		}

		reward, err := s.decodeReward(blockTime, log)
		if err != nil {
			return nil, nil, err
		}
		if reward != nil {
			rewards = append(rewards, reward)
		}
	}

	return submits, rewards, nil
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

	reward.MinerID = minerID

	return reward, nil
}

func (s *CatchupSyncer) interrupted(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}
