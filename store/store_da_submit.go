package store

import (
	"time"

	"github.com/0glabs/0g-storage-scan/contract"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DASubmit struct {
	BlockNumber uint64 `gorm:"primaryKey;autoIncrement:false"`
	Epoch       uint64 `gorm:"primaryKey;autoIncrement:false"`
	QuorumID    uint64 `gorm:"primaryKey;autoIncrement:false"`
	Sender      string `gorm:"-"`
	SenderID    uint64 `gorm:"primaryKey;autoIncrement:false"`
	RootHash    string `gorm:"size:66;index:idx_root"`

	BlockTime time.Time `gorm:"not null;index:idx_bt"`
	TxHash    string    `gorm:"size:66;not null;index:idx_txHash,length:10"`

	BlobPrice           decimal.Decimal `gorm:"type:decimal(65);not null"`
	Verified            bool            `gorm:"not null;default:false"`
	BlockNumberVerified *uint64         `gorm:"index:idx_bn_verified"`
	BlockTimeVerified   *time.Time      `gorm:"index:idx_bt_verified"`
	TxHashVerified      *string         `gorm:"size:66;index:idx_txHash_verified,length:10"`
}

func NewDASubmit(blockTime time.Time, log types.Log, filter *contract.DAEntranceFilterer) (*DASubmit, error) {
	dataUpload, err := filter.ParseDataUpload(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	submit := &DASubmit{
		Epoch:     dataUpload.Epoch.Uint64(),
		QuorumID:  dataUpload.QuorumId.Uint64(),
		RootHash:  common.Hash(dataUpload.DataRoot[:]).String(),
		Sender:    dataUpload.Sender.String(),
		BlobPrice: decimal.NewFromBigInt(dataUpload.BlobPrice, 0),

		BlockNumber: log.BlockNumber,
		BlockTime:   blockTime,
		TxHash:      log.TxHash.String(),
	}

	return submit, nil
}

func NewDASubmitVerified(blockTime time.Time, log types.Log, filter *contract.DAEntranceFilterer) (*DASubmit, error) {
	commitVerified, err := filter.ParseErasureCommitmentVerified(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	txHash := log.TxHash.String()
	submit := &DASubmit{
		Epoch:    commitVerified.Epoch.Uint64(),
		QuorumID: commitVerified.QuorumId.Uint64(),
		RootHash: common.BytesToHash(commitVerified.DataRoot[:]).String(),

		Verified:            true,
		BlockNumberVerified: &log.BlockNumber,
		BlockTimeVerified:   &blockTime,
		TxHashVerified:      &txHash,
	}

	return submit, nil
}

func (DASubmit) TableName() string {
	return "da_submits"
}

type DASubmitStore struct {
	*mysql.Store
}

func newDASubmitStore(db *gorm.DB) *DASubmitStore {
	return &DASubmitStore{
		Store: mysql.NewStore(db),
	}
}

func (ss *DASubmitStore) Add(dbTx *gorm.DB, submits []DASubmit) error {
	return dbTx.CreateInBatches(submits, batchSizeInsert).Error
}

func (ss *DASubmitStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&DASubmit{}).Error
}

func (ss *DASubmitStore) UpdateByPrimaryKey(dbTx *gorm.DB, s DASubmit) error {
	db := ss.DB
	if dbTx != nil {
		db = dbTx
	}

	if err := db.Model(&s).Where("epoch=? and quorum_id=? and root_hash=?", s.Epoch, s.QuorumID, s.RootHash).
		Updates(s).Error; err != nil {
		return err
	}

	return nil
}

func (ss *DASubmitStore) List(rootHash *string, txHash *string, idDesc bool, skip, limit int) (int64, []DASubmit, error) {
	dbRaw := ss.DB.Model(&DASubmit{})

	var conds []func(db *gorm.DB) *gorm.DB
	if rootHash != nil {
		conds = append(conds, RootHash(*rootHash))
	}
	if txHash != nil {
		conds = append(conds, TxHash(*txHash))
	}
	dbRaw.Scopes(conds...)

	var orderBy string
	if idDesc {
		orderBy = "block_number DESC, epoch DESC, quorum_id DESC"
	} else {
		orderBy = "block_number ASC, epoch ASC, quorum_id ASC"
	}

	list := new([]DASubmit)
	total, err := ss.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
