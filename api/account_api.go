package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func getAccountInfo(c *gin.Context) (interface{}, error) {
	addressInfo, err := getAddressInfo(c)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(addressInfo.address)
	balance, err := sdk.Eth.Balance(addr, nil)
	if err != nil {
		logrus.WithError(err).WithField("address", addressInfo.address).Error("Failed to get balance")
		return nil, errors.Errorf("Get balance error, address %v", addressInfo.address)
	}

	submitStat, err := db.AddressSubmitStore.Count(&addressInfo.addressId)
	if err != nil {
		return nil, err
	}

	rewardStat, err := db.AddressRewardStore.Count(&addressInfo.addressId)
	if err != nil {
		return nil, err
	}

	accountInfo := AccountInfo{
		Balance:    decimal.NewFromBigInt(balance, 0),
		FileCount:  submitStat.FileCount,
		TxCount:    submitStat.TxCount,
		DataSize:   submitStat.DataSize,
		StorageFee: submitStat.BaseFee,
	}

	if accountInfo.TxCount > 0 {
		accountInfo.TxTab = 1
	}
	if rewardStat.RewardCount > 0 {
		accountInfo.RewardTab = 1
	}

	return accountInfo, nil
}
