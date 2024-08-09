package storage

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/0glabs/0g-storage-scan/stat"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
)

type Type int

const (
	StorageStatType Type = iota
	TxStatType
	FeeStatType
)

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
		return nil, api.ErrValidation(errors.Errorf("Query param error"))
	}

	total, records, err := db.SubmitStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, api.ErrInternal(err)
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
				StatTime:        r.StatTime,
				StorageFee:      r.BaseFee,
				StorageFeeTotal: r.BaseFeeTotal,
			})
		}
		result["list"] = list
	default:
		return nil, commonApi.ErrValidation(errors.Errorf("Invalid stat type %v", t))
	}

	return result, nil
}

func listAddressStat(c *gin.Context) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Query param error"))
	}

	total, records, err := db.AddressStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, api.ErrInternal(err)
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]AddressStat, 0)
	for _, r := range records {
		list = append(list, AddressStat{
			StatTime:      r.StatTime,
			AddressNew:    r.AddrCount,
			AddressActive: r.AddrActive,
			AddressTotal:  r.AddrTotal,
		})
	}
	result["list"] = list

	return result, nil
}

func listMinerStat(c *gin.Context) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Query param error"))
	}

	total, records, err := db.MinerStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, api.ErrInternal(err)
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]MinerStat, 0)
	for _, r := range records {
		list = append(list, MinerStat{
			StatTime:    r.StatTime,
			MinerNew:    r.MinerCount,
			MinerActive: r.MinerActive,
			MinerTotal:  r.MinerTotal,
		})
	}
	result["list"] = list

	return result, nil
}

func summary(_ *gin.Context) (interface{}, error) {
	value, exist, err := db.ConfigStore.Get(store.KeyLogSyncInfo)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
	}
	if !exist {
		return nil, commonApi.ErrInternal(errors.Errorf("Log sync info not sync"))
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
		return nil, commonApi.ErrInternal(errors.Errorf("Storage base fee not stat"))
	}

	storageFee := StorageFeeStat{
		TokenInfo:       *chargeToken,
		StorageFeeTotal: submitStat.BaseFeeTotal,
	}
	result := Summary{
		StorageFeeStat: storageFee,
		LogSyncInfo:    logSyncInfo,
	}

	return result, nil
}
