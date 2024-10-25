package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StatAddress struct {
	*BaseStat
	statType string
}

func MustNewStatAddress(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statAddress := &StatAddress{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalAddress,
	}

	return &AbsStat{
		Stat: statAddress,
		sdk:  baseStat.Sdk,
	}
}

func (sa *StatAddress) nextTimeRange() (*TimeRange, error) {
	lastStat, err := sa.DB.AddressStatStore.LastByType(sa.statType)
	if err != nil {
		return nil, err
	}

	var nextRangeStart time.Time
	if lastStat == nil {
		nextRangeStart = sa.StartTime
	} else {
		t := lastStat.StatTime.Add(store.Intervals[sa.statType])
		nextRangeStart = t
	}

	timeRange, err := sa.calStatRange(nextRangeStart, store.Intervals[sa.statType])
	if err != nil {
		return nil, err
	}
	timeRange.intervalType = sa.statType

	return timeRange, nil
}

func (sa *StatAddress) calculateStat(tr TimeRange) error {
	stat, err := sa.statByTimeRange(tr.start, tr.end, sa.statType)
	if err != nil {
		return err
	}
	hStat, err := sa.statByTimeRange(time.Time{}, tr.end, store.Hour)
	if err != nil {
		return err
	}
	dStat, err := sa.statByTimeRange(time.Time{}, tr.end, store.Day)
	if err != nil {
		return err
	}

	stats := []*store.AddressStat{stat, hStat, dStat}
	return sa.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := sa.DB.AddressStatStore.Del(dbTx, hStat); err != nil {
			return errors.WithMessage(err, "failed to del hour stat")
		}
		if err := sa.DB.AddressStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := sa.DB.AddressStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (sa *StatAddress) statByTimeRange(start, end time.Time, statType string) (*store.AddressStat, error) {
	// cal range start if not exist
	nilTime := time.Time{}
	if start == nilTime {
		rangeStart, err := sa.calStatRangeStart(end, statType)
		if err != nil {
			return nil, err
		}
		start = rangeStart
	}

	// delta count
	delta, err := sa.DB.AddressStore.Count(start, end)
	if err != nil {
		return nil, err
	}

	// total count of the previous range
	preTimeRange, err := sa.calStatRange(start, -store.Intervals[statType])
	if err != nil {
		return nil, err
	}

	var addressStat store.AddressStat
	exist, err := sa.DB.AddressStatStore.Exists(&addressStat, "stat_type = ? and stat_time = ?",
		statType, preTimeRange.start)
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, err
	}

	var total uint64
	if !exist {
		count, err := sa.DB.AddressStore.Count(time.Time{}, start)
		if err != nil {
			return nil, err
		}
		total = count
	} else {
		total = addressStat.AddrTotal
	}

	// active count
	submitStatResult, err := sa.DB.SubmitStore.Count(start, end)
	if err != nil {
		return nil, err
	}

	return &store.AddressStat{
		StatTime:   start,
		StatType:   statType,
		AddrNew:    delta,
		AddrActive: submitStatResult.SenderCount,
		AddrTotal:  total + delta,
	}, nil
}
