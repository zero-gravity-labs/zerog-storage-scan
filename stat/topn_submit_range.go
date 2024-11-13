package stat

import (
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TopnSubmitRange struct {
	*BaseStat
	statType string
}

func MustNewTopnSubmitRange(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		DB:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	topnSubmitRange := &TopnSubmitRange{
		BaseStat: baseStat,
		statType: baseStat.Config.MinTopnIntervalSubmit,
	}

	return &AbsStat{
		Stat: topnSubmitRange,
		sdk:  baseStat.Sdk,
	}
}

func (tsr *TopnSubmitRange) nextTimeRange() (*TimeRange, error) {
	value, ok, err := tsr.DB.ConfigStore.Get(store.StatTopnSubmitTime)
	if err != nil {
		return nil, err
	}

	var nextRangeStart time.Time
	if ok {
		lastStatTime, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return nil, err
		}
		nextRangeStart = lastStatTime.Add(store.Intervals[tsr.statType])
	} else {
		nextRangeStart = tsr.StartTime
	}

	timeRange, err := tsr.calStatRange(nextRangeStart, store.Intervals[tsr.statType])
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (tsr *TopnSubmitRange) calculateStat(r TimeRange) error {
	submits, err := tsr.DB.GroupBySenderByTime(r.start, r.end)
	if err != nil {
		return err
	}

	submitStats := make([]store.SubmitTopnStat, 0)
	statTime := r.start.Truncate(time.Hour)
	for _, s := range submits {
		submitStats = append(submitStats, store.SubmitTopnStat{
			StatTime:   statTime,
			AddressID:  s.SenderID,
			DataSize:   s.DataSize,
			StorageFee: s.StorageFee,
			Txs:        s.Txs,
			Files:      s.Files,
		})
	}

	return tsr.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(submitStats) > 0 {
			if err := tsr.DB.SubmitTopnStatStore.BatchDeltaUpsert(dbTx, submitStats); err != nil {
				return errors.WithMessage(err, "failed to batch delta upsert submits")
			}
		}
		if err := tsr.DB.ConfigStore.Upsert(dbTx, store.StatTopnSubmitTime, r.start.Format(time.RFC3339)); err != nil {
			return err
		}
		return nil
	})
}
