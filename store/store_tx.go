package store

import (
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/openweb3/web3go/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type Tx struct {
	ID          uint64           `gorm:"primaryKey"`
	BlockNumber uint64           `gorm:"not null;index:idx_bn"`
	Hash        string           `gorm:"type:varchar(64);not null;index:idx_hash,length:10"`
	From        string           `gorm:"-"`
	FromId      uint64           `gorm:"not null"`
	To          string           `gorm:"-"`
	ToId        uint64           `gorm:"not null"`
	Nonce       uint64           `gorm:"not null"`
	MethodId    string           `gorm:"type:varchar(8);default:null"` // MethodId is function selector
	DripValue   *decimal.Decimal `gorm:"type:varchar(78);not null"`
	GasPrice    uint64           `gorm:"not null;default:0"`
	GasLimit    uint64           `gorm:"not null;default:0"`
	GasUsed     uint64           `gorm:"not null;default:0"`
	GasFee      uint64           `gorm:"not null;default:0"`
	Status      uint64           `gorm:"not null;default:0"`
	CreatedAt   *time.Time       `gorm:"not null;index:idx_createdAt,sort:desc"`
}

func NewTx(blockTime *time.Time, tx *types.TransactionDetail, rcpt *types.Receipt) *Tx {
	val := decimal.NewFromBigInt(tx.Value, 0)

	var gasUsed, gasFee uint64
	if rcpt != nil {
		gasUsed = rcpt.GasUsed
		gasFee = new(big.Int).Mul(
			new(big.Int).SetUint64(rcpt.EffectiveGasPrice),
			new(big.Int).SetUint64(rcpt.GasUsed),
		).Uint64()
	}

	return &Tx{
		BlockNumber: tx.BlockNumber.Uint64(),
		Hash:        tx.Hash.String()[2:],
		From:        tx.From.String()[2:],
		To:          tx.To.String()[2:],
		Nonce:       tx.Nonce,
		MethodId:    tx.Input.String()[2:10],
		DripValue:   &val,
		GasPrice:    tx.GasPrice.Uint64(),
		GasLimit:    tx.Gas,
		GasUsed:     gasUsed,
		GasFee:      gasFee,
		Status:      *tx.Status,
		CreatedAt:   blockTime,
	}
}

func (Tx) TableName() string {
	return "txs"
}

type TxStore struct {
	*mysql.Store
	as *AddressStore
}

func newTxStore(db *gorm.DB) *TxStore {
	return &TxStore{
		Store: mysql.NewStore(db),
		as:    newAddressStore(db),
	}
}

func (ts *TxStore) Add(dbTx *gorm.DB, txs []*Tx) error {
	return dbTx.CreateInBatches(txs, batchSizeInsert).Error
}

func (ts *TxStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&Tx{}).Error
}

func (ts *TxStore) Count(startTime, endTime *time.Time) (uint64, error) {
	var count int64
	ts.DB.Model(&Tx{}).Where("created_at >= ? and created_at < ? and status = ?", startTime, endTime, 1).
		Count(&count)
	return uint64(count), nil
}

// MapTxHashToTx TODO LRU cache
func (ts *TxStore) MapTxHashToTx(txHashes []string) (map[string]Tx, error) {
	txs := new([]Tx)
	err := ts.DB.Raw("select * from txs where hash in ?", txHashes).Scan(txs).Error
	if err != nil {
		return nil, err
	}

	m := make(map[string]Tx)
	for _, tx := range *txs {
		m[tx.Hash] = tx
	}

	return m, nil
}
