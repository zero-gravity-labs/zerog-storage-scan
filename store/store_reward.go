package store

import (
	"time"

	nhContract "github.com/0glabs/0g-storage-scan/contract"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/openweb3/web3go/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Reward struct {
	BlockNumber  uint64          `gorm:"primaryKey;autoIncrement:false"`
	BlockTime    time.Time       `gorm:"not null"`
	TxHash       string          `gorm:"size:66;not null"`
	Miner        string          `gorm:"-"`
	MinerID      uint64          `gorm:"not null"`
	PricingIndex uint64          `gorm:"not null"`
	Amount       decimal.Decimal `gorm:"type:decimal(65);not null"`
}

func NewReward(blockTime time.Time, log types.Log, filter *nhContract.OnePoolRewardFilterer) (*Reward, error) {
	distributeReward, err := filter.ParseDistributeReward(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	reward := &Reward{
		PricingIndex: distributeReward.PricingIndex.Uint64(),
		Miner:        distributeReward.Beneficiary.String(),
		Amount:       decimal.NewFromBigInt(distributeReward.Amount, 0),
		BlockNumber:  log.BlockNumber,
		BlockTime:    blockTime,
		TxHash:       log.TxHash.String(),
	}

	return reward, nil
}

func (Reward) TableName() string {
	return "rewards"
}

type RewardStore struct {
	*mysql.Store
}

func newRewardStore(db *gorm.DB) *RewardStore {
	return &RewardStore{
		Store: mysql.NewStore(db),
	}
}

func (rs *RewardStore) Add(dbTx *gorm.DB, rewards []*Reward) error {
	return dbTx.CreateInBatches(rewards, batchSizeInsert).Error
}

func (rs *RewardStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&Reward{}).Error
}

func (rs *RewardStore) List(idDesc bool, skip, limit int) (int64, []Reward, error) {
	dbRaw := rs.DB.Model(&Reward{})

	var orderBy string
	if idDesc {
		orderBy = "block_number DESC"
	} else {
		orderBy = "block_number ASC"
	}

	list := new([]Reward)
	total, err := rs.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
