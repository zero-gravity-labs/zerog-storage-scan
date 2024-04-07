package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AddressReward struct {
	MinerID      uint64          `gorm:"primaryKey;autoIncrement:false"`
	PricingIndex uint64          `gorm:"primaryKey;autoIncrement:false"`
	Amount       decimal.Decimal `gorm:"type:decimal(65);not null"`
	BlockNumber  uint64          `gorm:"not null;index:idx_bn"`
	BlockTime    time.Time       `gorm:"not null"`
	TxHash       string          `gorm:"size:66;not null"`
}

func (AddressReward) TableName() string {
	return "address_rewards"
}

type AddressRewardStore struct {
	*mysql.Store
}

func newAddressRewardStore(db *gorm.DB) *AddressRewardStore {
	return &AddressRewardStore{
		Store: mysql.NewStore(db),
	}
}

func (ars *AddressRewardStore) Add(dbTx *gorm.DB, addressRewards []AddressReward) error {
	return dbTx.CreateInBatches(addressRewards, batchSizeInsert).Error
}

func (ars *AddressRewardStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&AddressReward{}).Error
}
