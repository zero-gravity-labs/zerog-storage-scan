package storage

import (
	scanApi "github.com/0glabs/0g-storage-scan/api"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func getAccountInfo(c *gin.Context) (interface{}, error) {
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

	accountInfo := AccountInfo{
		Balance:      decimal.NewFromBigInt(balance, 0),
		DataSize:     addr.DataSize,
		StorageFee:   addr.StorageFee,
		Txs:          addr.Txs,
		Files:        addr.Files,
		ExpiredFiles: addr.ExpiredFiles,
		PrunedFiles:  addr.PrunedFiles,
		TotalReward:  amount,
	}

	return accountInfo, nil
}
