package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StatDAClient struct {
	*BaseStat
	statType string
}

func MustNewStatDAClient(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statDAClient := &StatDAClient{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalDailyDAClient,
	}

	return &AbsStat{
		Stat: statDAClient,
		sdk:  baseStat.Sdk,
	}
}

func (sm *StatDAClient) nextTimeRange() (*TimeRange, error) {
	lastStat, err := sm.DB.DAClientStatStore.LastByType(sm.statType)
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

func (sm *StatDAClient) calculateStat(tr TimeRange) error {
	stat, err := sm.statByTimeRange(tr.start, tr.end, sm.statType)
	if err != nil {
		return err
	}
	dStat, err := sm.statByTimeRange(time.Time{}, tr.end, store.Day)
	if err != nil {
		return err
	}

	stats := []*store.DAClientStat{stat, dStat}
	return sm.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := sm.DB.DAClientStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := sm.DB.DAClientStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (sm *StatDAClient) statByTimeRange(start, end time.Time, statType string) (*store.DAClientStat, error) {
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
	delta, err := sm.DB.DAClientStore.Count(start, end)
	if err != nil {
		return nil, err
	}

	// total count of the previous range
	preTimeRange, err := sm.calStatRange(start, -store.Intervals[statType])
	if err != nil {
		return nil, err
	}

	var daClientStat store.DAClientStat
	exist, err := sm.DB.MinerStatStore.Exists(&daClientStat, "stat_type = ? and stat_time = ?",
		statType, preTimeRange.start)
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, err
	}

	var total uint64
	if !exist {
		count, err := sm.DB.DAClientStore.Count(time.Time{}, start)
		if err != nil {
			return nil, err
		}
		total = count
	} else {
		total = daClientStat.ClientTotal
	}

	// active count
	active, err := sm.DB.DASubmitStore.CountActive(start, end)
	if err != nil {
		return nil, err
	}

	return &store.DAClientStat{
		StatTime:     start,
		StatType:     statType,
		ClientNew:    delta,
		ClientActive: active,
		ClientTotal:  total + delta,
	}, nil
}
