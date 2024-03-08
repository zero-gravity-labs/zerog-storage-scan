package api

import (
	"encoding/json"

	"github.com/0glabs/0g-storage-scan/stat"
	"github.com/0glabs/0g-storage-scan/store"
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
)

type Type int

const (
	StorageStatType Type = iota
	TxStatType
	FeeStatType
)

func dashboard(_ *gin.Context) (interface{}, error) {
	value, exist, err := db.ConfigStore.Get(store.KeyLogSyncInfo)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
	}
	if !exist {
		return nil, ErrConfigNotFound
	}

	var logSyncInfo stat.LogSyncInfo
	if err := json.Unmarshal([]byte(value), &logSyncInfo); err != nil {
		return nil, commonApi.ErrInternal(err)
	}

	submitStat, err := db.SubmitStatStore.LastByType(store.Day)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
	}
	if submitStat == nil {
		return nil, ErrStorageBaseFeeNotStat
	}

	storageBasicCost := StorageBasicCost{
		TokenInfo:      *chargeToken,
		BasicCostTotal: submitStat.BaseFeeTotal,
	}
	result := Dashboard{
		StorageBasicCost: storageBasicCost,
		LogSyncInfo:      logSyncInfo,
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

	total, records, err := db.SubmitStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["total"] = total

	switch t {
	case StorageStatType:
		list := make([]DataStat, 0)
		for _, r := range records {
			list = append(list, DataStat{
				StatTime:  r.StatTime,
				FileCount: r.FileCount,
				FileTotal: r.FileTotal,
				DataSize:  r.DataSize,
				DataTotal: r.DataTotal,
			})
		}
		result["list"] = list
	case TxStatType:
		list := make([]TxStat, 0)
		for _, r := range records {
			list = append(list, TxStat{
				StatTime: r.StatTime,
				TxCount:  r.TxCount,
				TxTotal:  r.TxTotal,
			})
		}
		result["list"] = list
	case FeeStatType:
		list := make([]FeeStat, 0)
		for _, r := range records {
			list = append(list, FeeStat{
				StatTime:     r.StatTime,
				BaseFee:      r.BaseFee,
				BaseFeeTotal: r.BaseFeeTotal,
			})
		}
		result["list"] = list
	default:
		return nil, ErrStatTypeNotSupported
	}

	return result, nil
}
