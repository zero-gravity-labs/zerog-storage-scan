package da

import (
	scanApi "github.com/0glabs/0g-storage-scan/api"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func listDARewards(c *gin.Context) (interface{}, error) {
	var param PageParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, api.ErrValidation(errors.New("Invalid page param"))
	}

	total, rewards, err := listRewards(param.isDesc(), param.Skip, param.Limit)
	if err != nil {
		return nil, err
	}

	return convertDARewards(total, rewards)
}

func listRewards(idDesc bool, skip, limit int) (int64, []store.DAReward, error) {
	total, rewards, err := db.DARewardStore.List(idDesc, skip, limit)
	if err != nil {
		return 0, nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get da reward list"))
	}
	return total, rewards, nil
}

func convertDARewards(total int64, rewards []store.DAReward) (*RewardList, error) {
	addrIDs := make([]uint64, 0)
	for _, r := range rewards {
		addrIDs = append(addrIDs, r.MinerID)
	}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, scanApi.ErrBatchGetAddress(err)
	}

	daRewards := make([]Reward, 0)
	for _, r := range rewards {
		storageReward := Reward{
			Miner:       addrMap[r.MinerID].Address,
			Amount:      r.Reward,
			BlockNumber: r.BlockNumber,
			TxHash:      r.TxHash,
			Timestamp:   r.BlockTime.Unix(),
			SampleRound: r.SampleRound,
			Epoch:       r.Epoch,
			QuorumID:    r.QuorumID,
			RootHash:    r.RootHash,
		}
		daRewards = append(daRewards, storageReward)
	}

	return &RewardList{
		Total: total,
		List:  daRewards,
	}, nil
}
