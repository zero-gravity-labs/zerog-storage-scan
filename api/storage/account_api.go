package storage

import (
	"github.com/0glabs/0g-storage-scan/api"
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
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
		return nil, api.ErrBlockchainRPC(err)
	}

	submitStat, err := db.AddressSubmitStore.Count(&addressInfo.addressId)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
	}

	rewardStat, err := db.AddressRewardStore.Count(&addressInfo.addressId)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
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
