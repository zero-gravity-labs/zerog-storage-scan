package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TopnRewardRange struct {
	*BaseStat
	statType string
}

func MustNewTopnRewardRange(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	topnRewardRange := &TopnRewardRange{
		BaseStat: baseStat,
		statType: baseStat.Config.MinTopnIntervalReward,
	}

	return &AbsStat{
		Stat: topnRewardRange,
		sdk:  baseStat.Sdk,
	}
}

func (trr *TopnRewardRange) nextTimeRange() (*TimeRange, error) {
	value, ok, err := trr.DB.ConfigStore.Get(store.StatTopnRewardTime)
	if err != nil {
		return nil, err
	}

	var nextRangeStart time.Time
	if ok {
		lastStatTime, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return nil, err
		}
		nextRangeStart = lastStatTime.Add(store.Intervals[trr.statType])
	} else {
		nextRangeStart = trr.StartTime
	}

	timeRange, err := trr.calStatRange(nextRangeStart, store.Intervals[trr.statType])
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (trr *TopnRewardRange) calculateStat(r TimeRange) error {
	rewards, err := trr.DB.GroupByMinerByTime(r.start, r.end)
	if err != nil {
		return err
	}

	rewardStats := make([]store.RewardTopnStat, 0)
	statTime := r.start.Truncate(time.Hour)
	for _, r := range rewards {
		rewardStats = append(rewardStats, store.RewardTopnStat{
			StatTime:  statTime,
			AddressID: r.MinerID,
			Amount:    r.Amount,
		})
	}

	return trr.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(rewardStats) > 0 {
			if err := trr.DB.RewardTopnStatStore.BatchDeltaUpsert(dbTx, rewardStats); err != nil {
				return errors.WithMessage(err, "failed to batch delta upsert rewards")
			}
		}
		if err := trr.DB.ConfigStore.Upsert(dbTx, store.StatTopnRewardTime, r.start.Format(time.RFC3339)); err != nil {
			return err
		}
		return nil
	})
}
