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
	stat, err := ts.statByTimeRange(tr.start, tr.end, ts.statType)
	if err != nil {
		return err
	}
	hStat, err := ts.statByTimeRange(time.Time{}, tr.end, store.Hour)
	if err != nil {
		return err
	}
	dStat, err := ts.statByTimeRange(time.Time{}, tr.end, store.Day)
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

func (ts *StatSubmit) statByTimeRange(start, end time.Time, statType string) (*store.SubmitStat, error) {
	// cal range start if not exist
	nilTime := time.Time{}
	if start == nilTime {
		rangeStart, err := ts.calStatRangeStart(end, statType)
		if err != nil {
			return nil, err
		}
		start = rangeStart
	}

	// delta count
	delta, err := ts.DB.SubmitStore.Count(start, end)
	if err != nil {
		return nil, err
	}

	// total count of the previous range
	preTimeRange, err := ts.calStatRange(start, -store.Intervals[statType])
	if err != nil {
		return nil, err
	}

	var preStat store.SubmitStat
	exist, err := ts.DB.SubmitStatStore.Exists(&preStat, "stat_type = ? and stat_time = ?",
		statType, preTimeRange.start)
	if err != nil {
		return nil, err
	}

	if !exist {
		total, err := ts.DB.SubmitStore.Count(time.Time{}, start)
		if err != nil {
			return nil, err
		}
		preStat.FileTotal = total.FileCount
		preStat.DataTotal = total.DataSize
		preStat.BaseFeeTotal = total.BaseFee
		preStat.TxTotal = total.TxCount
	}

	return &store.SubmitStat{
		StatTime:     start,
		StatType:     statType,
		FileCount:    delta.FileCount,
		FileTotal:    preStat.FileTotal + delta.FileCount,
		DataSize:     delta.DataSize,
		DataTotal:    preStat.DataTotal + delta.DataSize,
		BaseFee:      delta.BaseFee,
		BaseFeeTotal: preStat.BaseFeeTotal.Add(delta.BaseFee),
		TxCount:      delta.TxCount,
		TxTotal:      preStat.TxTotal + delta.TxCount,
	}, nil
}
