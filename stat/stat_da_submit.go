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

func MustNewDAStatSubmit(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statDASubmit := &StatDASubmit{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalDailyDASubmit,
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
	stat, err := sds.statBasicRange(tr)
	if err != nil {
		return err
	}
	hStat, err := sds.statRange(tr.end, sds.statType, store.Hour, stat)
	if err != nil {
		return err
	}
	dStat, err := sds.statRange(tr.end, sds.statType, store.Day, stat)
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

func (sds *StatDASubmit) statBasicRange(tr TimeRange) (*store.DASubmitStat, error) {
	delta, err := sds.DB.DASubmitStore.Count(tr.start, tr.end)
	if err != nil {
		return nil, err
	}
	total, err := sds.DB.DASubmitStatStore.Sum(time.Time{}, tr.start, sds.statType)
	if err != nil {
		return nil, err
	}

	return &store.DASubmitStat{
		StatTime: tr.start,
		StatType: sds.statType,

		BlobNew:         delta.Blobs,
		BlobTotal:       total.Blobs + delta.Blobs,
		DataSizeNew:     delta.Blobs * 32,
		DataSizeTotal:   (total.Blobs + delta.Blobs) * 32,
		StorageFeeNew:   delta.StorageFee,
		StorageFeeTotal: total.StorageFee.Add(delta.StorageFee),
	}, nil
}

func (sds *StatDASubmit) statRange(rangEnd time.Time, srcStatType, descStatType string, latestStat *store.DASubmitStat) (*store.DASubmitStat, error) {
	rangeStart, err := sds.calStatRangeStart(rangEnd, descStatType)
	if err != nil {
		return nil, err
	}

	srcStat, err := sds.DB.DASubmitStatStore.Sum(rangeStart, rangEnd, srcStatType)
	if err != nil {
		return nil, err
	}
	destStat, err := sds.DB.DASubmitStatStore.Sum(time.Time{}, rangeStart, descStatType)
	if err != nil {
		return nil, err
	}

	if latestStat != nil {
		srcStat.Blobs += latestStat.BlobNew
		srcStat.StorageFee = srcStat.StorageFee.Add(latestStat.StorageFeeNew)
	}

	return &store.DASubmitStat{
		StatTime: rangeStart,
		StatType: descStatType,

		BlobNew:         srcStat.Blobs,
		BlobTotal:       destStat.Blobs + srcStat.Blobs,
		DataSizeNew:     srcStat.Blobs * 32,
		DataSizeTotal:   (destStat.Blobs + srcStat.Blobs) * 32,
		StorageFeeNew:   srcStat.StorageFee,
		StorageFeeTotal: destStat.StorageFee.Add(srcStat.StorageFee),
	}, nil
}
