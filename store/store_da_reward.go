package store

import (
	"time"

	"github.com/ethereum/go-ethereum/common"

	nhContract "github.com/0glabs/0g-storage-scan/contract"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/openweb3/web3go/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DAReward struct {
	BlockNumber uint64 `gorm:"primaryKey;autoIncrement:false"`
	Epoch       uint64 `gorm:"primaryKey;autoIncrement:false"`
	QuorumID    uint64 `gorm:"primaryKey;autoIncrement:false"`
	Miner       string `gorm:"-"`
	MinerID     uint64 `gorm:"primaryKey;autoIncrement:false"`
	RootHash    string `gorm:"size:66;index:idx_root"`

	BlockTime time.Time `gorm:"not null"`
	TxHash    string    `gorm:"size:66;not null"`

	SampleRound  uint64          `gorm:"not null"`
	Quality      uint64          `gorm:"not null"`
	LineIndex    uint64          `gorm:"not null"`
	SubLineIndex uint64          `gorm:"not null"`
	Reward       decimal.Decimal `gorm:"type:decimal(65);not null"`
}

func NewDAReward(blockTime time.Time, log types.Log, filter *nhContract.DAEntranceFilterer) (*DAReward, error) {
	daReward, err := filter.ParseDAReward(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	reward := &DAReward{
		BlockNumber: log.BlockNumber,
		Epoch:       daReward.Epoch.Uint64(),
		QuorumID:    daReward.QuorumId.Uint64(),
		RootHash:    common.Hash(daReward.DataRoot[:]).String(),
		Miner:       daReward.Beneficiary.String(),

		BlockTime:    blockTime,
		TxHash:       log.TxHash.String(),
		SampleRound:  daReward.SampleRound.Uint64(),
		Quality:      daReward.Quality.Uint64(),
		LineIndex:    daReward.LineIndex.Uint64(),
		SubLineIndex: daReward.SublineIndex.Uint64(),
		Reward:       decimal.NewFromBigInt(daReward.Reward, 0),
	}

	return reward, nil
}

func (DAReward) TableName() string {
	return "da_rewards"
}

type DARewardStore struct {
	*mysql.Store
}

func newDARewardStore(db *gorm.DB) *DARewardStore {
	return &DARewardStore{
		Store: mysql.NewStore(db),
	}
}

func (rs *DARewardStore) Add(dbTx *gorm.DB, rewards []DAReward) error {
	return dbTx.CreateInBatches(rewards, batchSizeInsert).Error
}

func (rs *DARewardStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&DAReward{}).Error
}

func (rs *DARewardStore) List(idDesc bool, skip, limit int) (int64, []DAReward, error) {
	dbRaw := rs.DB.Model(&DAReward{})

	var orderBy string
	if idDesc {
		orderBy = "block_number DESC"
	} else {
		orderBy = "block_number ASC"
	}

	list := new([]DAReward)
	total, err := rs.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
