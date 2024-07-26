package store

import (
	"time"

	"github.com/0glabs/0g-storage-scan/contract"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/openweb3/web3go/types"
	"gorm.io/gorm"
)

type DASigner struct {
	SignerID uint64 `gorm:"primaryKey;autoIncrement:false"`
	Address  string `gorm:"_"`
	Socket   string `gorm:"size:128"`

	BlockNumber uint64    `gorm:"not null;index:idx_bn"`
	BlockTime   time.Time `gorm:"not null;index:idx_bt"`
	TxHash      string    `gorm:"size:66;not null;index:idx_txHash,length:10"`
}

func NewDASigner(blockTime time.Time, signerLog types.Log, filter *contract.DASignersFilterer) (*DASigner, error) {
	newSigner, err := filter.ParseNewSigner(*signerLog.ToEthLog())
	if err != nil {
		return nil, err
	}

	signer := &DASigner{
		Address:     newSigner.Signer.String(),
		BlockNumber: signerLog.BlockNumber,
		BlockTime:   blockTime,
		TxHash:      signerLog.TxHash.String(),
	}

	return signer, nil
}

func NewDASignerSocket(signerLog types.Log, filter *contract.DASignersFilterer) (*DASigner, error) {
	socketUpdated, err := filter.ParseSocketUpdated(*signerLog.ToEthLog())
	if err != nil {
		return nil, err
	}

	signer := &DASigner{
		Address: socketUpdated.Signer.String(),
		Socket:  socketUpdated.Socket,
	}

	return signer, nil
}

func (DASigner) TableName() string {
	return "da_signers"
}

type DASignerStore struct {
	*mysql.Store
}

func newDASignerStore(db *gorm.DB) *DASignerStore {
	return &DASignerStore{
		Store: mysql.NewStore(db),
	}
}

func (ss *DASignerStore) Add(dbTx *gorm.DB, signers []DASigner) error {
	return dbTx.CreateInBatches(signers, batchSizeInsert).Error
}

func (ss *DASignerStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&DASigner{}).Error
}

func (ss *DASignerStore) UpdateByPrimaryKey(dbTx *gorm.DB, s DASigner) error {
	db := ss.DB
	if dbTx != nil {
		db = dbTx
	}

	if err := db.Model(&s).Where("signer_id=?", s.SignerID).Updates(s).Error; err != nil {
		return err
	}

	return nil
}
