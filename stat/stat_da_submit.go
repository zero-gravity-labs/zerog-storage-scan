package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type StatDASubmit struct {
	*BaseStat
	statType string
}

func MustNewStatDASubmit(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statDASubmit := &StatDASubmit{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalDASubmit,
	}

	return &AbsStat{
		Stat: statDASubmit,
		sdk:  baseStat.Sdk,
	}
}

func (sds *StatDASubmit) nextTimeRange() (*TimeRange, error) {
	lastStat, err := sds.DB.DASubmitStatStore.LastByType(sds.statType)
	if err != nil {
		return nil, err
	}

	var nextRangeStart time.Time
	if lastStat == nil {
		nextRangeStart = sds.StartTime
	} else {
		t := lastStat.StatTime.Add(store.Intervals[sds.statType])
		nextRangeStart = t
	}

	timeRange, err := sds.calStatRange(nextRangeStart, store.Intervals[sds.statType])
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (sds *StatDASubmit) calculateStat(tr TimeRange) error {
	stat, err := sds.statByTimeRange(tr.start, tr.end, sds.statType)
	if err != nil {
		return err
	}
	hStat, err := sds.statByTimeRange(time.Time{}, tr.end, store.Hour)
	if err != nil {
		return err
	}
	dStat, err := sds.statByTimeRange(time.Time{}, tr.end, store.Day)
	if err != nil {
		return err
	}

	stats := []*store.DASubmitStat{stat, hStat, dStat}
	return sds.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := sds.DB.DASubmitStatStore.Del(dbTx, hStat); err != nil {
			return errors.WithMessage(err, "failed to del hour stat")
		}
		if err := sds.DB.DASubmitStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := sds.DB.DASubmitStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (sds *StatDASubmit) statByTimeRange(start, end time.Time, statType string) (*store.DASubmitStat, error) {
	// cal range start if not exist
	nilTime := time.Time{}
	if start == nilTime {
		rangeStart, err := sds.calStatRangeStart(end, statType)
		if err != nil {
			return nil, err
		}
		start = rangeStart
	}

	// delta count
	delta, err := sds.DB.DASubmitStore.Count(start, end)
	if err != nil {
		return nil, err
	}

	// total count of the previous range
	preTimeRange, err := sds.calStatRange(start, -store.Intervals[statType])
	if err != nil {
		return nil, err
	}

	var preStat store.DASubmitStat
	exist, err := sds.DB.DASubmitStatStore.Exists(&preStat, "stat_type = ? and stat_time = ?",
		statType, preTimeRange.start)
	if err != nil {
		return nil, err
	}

	if !exist {
		total, err := sds.DB.DASubmitStore.Count(time.Time{}, start)
		if err != nil {
			return nil, err
		}
		preStat.BlobTotal = total.Blobs
		preStat.DataSizeTotal = total.Blobs * 32
		preStat.StorageFeeTotal = total.StorageFee
	}

	return &store.DASubmitStat{
		StatTime: start,
		StatType: statType,

		BlobNew:         delta.Blobs,
		BlobTotal:       preStat.BlobTotal + delta.Blobs,
		DataSizeNew:     delta.Blobs * 32,
		DataSizeTotal:   (preStat.BlobTotal + delta.Blobs) * 32,
		StorageFeeNew:   delta.StorageFee,
		StorageFeeTotal: preStat.StorageFeeTotal.Add(delta.StorageFee),
	}, nil
}
