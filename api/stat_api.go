package api

import (
	"encoding/json"
	"strconv"

	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zero-gravity-labs/zerog-storage-scan/stat"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"gorm.io/gorm"
)

type Type int

const (
	StorageStatType Type = iota
	TxStatType
	FeeStatType
)

func dashboard(c *gin.Context) (interface{}, error) {
	submitStat, err := db.SubmitStatStore.LastByType(stat.Day)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
	}
	if submitStat == nil {
		return nil, ErrStorageBaseFeeNotStat
	}

	storageBasicCost := StorageBasicCost{
		TokenInfo:      *chargeToken,
		BasicCostTotal: strconv.FormatUint(submitStat.BaseFeeTotal, 10),
	}
	result := Dashboard{
		StorageBasicCost: storageBasicCost,
	}

	return result, nil
}

func listDataStat(c *gin.Context) (interface{}, error) {
	return getSubmitStatByType(c, StorageStatType)
}

func listTxStat(c *gin.Context) (interface{}, error) {
	return getSubmitStatByType(c, TxStatType)
}

func listFeeStat(c *gin.Context) (interface{}, error) {
	return getSubmitStatByType(c, FeeStatType)
}

func getSubmitStatByType(c *gin.Context, t Type) (interface{}, error) {
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
	dbRaw := db.DB.Model(&store.SubmitStat{})
	dbRaw.Scopes(conds...)

	records := new([]store.SubmitStat)
	total, err := db.List(dbRaw, statP.isDesc(), statP.Skip, statP.Limit, records)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["total"] = total

	switch t {
	case StorageStatType:
		list := make([]DataStat, 0)
		for _, stat := range *records {
			list = append(list, DataStat{
				StatTime:  &stat.StatTime,
				FileCount: stat.FileCount,
				FileTotal: stat.FileTotal,
				DataSize:  stat.DataSize,
				DataTotal: stat.DataTotal,
			})
		}
		result["list"] = list
	case TxStatType:
		list := make([]TxStat, 0)
		for _, stat := range *records {
			list = append(list, TxStat{
				StatTime: &stat.StatTime,
				TxCount:  stat.FileCount,
				TxTotal:  stat.FileTotal,
			})
		}
		result["list"] = list
	case FeeStatType:
		list := make([]FeeStat, 0)
		for _, stat := range *records {
			list = append(list, FeeStat{
				StatTime:     &stat.StatTime,
				BaseFee:      stat.BaseFee,
				BaseFeeTotal: stat.BaseFeeTotal,
			})
		}
		result["list"] = list
	default:
		return nil, ErrStatTypeNotSupported
	}

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
