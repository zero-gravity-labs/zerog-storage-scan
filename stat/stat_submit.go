package stat

import (
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"gorm.io/gorm"
	"time"
)

type StatSubmit struct {
	*BaseStat
	statType string
}

func MustNewStatSubmit(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime *time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		Db:        db,
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
	lastStat, err := ts.Db.SubmitStatStore.LastByType(ts.statType)
	if err != nil {
		return nil, err
	}

	var nextRangeStart *time.Time
	if lastStat == nil {
		nextRangeStart = ts.StartTime
	} else {
		t := lastStat.StatTime.Add(Intervals[ts.statType])
		nextRangeStart = &t
	}

	timeRange, err := ts.calStatRange(nextRangeStart, Intervals[ts.statType])
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (ts *StatSubmit) calculateStat(tr *TimeRange) error {
	stat, err := ts.statBasicRange(tr)
	if err != nil {
		return err
	}
	hStat, err := ts.statRange(tr.end, ts.statType, Hour, stat)
	if err != nil {
		return err
	}
	dStat, err := ts.statRange(tr.end, ts.statType, Day, stat)
	if err != nil {
		return err
	}

	stats := []*store.SubmitStat{stat, hStat, dStat}
	return ts.Db.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := ts.Db.SubmitStatStore.Del(dbTx, hStat); err != nil {
			return errors.WithMessage(err, "failed to del hour stat")
		}
		if err := ts.Db.SubmitStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := ts.Db.SubmitStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (ts *StatSubmit) statBasicRange(tr *TimeRange) (*store.SubmitStat, error) {
	delta, err := ts.Db.SubmitStore.Count(*tr.start, *tr.end)
	if err != nil {
		return nil, err
	}
	total, err := ts.Db.SubmitStatStore.Sum(nil, tr.start, ts.statType)
	if err != nil {
		return nil, err
	}

	return &store.SubmitStat{
		StatTime:       tr.start,
		StatType:       ts.statType,
		FileCount:      delta.FileCount,
		FileTotal:      total.FileCount + delta.FileCount,
		DataSize:       delta.DataSize,
		DataTotal:      total.DataSize + delta.DataSize,
		BasicCost:      delta.BasicCost,
		BasicCostTotal: total.BasicCost + delta.BasicCost,
	}, nil
}

func (ts *StatSubmit) statRange(rangEnd *time.Time, srcStatType, descStatType string, latestStat *store.SubmitStat) (*store.SubmitStat, error) {
	rangeStart, err := ts.calStatRangeStart(rangEnd, descStatType)
	if err != nil {
		return nil, err
	}

	srcStat, err := ts.Db.SubmitStatStore.Sum(rangeStart, rangEnd, srcStatType)
	if err != nil {
		return nil, err
	}
	destStat, err := ts.Db.SubmitStatStore.Sum(nil, rangeStart, descStatType)
	if err != nil {
		return nil, err
	}

	if latestStat != nil {
		srcStat.FileCount += latestStat.FileCount
		srcStat.DataSize += latestStat.DataSize
		srcStat.BasicCost += latestStat.BasicCost
	}

	return &store.SubmitStat{
		StatTime:       rangeStart,
		StatType:       descStatType,
		FileCount:      srcStat.FileCount,
		FileTotal:      destStat.FileCount + srcStat.FileCount,
		DataSize:       srcStat.DataSize,
		DataTotal:      destStat.DataSize + srcStat.DataSize,
		BasicCost:      srcStat.BasicCost,
		BasicCostTotal: destStat.BasicCost + srcStat.BasicCost,
	}, nil
}
