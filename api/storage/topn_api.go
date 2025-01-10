package storage

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	dataSizeTopn   = "data_size"
	storageFeeTopn = "storage_fee"
	txsTopn        = "txs"
	filesTopn      = "files"

	maxRecords = 100
)

var (
	spanTypes = map[string]time.Duration{
		"24h": time.Hour * 24,
		"3d":  time.Hour * 24 * 3,
		"7d":  time.Hour * 24 * 7,
	}
)

func topnDataSize(c *gin.Context) (interface{}, error) {
	return topnByType(c, dataSizeTopn)
}

func topnStorageFee(c *gin.Context) (interface{}, error) {
	return topnByType(c, storageFeeTopn)
}

func topnTxs(c *gin.Context) (interface{}, error) {
	return topnByType(c, txsTopn)
}

func topnFiles(c *gin.Context) (interface{}, error) {
	return topnByType(c, filesTopn)
}

func topnByType(c *gin.Context, t string) (interface{}, error) {
	var topnP topnParam
	if err := c.ShouldBind(&topnP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid topn query param"))
	}

	statSpan := spanTypes[topnP.SpanType]
	records := cache.topnAddresses[t][statSpan]

	result := make(map[string]interface{})
	switch t {
	case dataSizeTopn:
		list := make([]DataTopn, 0)
		for rank, r := range records {
			list = append(list, DataTopn{
				Topn:     Topn{Rank: rank + 1, Address: r.Address},
				DataSize: r.DataSize,
			})
		}
		result["list"] = list
	case storageFeeTopn:
		list := make([]FeeTopn, 0)
		for rank, r := range records {
			list = append(list, FeeTopn{
				Topn:       Topn{Rank: rank + 1, Address: r.Address},
				StorageFee: r.StorageFee,
			})
		}
		result["list"] = list
	case txsTopn:
		list := make([]TxsTopn, 0)
		for rank, r := range records {
			list = append(list, TxsTopn{
				Topn: Topn{Rank: rank + 1, Address: r.Address},
				Txs:  r.Txs,
			})
		}
		result["list"] = list
	case filesTopn:
		list := make([]FilesTopn, 0)
		for rank, r := range records {
			list = append(list, FilesTopn{
				Topn:  Topn{Rank: rank + 1, Address: r.Address},
				Files: r.Files,
			})
		}
		result["list"] = list
	default:
		return nil, api.ErrValidation(errors.Errorf("Invalid topn type %v", t))
	}

	return result, nil
}

func topnReward(c *gin.Context) (interface{}, error) {
	var topnP topnParam
	if err := c.ShouldBind(&topnP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid topn query param"))
	}

	statSpan := spanTypes[topnP.SpanType]
	miners := cache.topnMiners[statSpan]
	if len(miners) == 0 {
		return map[string]interface{}{"list": []RewardTopn{}}, nil
	}

	list := make([]RewardTopn, 0)
	for rank, m := range miners {
		list = append(list, RewardTopn{
			Rank:    rank + 1,
			Address: m.Address,
			Amount:  m.Amount,
		})
	}

	return map[string]interface{}{"list": list}, nil
}
