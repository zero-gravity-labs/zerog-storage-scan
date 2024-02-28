package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zero-gravity-labs/zerog-storage-scan/stat"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"gorm.io/gorm"
)

// TO add logic when refactor submit db domain
func dashboard(c *gin.Context) (interface{}, error) {
	storageBasicCost := StorageBasicCost{
		TokenInfo: *chargeToken,
	}
	result := Dashboard{
		StorageBasicCost: storageBasicCost,
	}

	return result, nil
}

// TO add logic when refactor submit db domain
func listTxStat(c *gin.Context) (interface{}, error) {
	return nil, nil
}

func listDataStat(c *gin.Context) (interface{}, error) {
	return queryStat(c, db.DB.Model(&store.SubmitStat{}), new([]store.SubmitStat))
}

// TO add logic when refactor submit db domain
func listBasicCostStat(c *gin.Context) (interface{}, error) {
	return nil, nil
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
