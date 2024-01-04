package stat

import (
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type StatBasicCost struct {
	*BaseStat
	statType string
}

func MustNewStatBasicCost(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client, startTime *time.Time) *AbsStat {
	baseStat := &BaseStat{
		Config:    cfg,
		Db:        db,
		Sdk:       sdk,
		StartTime: startTime,
	}

	statBasicCost := &StatBasicCost{
		BaseStat: baseStat,
		statType: baseStat.Config.MinStatIntervalDailySubmit,
	}

	return &AbsStat{
		Stat: statBasicCost,
		sdk:  baseStat.Sdk,
	}
}

func (sbc *StatBasicCost) nextTimeRange() (*TimeRange, error) {
	lastStat, err := sbc.Db.CostStatStore.LastByType(sbc.statType)
	if err != nil {
		return nil, err
	}

	var nextRangeStart *time.Time
	if lastStat == nil {
		nextRangeStart = sbc.StartTime
	} else {
		t := lastStat.StatTime.Add(Intervals[sbc.statType])
		nextRangeStart = &t
	}

	timeRange, err := sbc.calStatRange(nextRangeStart, Intervals[sbc.statType])
	if err != nil {
		return nil, err
	}

	return timeRange, nil
}

func (sbc *StatBasicCost) calculateStat(tr *TimeRange) error {
	stat, err := sbc.statBasicRange(tr)
	if err != nil {
		return err
	}
	hStat, err := sbc.statRange(tr.end, sbc.statType, Hour, stat)
	if err != nil {
		return err
	}
	dStat, err := sbc.statRange(tr.end, sbc.statType, Day, stat)
	if err != nil {
		return err
	}

	stats := []*store.CostStat{stat, hStat, dStat}
	return sbc.Db.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := sbc.Db.CostStatStore.Del(dbTx, hStat); err != nil {
			return errors.WithMessage(err, "failed to del hour stat")
		}
		if err := sbc.Db.CostStatStore.Del(dbTx, dStat); err != nil {
			return errors.WithMessage(err, "failed to del day stat")
		}
		if err := sbc.Db.CostStatStore.Add(dbTx, stats); err != nil {
			return errors.WithMessage(err, "failed to save stats")
		}
		return nil
	})
}

func (sbc *StatBasicCost) statBasicRange(tr *TimeRange) (*store.CostStat, error) {
	basicCost, err := sbc.Db.Erc20TransferStore.Sum(tr.start, tr.end)
	if err != nil {
		return nil, err
	}
	basicCostSum, err := sbc.Db.CostStatStore.Sum(nil, tr.start, sbc.statType)
	if err != nil {
		return nil, err
	}

	return &store.CostStat{
		StatTime:       tr.start,
		StatType:       sbc.statType,
		BasicCost:      basicCost,
		BasicCostTotal: basicCostSum + basicCost,
	}, nil
}

func (sbc *StatBasicCost) statRange(rangEnd *time.Time, srcStatType, descStatType string, latestStat *store.CostStat) (*store.CostStat, error) {
	rangeStart, err := sbc.calStatRangeStart(rangEnd, descStatType)
	if err != nil {
		return nil, err
	}

	basicCost, err := sbc.Db.CostStatStore.Sum(rangeStart, rangEnd, srcStatType)
	if err != nil {
		return nil, err
	}
	basicCostTotal, err := sbc.Db.CostStatStore.Sum(nil, rangeStart, descStatType)
	if err != nil {
		return nil, err
	}

	if latestStat != nil {
		basicCost += latestStat.BasicCost
	}

	return &store.CostStat{
		StatTime:       rangeStart,
		StatType:       descStatType,
		BasicCost:      basicCost,
		BasicCostTotal: basicCostTotal + basicCost,
	}, nil
}
