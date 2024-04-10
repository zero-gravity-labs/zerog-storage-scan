package api

import (
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/gin-gonic/gin"
)

func listStorageRewards(c *gin.Context) (interface{}, error) {
	var param PageParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	total, rewards, err := listRewards(nil, param.isDesc(), param.Skip, param.Limit)
	if err != nil {
		return nil, err
	}

	return convertStorageRewards(total, rewards)
}
func listAddressStorageRewards(c *gin.Context) (interface{}, error) {
	addressInfo, err := getAddressInfo(c)
	if err != nil {
		return nil, err
	}
	addrIDPtr := &addressInfo.addressId

	var param PageParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	total, rewards, err := listRewards(addrIDPtr, param.isDesc(), param.Skip, param.Limit)
	if err != nil {
		return nil, err
	}

	return convertStorageRewards(total, rewards)
}

func listRewards(addressID *uint64, idDesc bool, skip, limit int) (int64, []store.Reward, error) {
	if addressID == nil {
		return db.RewardStore.List(idDesc, skip, limit)
	}

	total, addrRewards, err := db.AddressRewardStore.List(addressID, idDesc, skip, limit)
	if err != nil {
		return 0, nil, err
	}

	rewards := make([]store.Reward, 0)
	for _, ar := range addrRewards {
		rewards = append(rewards, store.Reward{
			PricingIndex: ar.PricingIndex,
			MinerID:      ar.MinerID,
			Amount:       ar.Amount,
			BlockNumber:  ar.BlockNumber,
			BlockTime:    ar.BlockTime,
			TxHash:       ar.TxHash,
		})
	}

	return total, rewards, nil
}

func convertStorageRewards(total int64, rewards []store.Reward) (*RewardList, error) {
	addrIDs := make([]uint64, 0)
	for _, r := range rewards {
		addrIDs = append(addrIDs, r.MinerID)
	}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, err
	}

	storageRewards := make([]Reward, 0)
	for _, r := range rewards {
		storageReward := Reward{
			Miner:       addrMap[r.MinerID].Address,
			Amount:      r.Amount,
			BlockNumber: r.BlockNumber,
			TxHash:      r.TxHash,
			Timestamp:   r.BlockTime.Unix(),
		}
		storageRewards = append(storageRewards, storageReward)
	}

	return &RewardList{
		Total: total,
		List:  storageRewards,
	}, nil
}
