package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StatMiner struct {
	*BaseStat
	statType string
}

func MustNewStatMiner(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statMiner := &StatMiner{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalMiner,
	}

	return &AbsStat{
		Stat: statMiner,
		sdk:  baseStat.Sdk,
	}
}

func (sm *StatMiner) nextTimeRange() (*TimeRange, error) {
	lastStat, err := sm.DB.MinerStatStore.LastByType(sm.statType)
	if err != nil {
		return nil, err
	}

	var nextRangeStart time.Time
	if lastStat == nil {
		nextRangeStart = sm.StartTime
	} else {
		t := lastStat.StatTime.Add(store.Intervals[sm.statType])
		nextRangeStart = t
	}

	timeRange, err := sm.calStatRange(nextRangeStart, store.Intervals[sm.statType])
	if err != nil {
		return nil, err
	}
	timeRange.intervalType = sm.statType

	return timeRange, nil
}

func (sm *StatMiner) calculateStat(tr TimeRange) error {
	stat, err := sm.statByTimeRange(tr.start, tr.end, sm.statType)
	if err != nil {
		return err
	}
	dStat, err := sm.statByTimeRange(time.Time{}, tr.end, store.Day)
	if err != nil {
		return err
	}

	stats := []*store.MinerStat{stat, dStat}
	return sm.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := sm.DB.MinerStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := sm.DB.MinerStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (sm *StatMiner) statByTimeRange(start, end time.Time, statType string) (*store.MinerStat, error) {
	// cal range start if not exist
	nilTime := time.Time{}
	if start == nilTime {
		rangeStart, err := sm.calStatRangeStart(end, statType)
		if err != nil {
			return nil, err
		}
		start = rangeStart
	}

	// delta count
	delta, err := sm.DB.MinerStore.Count(start, end)
	if err != nil {
		return nil, err
	}

	// total count of the previous range
	preTimeRange, err := sm.calStatRange(start, -store.Intervals[statType])
	if err != nil {
		return nil, err
	}

	var minerStat store.MinerStat
	exist, err := sm.DB.MinerStatStore.Exists(&minerStat, "stat_type = ? and stat_time = ?",
		statType, preTimeRange.start)
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, err
	}

	var total uint64
	if !exist {
		count, err := sm.DB.MinerStore.Count(time.Time{}, start)
		if err != nil {
			return nil, err
		}
		total = count
	} else {
		total = minerStat.MinerTotal
	}

	// active count
	countActive, err := sm.DB.RewardStore.CountActive(start, end)
	if err != nil {
		return nil, err
	}

	return &store.MinerStat{
		StatTime:    start,
		StatType:    statType,
		MinerNew:    delta,
		MinerActive: countActive,
		MinerTotal:  total + delta,
	}, nil
}
