package sync

import (
	"context"
	"strings"
	"time"

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
	baseSyncer
	finalizedBlock uint64

	alertChannel  string
	healthReport  health.TimedCounterConfig
	nodeRpcHealth health.TimedCounter
}

func MustNewCatchupSyncer(sdk *web3go.Client, db *store.MysqlStore, conf SyncConfig, alertChannel string,
	healthReport health.TimedCounterConfig) *CatchupSyncer {
	var flow flowConfig
	viperUtil.MustUnmarshalKey("flow", &flow)
	var reward rewardConfig
	viperUtil.MustUnmarshalKey("reward", &reward)
	var daEntrance daEntranceConfig
	viperUtil.MustUnmarshalKey("daEntrance", &daEntrance)
	var daSigners daSignersConfig
	viperUtil.MustUnmarshalKey("daSigners", &daSigners)

	base := baseSyncer{
		conf: &conf,
		sdk:  sdk,
		db:   db,

		flowAddr:        flow.Address,
		flowSubmitSig:   flow.SubmitEventSignature,
		flowNewEpochSig: flow.NewEpochEventSignature,
		rewardAddr:      reward.Address,
		rewardSig:       reward.RewardEventSignature,

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
			common.HexToHash(flow.NewEpochEventSignature),
			common.HexToHash(reward.RewardEventSignature),
			common.HexToHash(daEntrance.DataUploadSignature),
			common.HexToHash(daEntrance.ErasureCommitmentVerifiedSignature),
			common.HexToHash(daEntrance.DARewardSignature),
			common.HexToHash(daSigners.NewSignerSignature),
			common.HexToHash(daSigners.SocketUpdatedSignature),
		}},
	}
	return &CatchupSyncer{
		baseSyncer:    base,
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

			_, en, e := s.findClosedInterval(err.Error(), `suggested block range is \[(\d+)\s*,\s*(\d+)\]`)
			if e != nil {
				end = start + (end-start)/2
			} else {
				end = en // use suggested range end when matched
			}

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

func interrupted(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}
