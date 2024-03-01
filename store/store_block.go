package store

import (
	"database/sql"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/openweb3/web3go/types"
	"gorm.io/gorm"
)

type Block struct {
	BlockNumber uint64    `gorm:"primaryKey;autoIncrement:false"`
	Hash        string    `gorm:"size:66;not null"`
	BlockTime   time.Time `gorm:"not null;index:idx_block_time"`
}

func NewBlock(data *types.Block) *Block {
	blockTime := time.Unix(int64(data.Timestamp), 0)
	return &Block{
		BlockNumber: data.Number.Uint64(),
		Hash:        data.Hash.String(),
		BlockTime:   blockTime,
	}
}

func (Block) TableName() string {
	return "blocks"
}

type BlockStore struct {
	*mysql.Store
}

func newBlockStore(db *gorm.DB) *BlockStore {
	return &BlockStore{
		Store: mysql.NewStore(db),
	}
}

func (bs *BlockStore) Add(dbTx *gorm.DB, block *Block) error {
	return dbTx.Create(block).Error
}

func (bs *BlockStore) MaxBlock() (uint64, bool, error) {
	var maxBlock sql.NullInt64

	db := bs.Store.DB.Model(&Block{}).Select("MAX(block_number)")
	if err := db.Find(&maxBlock).Error; err != nil {
		return 0, false, err
	}

	if !maxBlock.Valid {
		return 0, false, nil
	}

	return uint64(maxBlock.Int64), true, nil
}

func (bs *BlockStore) BlockHash(blockNumber uint64) (string, bool, error) {
	var blk Block

	existed, err := bs.Store.Exists(&blk, "block_number = ?", blockNumber)
	if err != nil {
		return "", false, err
	}

	return blk.Hash, existed, nil
}

func (bs *BlockStore) FirstBlockAfterTime(t *time.Time) (uint64, bool, error) {
	var blk Block

	result := bs.DB.Where("block_time >= ?", t).Order("block_time asc").Limit(1).Find(&blk)
	if result.Error != nil {
		return 0, false, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, false, nil
	}

	return blk.BlockNumber, true, nil
}

func (bs *BlockStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&Block{}).Error
}
