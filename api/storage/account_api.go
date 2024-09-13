package storage

import (
	"github.com/0glabs/0g-storage-scan/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func getAccountInfo(c *gin.Context) (interface{}, error) {
	addressInfo, err := getAddressInfo(c)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(addressInfo.address)
	balance, err := sdk.Eth.Balance(addr, nil)
	if err != nil {
		return nil, api.ErrBlockchainRPC(errors.WithMessagef(err, "Failed to get balance, address %v", addressInfo.address))
	}

	submitStat, err := db.AddressSubmitStore.Count(&addressInfo.addressId)
	if err != nil {
		return nil, api.ErrDatabase(errors.WithMessage(err, "Failed to get submit stat"))
	}

	rewardStat, err := db.AddressRewardStore.Count(&addressInfo.addressId)
	if err != nil {
		return nil, api.ErrDatabase(errors.WithMessage(err, "Failed to get reward stat"))
	}

	accountInfo := AccountInfo{
		Balance:     decimal.NewFromBigInt(balance, 0),
		FileCount:   submitStat.FileCount,
		TxCount:     submitStat.TxCount,
		DataSize:    submitStat.DataSize,
		StorageFee:  submitStat.BaseFee,
		RewardCount: rewardStat.RewardCount,
	}

	return accountInfo, nil
}
