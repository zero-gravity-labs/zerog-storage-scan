package stat

import (
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type StatTx struct {
	*BaseStat
	statType string
}

func MustNewStatTx(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime *time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		Db:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statTx := &StatTx{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalDailyTx,
	}

	return &AbsStat{
		Stat: statTx,
		sdk:  baseStat.Sdk,
	}
}

func (ts *StatTx) nextTimeRange() (*TimeRange, error) {
	lastStat, err := ts.Db.TxStatStore.LastByType(ts.statType)
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

func (ts *StatTx) calculateStat(tr *TimeRange) error {
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

	stats := []*store.TxStat{stat, hStat, dStat}
	return ts.Db.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := ts.Db.TxStatStore.Del(dbTx, hStat); err != nil {
			return errors.WithMessage(err, "failed to del hour stat")
		}
		if err := ts.Db.TxStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := ts.Db.TxStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (ts *StatTx) statBasicRange(tr *TimeRange) (*store.TxStat, error) {
	count, err := ts.Db.TxStore.Count(tr.start, tr.end)
	if err != nil {
		return nil, err
	}
	total, err := ts.Db.TxStatStore.Sum(nil, tr.start, ts.statType)
	if err != nil {
		return nil, err
	}

	return &store.TxStat{
		StatTime: tr.start,
		StatType: ts.statType,
		TxCount:  count,
		TxTotal:  total + count,
	}, nil
}

func (ts *StatTx) statRange(rangEnd *time.Time, srcStatType, descStatType string, latestStat *store.TxStat) (*store.TxStat, error) {
	rangeStart, err := ts.calStatRangeStart(rangEnd, descStatType)
	if err != nil {
		return nil, err
	}

	count, err := ts.Db.TxStatStore.Sum(rangeStart, rangEnd, srcStatType)
	if err != nil {
		return nil, err
	}
	total, err := ts.Db.TxStatStore.Sum(nil, rangeStart, descStatType)
	if err != nil {
		return nil, err
	}

	if latestStat != nil {
		count += latestStat.TxCount
	}

	return &store.TxStat{
		StatTime: rangeStart,
		StatType: descStatType,
		TxCount:  count,
		TxTotal:  total + count,
	}, nil
}
