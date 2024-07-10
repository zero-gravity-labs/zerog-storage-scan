package store

import (
	"strings"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	batchSizeInsert = 100
)

type MysqlStore struct {
	*mysql.Store
	*AddressStore
	*BlockStore
	*ConfigStore
	*SubmitStore
	*AddressSubmitStore
	*RewardStore
	*AddressRewardStore
	*SubmitStatStore
	*AddressStatStore
}

func MustNewStore(db *gorm.DB) *MysqlStore {
	return &MysqlStore{
		Store:              mysql.NewStore(db),
		AddressStore:       newAddressStore(db),
		BlockStore:         newBlockStore(db),
		ConfigStore:        newConfigStore(db),
		SubmitStore:        newSubmitStore(db),
		AddressSubmitStore: newAddressSubmitStore(db),
		RewardStore:        newRewardStore(db),
		AddressRewardStore: newAddressRewardStore(db),
		SubmitStatStore:    newSubmitStatStore(db),
		AddressStatStore:   newAddressStatStore(db),
	}
}

func (ms *MysqlStore) Push(block *Block, submits []*Submit, rewards []*Reward) error {
	addressSubmits := make([]AddressSubmit, 0)
	if len(submits) > 0 {
		for _, submit := range submits {
			addressSubmit := AddressSubmit{
				SenderID:        submit.SenderID,
				SubmissionIndex: submit.SubmissionIndex,
				RootHash:        submit.RootHash,
				Length:          submit.Length,
				BlockNumber:     submit.BlockNumber,
				BlockTime:       submit.BlockTime,
				TxHash:          submit.TxHash,
				Fee:             submit.Fee,
				TotalSegNum:     submit.TotalSegNum,
			}
			addressSubmits = append(addressSubmits, addressSubmit)
		}
	}

	addressRewards := make([]AddressReward, 0)
	if len(rewards) > 0 {
		for _, reward := range rewards {
			addressReward := AddressReward{
				MinerID:      reward.MinerID,
				PricingIndex: reward.PricingIndex,
				Amount:       reward.Amount,
				BlockNumber:  reward.BlockNumber,
				BlockTime:    reward.BlockTime,
				TxHash:       reward.TxHash,
			}
			addressRewards = append(addressRewards, addressReward)
		}
	}

	return ms.Store.DB.Transaction(func(dbTx *gorm.DB) error {
		// save blocks
		if err := ms.BlockStore.Add(dbTx, block); err != nil {
			return errors.WithMessage(err, "failed to save block")
		}

		// save flow submits
		if len(submits) > 0 {
			if err := ms.SubmitStore.Add(dbTx, submits); err != nil {
				return errors.WithMessage(err, "failed to save flow submits")
			}
			if err := ms.AddressSubmitStore.Add(dbTx, addressSubmits); err != nil {
				return errors.WithMessage(err, "failed to save address flow submits")
			}
		}

		// save distribute rewards
		if len(rewards) > 0 {
			if err := ms.RewardStore.Add(dbTx, rewards); err != nil {
				return errors.WithMessage(err, "failed to save rewards")
			}
			if err := ms.AddressRewardStore.Add(dbTx, addressRewards); err != nil {
				return errors.WithMessage(err, "failed to save address rewards")
			}
		}

		return nil
	})
}

func (ms *MysqlStore) Pop(block uint64) error {
	maxBlock, ok, err := ms.MaxBlock()
	if err != nil {
		return errors.WithMessage(err, "failed to get max block")
	}
	if !ok || block > maxBlock {
		return nil
	}

	return ms.Store.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := ms.BlockStore.Pop(dbTx, block); err != nil {
			return errors.WithMessage(err, "failed to remove blocks")
		}
		if err := ms.SubmitStore.Pop(dbTx, block); err != nil {
			return errors.WithMessage(err, "failed to remove submits")
		}
		if err := ms.AddressSubmitStore.Pop(dbTx, block); err != nil {
			return errors.WithMessage(err, "failed to remove address submits")
		}
		if err := ms.RewardStore.Pop(dbTx, block); err != nil {
			return errors.WithMessage(err, "failed to remove rewards")
		}
		if err := ms.AddressRewardStore.Pop(dbTx, block); err != nil {
			return errors.WithMessage(err, "failed to remove address rewards")
		}
		return nil
	})
}

func (ms *MysqlStore) Close() error {
	return ms.Store.Close()
}

func (ms *MysqlStore) UpdateSubmitByPrimaryKey(s *Submit, as *AddressSubmit) error {
	return ms.Store.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := ms.SubmitStore.UpdateByPrimaryKey(dbTx, s); err != nil {
			return errors.WithMessage(err, "failed to update submit")
		}
		if err := ms.AddressSubmitStore.UpdateByPrimaryKey(dbTx, as); err != nil {
			return errors.WithMessage(err, "failed to update address submit")
		}
		return nil
	})
}

func SenderID(si uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("sender_id = ?", si)
	}
}

func RootHash(rh string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("root_hash = ?", strings.ToLower(rh))
	}
}

func TxHash(rh string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("tx_hash = ?", strings.ToLower(rh))
	}
}

func StatType(t string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stat_type = ?", t)
	}
}

func MinTimestamp(minTimestamp int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stat_time >= ?", minTimestamp)
	}
}

func MaxTimestamp(maxTimestamp int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stat_time <= ?", maxTimestamp)
	}
}

func MinerID(mi uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("miner_id = ?", mi)
	}
}
