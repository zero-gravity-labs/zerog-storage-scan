package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AddressSubmit struct {
	SenderID        uint64 `gorm:"primary_key;autoIncrement:false"`
	SubmissionIndex uint64 `gorm:"primary_key;autoIncrement:false"`
	RootHash        string `gorm:"size:66;index:idx_root"`
	Length          uint64 `gorm:"not null"`

	BlockNumber uint64    `gorm:"not null"`
	BlockTime   time.Time `gorm:"not null"`
	TxHash      string    `gorm:"size:66;not null"`

	Status uint64          `gorm:"not null;default:0"`
	Fee    decimal.Decimal `gorm:"type:decimal(65);not null"`
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

func (ass *AddressSubmitStore) List(addressID *uint64, rootHash *string, idDesc bool, skip, limit int) (int64,
	[]AddressSubmit, error) {
	dbRaw := ass.DB.Model(&AddressSubmit{})

	var conds []func(db *gorm.DB) *gorm.DB
	if addressID != nil {
		conds = append(conds, SenderID(*addressID))
	}
	if rootHash != nil {
		conds = append(conds, RootHash(*rootHash))
	}
	dbRaw.Scopes(conds...)

	var orderBy string
	if idDesc {
		orderBy = "sender_id DESC"
	} else {
		orderBy = "sender_id ASC"
	}

	list := new([]AddressSubmit)
	total, err := ass.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
