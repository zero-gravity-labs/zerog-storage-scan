package storage

import (
	scanApi "github.com/0glabs/0g-storage-scan/api"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func listMiners(c *gin.Context) (interface{}, error) {
	var param PageParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, api.ErrValidation(errors.New("Invalid page param"))
	}

	total, miners, err := db.MinerStore.List(param.isDesc(), param.Skip, param.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get miner list"))
	}

	return convertMiners(total, miners)
}

func convertMiners(total int64, miners []store.Miner) (*MinerList, error) {
	addrIDs := make([]uint64, 0)
	for _, m := range miners {
		addrIDs = append(addrIDs, m.ID)
	}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, scanApi.ErrBatchGetAddress(err)
	}

	minerList := make([]Miner, 0)
	for _, m := range miners {
		miner := Miner{
			Miner:       addrMap[m.ID].Address,
			TotalReward: m.Amount,
			Timestamp:   m.UpdatedAt.Unix(),
		}
		minerList = append(minerList, miner)
	}

	return &MinerList{
		Total: total,
		List:  minerList,
	}, nil
}

func getMinerInfo(c *gin.Context) (interface{}, error) {
	addr, err := getAddressInfo(c)
	if err != nil {
		return nil, err
	}

	balance, err := sdk.Eth.Balance(common.HexToAddress(addr.Address), nil)
	if err != nil {
		return nil, scanApi.ErrBlockchainRPC(errors.WithMessagef(err, "Failed to get balance %v", addr.Address))
	}

	var amount decimal.Decimal
	var miner store.Miner
	exist, err := db.MinerStore.GetById(&miner, addr.ID)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessagef(err, "Failed to get miner info %v", addr.Address))
	}
	if exist {
		amount = miner.Amount
	}

	minerInfo := MinerInfo{
		Balance:     decimal.NewFromBigInt(balance, 0),
		TotalReward: amount,
	}

	return minerInfo, nil
}
