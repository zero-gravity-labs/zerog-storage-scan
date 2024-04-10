package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AddressReward struct {
	MinerID      uint64          `gorm:"primaryKey;autoIncrement:false"`
	BlockNumber  uint64          `gorm:"primaryKey;autoIncrement:false;index:idx_bn"`
	BlockTime    time.Time       `gorm:"not null"`
	TxHash       string          `gorm:"size:66;not null"`
	PricingIndex uint64          `gorm:"not null"`
	Amount       decimal.Decimal `gorm:"type:decimal(65);not null"`
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

func (ars *AddressRewardStore) List(addressID *uint64, idDesc bool, skip, limit int) (int64, []AddressReward, error) {
	if addressID == nil {
		return 0, nil, errors.New("nil addressID")
	}

	dbRaw := ars.DB.Model(&AddressReward{})
	var conds []func(db *gorm.DB) *gorm.DB
	conds = append(conds, MinerID(*addressID))
	dbRaw.Scopes(conds...)

	var orderBy string
	if idDesc {
		orderBy = "pricing_index DESC"
	} else {
		orderBy = "pricing_index ASC"
	}

	list := new([]AddressReward)
	total, err := ars.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}

type RewardStatResult struct {
	RewardCount  uint64
	RewardAmount decimal.Decimal
}

func (ars *AddressRewardStore) Count(addressID *uint64) (*RewardStatResult, error) {
	if addressID == nil {
		return nil, errors.New("nil addressID")
	}

	var result RewardStatResult
	err := ars.DB.Model(&AddressReward{}).
		Select(`count(*) as reward_count,IFNULL(sum(amount), 0) as reward_amount`).
		Where("miner_id = ?", addressID).
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
