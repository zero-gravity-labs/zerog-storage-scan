package store

import (
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
	*SubmitStatStore
}

func MustNewStore(db *gorm.DB) *MysqlStore {
	return &MysqlStore{
		Store:           mysql.NewStore(db),
		AddressStore:    newAddressStore(db),
		BlockStore:      newBlockStore(db),
		ConfigStore:     newConfigStore(db),
		SubmitStore:     newSubmitStore(db),
		SubmitStatStore: newSubmitStatStore(db),
	}
}

func (ms *MysqlStore) Push(block *Block, submits []*Submit) error {
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
			return errors.WithMessage(err, "failed to remove block")
		}
		if err := ms.SubmitStore.Pop(dbTx, block); err != nil {
			return errors.WithMessage(err, "failed to remove flow submits")
		}
		return nil
	})
}

func (ms *MysqlStore) Close() error {
	return ms.Store.Close()
}
