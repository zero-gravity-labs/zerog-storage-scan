package api

import (
	"encoding/json"
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/zero-gravity-labs/zerog-storage-scan/stat"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

func dashboard(c *gin.Context) (interface{}, error) {
	dataUplinkRate, exist, err := db.ConfigStore.Get(store.CfgDataUplinkRate)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
	}
	if !exist {
		return nil, ErrConfigNotFound
	}

	costStat, err := db.CostStatStore.LastByType(stat.Day)
	if err != nil {
		return nil, err
	}
	if costStat == nil {
		return nil, errors.New("Storage basic cost not stat.")
	}

	storageBasicCost := StorageBasicCost{
		TokenInfo:      *chargeToken,
		BasicCostTotal: strconv.FormatUint(costStat.BasicCostTotal, 10),
	}
	result := Dashboard{
		AverageUplinkRate: dataUplinkRate,
		StorageBasicCost:  storageBasicCost,
	}

	return result, nil
}

func listTxStat(c *gin.Context) (interface{}, error) {
	return queryStat(c, db.DB.Model(&store.TxStat{}), new([]store.TxStat))
}

func listDataStat(c *gin.Context) (interface{}, error) {
	return queryStat(c, db.DB.Model(&store.SubmitStat{}), new([]store.SubmitStat))
}

func listBasicCostStat(c *gin.Context) (interface{}, error) {
	return queryStat(c, db.DB.Model(&store.CostStat{}), new([]store.CostStat))
}

func queryStat(c *gin.Context, dbRaw *gorm.DB, records interface{}) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, err
	}

	r, _ := json.Marshal(statP)
	logrus.WithFields(logrus.Fields{
		"skip":         statP.Skip,
		"limit":        statP.Limit,
		"minTimestamp": statP.MinTimestamp,
		"maxTimestamp": statP.MaxTimestamp,
		"intervalType": statP.IntervalType,
		"sort":         statP.Sort,
	}).Infof("queryStat incoming %v", string(r))

	var conds []func(db *gorm.DB) *gorm.DB
	intervalType := stat.IntervalTypes[statP.IntervalType]
	conds = append(conds, StatType(intervalType))
	if statP.MinTimestamp != 0 {
		conds = append(conds, MinTimestamp(statP.MinTimestamp))
	}
	if statP.MaxTimestamp != 0 {
		conds = append(conds, MaxTimestamp(statP.MaxTimestamp))
	}
	dbRaw.Scopes(conds...)

	total, err := db.List(dbRaw, statP.isDesc(), statP.Skip, statP.Limit, records)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["total"] = total
	result["list"] = records
	return result, nil
}

func StatType(t string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stat_type = ?", t)
	}
}

func MinTimestamp(minTimestamp int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stat_time >= ?", minTimestamp)
	}
}

func MaxTimestamp(maxTimestamp int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stat_time <= ?", maxTimestamp)
	}
}
