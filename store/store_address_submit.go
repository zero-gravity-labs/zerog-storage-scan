package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AddressSubmit struct {
	SenderID        uint64 `gorm:"primaryKey;autoIncrement:false"`
	SubmissionIndex uint64 `gorm:"primaryKey;autoIncrement:false"`
	RootHash        string `gorm:"size:66;index:idx_root"`
	Length          uint64 `gorm:"not null"`

	BlockNumber uint64    `gorm:"not null;index:idx_bn"`
	BlockTime   time.Time `gorm:"not null"`
	TxHash      string    `gorm:"size:66;not null"`

	TotalSegNum    uint64          `gorm:"not null;default:0"`
	UploadedSegNum uint64          `gorm:"not null;default:0"`
	Status         uint8           `gorm:"not null;default:0"`
	Fee            decimal.Decimal `gorm:"type:decimal(65);not null"`
}

func (AddressSubmit) TableName() string {
	return "address_submits"
}

type AddressSubmitStore struct {
	*mysql.Store
}

func newAddressSubmitStore(db *gorm.DB) *AddressSubmitStore {
	return &AddressSubmitStore{
		Store: mysql.NewStore(db),
	}
}

func (ass *AddressSubmitStore) Add(dbTx *gorm.DB, addressSubmits []AddressSubmit) error {
	return dbTx.CreateInBatches(addressSubmits, batchSizeInsert).Error
}

func (ass *AddressSubmitStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&AddressSubmit{}).Error
}

func (ass *AddressSubmitStore) UpdateByPrimaryKey(dbTx *gorm.DB, s *AddressSubmit) error {
	db := ass.DB
	if dbTx != nil {
		db = dbTx
	}

	if err := db.Model(&s).Where("sender_id=? and submission_index=?", s.SenderID, s.SubmissionIndex).
		Updates(s).Error; err != nil {
		return err
	}

	return nil
}

func (ass *AddressSubmitStore) List(addressID *uint64, rootHash *string, idDesc bool, skip, limit int) (int64,
	[]AddressSubmit, error) {
	if addressID == nil {
		return 0, nil, errors.New("nil addressID")
	}

	dbRaw := ass.DB.Model(&AddressSubmit{})
	var conds []func(db *gorm.DB) *gorm.DB
	conds = append(conds, SenderID(*addressID))
	if rootHash != nil {
		conds = append(conds, RootHash(*rootHash))
	}
	dbRaw.Scopes(conds...)

	var orderBy string
	if idDesc {
		orderBy = "submission_index DESC"
	} else {
		orderBy = "submission_index ASC"
	}

	list := new([]AddressSubmit)
	total, err := ass.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}

func (ass *AddressSubmitStore) Count(addressID *uint64) (*SubmitStatResult, error) {
	if addressID == nil {
		return nil, errors.New("nil addressID")
	}

	var result SubmitStatResult
	err := ass.DB.Model(&AddressSubmit{}).Select(`count(submission_index) as file_count, 
		IFNULL(sum(length), 0) as data_size, IFNULL(sum(fee), 0) as base_fee, count(distinct tx_hash) as tx_count`).
		Where("sender_id = ?", addressID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
