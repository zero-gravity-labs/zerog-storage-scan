package store

import (
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	nhContract "github.com/zero-gravity-labs/zerog-storage-scan/contract"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Erc20Transfer struct {
	ID          uint64           `gorm:"primaryKey"`
	BlockNumber uint64           `gorm:"not null;index:idx_bn"`
	TxHash      string           `gorm:"type:varchar(64);not null;index:idx_hash,length:10"`
	Contract    string           `gorm:"-"`
	From        string           `gorm:"-"`
	To          string           `gorm:"-"`
	ContractId  uint64           `gorm:"not null;index:idx_contract_id"`
	FromId      uint64           `gorm:"not null;index:idx_from_id"`
	ToId        uint64           `gorm:"not null;index:idx_to_id"`
	Value       *decimal.Decimal `gorm:"type:varchar(78);not null"`
	CreatedAt   *time.Time       `gorm:"not null;index:idx_createdAt,sort:desc"`
}

func NewErc20Transfer(blockTime *time.Time, log *types.Log, filter *nhContract.Erc20TokenFilterer) (*Erc20Transfer, error) {
	transfer, err := filter.ParseTransfer(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	val := decimal.NewFromBigInt(transfer.Value, 0)
	erc20Transfer := &Erc20Transfer{
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash.String()[2:],
		Contract:    log.Address.String()[2:],
		From:        transfer.From.String()[2:],
		To:          transfer.To.String()[2:],
		Value:       &val,
		CreatedAt:   blockTime,
	}

	return erc20Transfer, nil
}

func (Erc20Transfer) TableName() string {
	return "erc20_transfers"
}

type Erc20TransferStore struct {
	*mysql.Store
}

func newErc20TransferStore(db *gorm.DB) *Erc20TransferStore {
	return &Erc20TransferStore{
		Store: mysql.NewStore(db),
	}
}

func (ets *Erc20TransferStore) Add(dbTx *gorm.DB, transfers []*Erc20Transfer) error {
	return dbTx.CreateInBatches(transfers, batchSizeInsert).Error
}

func (ets *Erc20TransferStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&Erc20Transfer{}).Error
}

func (ets *Erc20TransferStore) Sum(startTime, endTime *time.Time) (uint64, error) {
	if startTime == nil && endTime == nil {
		return 0, errors.New("At least provide one parameter for startTime and endTime")
	}

	db := ets.DB.Model(&Erc20Transfer{}).Select("IFNULL(sum(`value`), 0) as basic_cost")
	if startTime != nil && endTime != nil {
		db = db.Where("created_at >= ? and created_at < ?", startTime, endTime)
	}
	if startTime != nil && endTime == nil {
		db = db.Where("created_at >= ?", startTime)
	}
	if startTime == nil && endTime != nil {
		db = db.Where("created_at < ?", endTime)
	}

	var sum struct {
		BasicCost *decimal.Decimal
	}

	err := db.Find(&sum).Error
	if err != nil {
		return 0, err
	}

	return sum.BasicCost.BigInt().Uint64(), nil
}
