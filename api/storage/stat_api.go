package storage

import (
	"encoding/json"

	scanApi "github.com/0glabs/0g-storage-scan/api"
	"github.com/0glabs/0g-storage-scan/stat"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type statType int

const (
	storageStat statType = iota
	txStat
	feeStat
)

func listDataStat(c *gin.Context) (interface{}, error) {
	return getSubmitStatByType(c, storageStat)
}

func listTxStat(c *gin.Context) (interface{}, error) {
	return getSubmitStatByType(c, txStat)
}

func listFeeStat(c *gin.Context) (interface{}, error) {
	return getSubmitStatByType(c, feeStat)
}

func getSubmitStatByType(c *gin.Context, t statType) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid stat query param"))
	}

	total, records, err := db.SubmitStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get submit stat list"))
	}

	result := make(map[string]interface{})
	result["total"] = total

	switch t {
	case storageStat:
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
	case txStat:
		list := make([]TxStat, 0)
		for _, r := range records {
			list = append(list, TxStat{
				StatTime: r.StatTime,
				TxCount:  r.TxCount,
				TxTotal:  r.TxTotal,
			})
		}
		result["list"] = list
	case feeStat:
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
		return nil, api.ErrValidation(errors.Errorf("Invalid stat type %v", t))
	}

	return result, nil
}

func listAddressStat(c *gin.Context) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid stat query param"))
	}

	total, records, err := db.AddressStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get address stat"))
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]AddressStat, 0)
	for _, r := range records {
		list = append(list, AddressStat{
			StatTime:      r.StatTime,
			AddressNew:    r.AddrNew,
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
		return nil, api.ErrValidation(errors.Errorf("Invalid stat query param"))
	}

	total, records, err := db.MinerStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get miner stat"))
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]MinerStat, 0)
	for _, r := range records {
		list = append(list, MinerStat{
			StatTime:    r.StatTime,
			MinerNew:    r.MinerNew,
			MinerActive: r.MinerActive,
			MinerTotal:  r.MinerTotal,
		})
	}
	result["list"] = list

	return result, nil
}

func listRewardStat(c *gin.Context) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid stat query param"))
	}

	total, records, err := db.RewardStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get reward stat"))
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]RewardStat, 0)
	for _, r := range records {
		list = append(list, RewardStat{
			StatTime:    r.StatTime,
			RewardNew:   r.RewardNew,
			RewardTotal: r.RewardTotal,
		})
	}
	result["list"] = list

	return result, nil
}

func summary(_ *gin.Context) (interface{}, error) {
	value, exist, err := db.ConfigStore.Get(store.SyncStatusLog)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get log sync info"))
	}
	if !exist {
		return nil, api.ErrInternal(errors.New("No matching log-sync-info record found"))
	}

	var logSyncInfo stat.LogSyncInfo
	if err := json.Unmarshal([]byte(value), &logSyncInfo); err != nil {
		return nil, api.ErrInternal(errors.New("Failed to unmarshal log sync info"))
	}

	submitStat, err := db.SubmitStatStore.LastByType(store.Day)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get the latest submit stat"))
	}
	if submitStat == nil {
		return nil, api.ErrInternal(errors.New("No matching storage-fee-stat record found"))
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
