package stat

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	batchInBns = 1000
)

type TopnReward struct {
	*BaseStat
}

func MustNewTopnReward(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client) *AbsTopn[StatRange] {
	baseStat := &BaseStat{
		Config: cfg,
		DB:     db,
		Sdk:    sdk,
	}

	topnReward := &TopnReward{
		BaseStat: baseStat,
	}
	topnReward.mustLoadLastPos()

	return &AbsTopn[StatRange]{
		Topn: topnReward,
	}
}

func (ts *TopnReward) mustLoadLastPos() {
	loaded, err := ts.loadLastPos(store.StatTopnRewardBn)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load stat pos from db")
	}

	// Reward bn is set config value if not loaded.
	if !loaded {
		ts.currentPos = ts.Config.BlockOnStatBegin
	}
}

func (ts *TopnReward) nextStatRange() (*StatRange, error) {
	minPos := ts.currentPos
	maxPos := ts.currentPos + uint64(batchInBns) - 1

	block, err := ts.Sdk.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if err != nil {
		return nil, err
	}
	maxPosFinalized := block.Number.Uint64()

	if maxPosFinalized < minPos {
		return nil, ErrMinPosNotFinalized
	}
	if maxPosFinalized < maxPos {
		maxPos = maxPosFinalized
	}

	return &StatRange{minPos, maxPos}, nil
}

func (ts *TopnReward) calculateStat(r StatRange) error {
	groupedRewards, err := ts.DB.RewardStore.GroupByMiner(r.minPos, r.maxPos)
	if err != nil {
		return err
	}

	miners := make([]store.Miner, 0)
	for _, reward := range groupedRewards {
		miners = append(miners, store.Miner{
			ID:        reward.MinerID,
			Amount:    reward.Amount,
			UpdatedAt: reward.UpdatedAt,
		})
	}

	if err := ts.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(miners) > 0 {
			if err := ts.DB.MinerStore.BatchIncreaseStat(dbTx, miners); err != nil {
				return errors.WithMessage(err, "Failed to batch update miners for topn")
			}
		}
		if err := ts.DB.ConfigStore.Upsert(dbTx, store.StatTopnRewardBn,
			strconv.FormatUint(r.maxPos, 10)); err != nil {
			return errors.WithMessage(err, "Failed to batch update bn for topn")
		}
		return nil
	}); err != nil {
		return err
	}

	ts.currentPos = r.maxPos + 1

	return nil
}
