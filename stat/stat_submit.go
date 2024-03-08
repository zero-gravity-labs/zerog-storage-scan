package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type StatSubmit struct {
	*BaseStat
	statType string
}

func MustNewStatSubmit(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statSubmit := &StatSubmit{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalDailySubmit,
	}

	return &AbsStat{
		Stat: statSubmit,
		sdk:  baseStat.Sdk,
	}
}

func (ts *StatSubmit) nextTimeRange() (*TimeRange, error) {
	lastStat, err := ts.DB.SubmitStatStore.LastByType(ts.statType)
	if err != nil {
		return nil, err
	}

	var nextRangeStart time.Time
	if lastStat == nil {
		nextRangeStart = ts.StartTime
	} else {
		t := lastStat.StatTime.Add(store.Intervals[ts.statType])
		nextRangeStart = t
	}

	timeRange, err := ts.calStatRange(nextRangeStart, store.Intervals[ts.statType])
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (ts *StatSubmit) calculateStat(tr TimeRange) error {
	stat, err := ts.statBasicRange(tr)
	if err != nil {
		return err
	}
	hStat, err := ts.statRange(tr.end, ts.statType, store.Hour, stat)
	if err != nil {
		return err
	}
	dStat, err := ts.statRange(tr.end, ts.statType, store.Day, stat)
	if err != nil {
		return err
	}

	stats := []*store.SubmitStat{stat, hStat, dStat}
	return ts.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := ts.DB.SubmitStatStore.Del(dbTx, hStat); err != nil {
			return errors.WithMessage(err, "failed to del hour stat")
		}
		if err := ts.DB.SubmitStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := ts.DB.SubmitStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (ts *StatSubmit) statBasicRange(tr TimeRange) (*store.SubmitStat, error) {
	delta, err := ts.DB.SubmitStore.Count(tr.start, tr.end)
	if err != nil {
		return nil, err
	}
	total, err := ts.DB.SubmitStatStore.Sum(time.Time{}, tr.start, ts.statType)
	if err != nil {
		return nil, err
	}

	return &store.SubmitStat{
		StatTime:     tr.start,
		StatType:     ts.statType,
		FileCount:    delta.FileCount,
		FileTotal:    total.FileCount + delta.FileCount,
		DataSize:     delta.DataSize,
		DataTotal:    total.DataSize + delta.DataSize,
		BaseFee:      delta.BaseFee,
		BaseFeeTotal: total.BaseFee.Add(delta.BaseFee),
		TxCount:      delta.TxCount,
		TxTotal:      total.TxCount + delta.TxCount,
	}, nil
}

func (ts *StatSubmit) statRange(rangEnd time.Time, srcStatType, descStatType string, latestStat *store.SubmitStat) (*store.SubmitStat, error) {
	rangeStart, err := ts.calStatRangeStart(rangEnd, descStatType)
	if err != nil {
		return nil, err
	}

	srcStat, err := ts.DB.SubmitStatStore.Sum(rangeStart, rangEnd, srcStatType)
	if err != nil {
		return nil, err
	}
	destStat, err := ts.DB.SubmitStatStore.Sum(time.Time{}, rangeStart, descStatType)
	if err != nil {
		return nil, err
	}

	if latestStat != nil {
		srcStat.FileCount += latestStat.FileCount
		srcStat.DataSize += latestStat.DataSize
		srcStat.BaseFee = srcStat.BaseFee.Add(latestStat.BaseFee)
		srcStat.TxCount += latestStat.TxCount
	}

	return &store.SubmitStat{
		StatTime:     rangeStart,
		StatType:     descStatType,
		FileCount:    srcStat.FileCount,
		FileTotal:    destStat.FileCount + srcStat.FileCount,
		DataSize:     srcStat.DataSize,
		DataTotal:    destStat.DataSize + srcStat.DataSize,
		BaseFee:      srcStat.BaseFee,
		BaseFeeTotal: destStat.BaseFee.Add(srcStat.BaseFee),
		TxCount:      srcStat.TxCount,
		TxTotal:      destStat.TxCount + srcStat.TxCount,
	}, nil
}
