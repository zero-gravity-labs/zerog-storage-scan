package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type StatReward struct {
	*BaseStat
	statType string
}

func MustNewStatReward(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statReward := &StatReward{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalReward,
	}

	return &AbsStat{
		Stat: statReward,
		sdk:  baseStat.Sdk,
	}
}

func (sr *StatReward) nextTimeRange() (*TimeRange, error) {
	lastStat, err := sr.DB.RewardStatStore.LastByType(sr.statType)
	if err != nil {
		return nil, err
	}

	var nextRangeStart time.Time
	if lastStat == nil {
		nextRangeStart = sr.StartTime
	} else {
		t := lastStat.StatTime.Add(store.Intervals[sr.statType])
		nextRangeStart = t
	}

	timeRange, err := sr.calStatRange(nextRangeStart, store.Intervals[sr.statType])
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (sr *StatReward) calculateStat(tr TimeRange) error {
	stat, err := sr.statByTimeRange(tr.start, tr.end, sr.statType)
	if err != nil {
		return err
	}
	hStat, err := sr.statByTimeRange(time.Time{}, tr.end, store.Hour)
	if err != nil {
		return err
	}
	dStat, err := sr.statByTimeRange(time.Time{}, tr.end, store.Day)
	if err != nil {
		return err
	}

	stats := []*store.RewardStat{stat, hStat, dStat}
	return sr.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := sr.DB.RewardStatStore.Del(dbTx, hStat); err != nil {
			return errors.WithMessage(err, "failed to del hour stat")
		}
		if err := sr.DB.RewardStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := sr.DB.RewardStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (sr *StatReward) statByTimeRange(start, end time.Time, statType string) (*store.RewardStat, error) {
	// cal range start if not exist
	nilTime := time.Time{}
	if start == nilTime {
		rangeStart, err := sr.calStatRangeStart(end, statType)
		if err != nil {
			return nil, err
		}
		start = rangeStart
	}

	// delta count
	delta, err := sr.DB.RewardStore.Sum(start, end)
	if err != nil {
		return nil, err
	}

	// total count of the previous range
	preTimeRange, err := sr.calStatRange(start, -store.Intervals[statType])
	if err != nil {
		return nil, err
	}

	var preStat store.RewardStat
	exist, err := sr.DB.RewardStatStore.Exists(&preStat, "stat_type = ? and stat_time = ?",
		statType, preTimeRange.start)
	if err != nil {
		return nil, err
	}

	if !exist {
		total, err := sr.DB.RewardStore.Sum(time.Time{}, start)
		if err != nil {
			return nil, err
		}
		preStat.RewardTotal = *total
	}

	return &store.RewardStat{
		StatTime:    start,
		StatType:    statType,
		RewardNew:   *delta,
		RewardTotal: preStat.RewardTotal.Add(*delta),
	}, nil
}
