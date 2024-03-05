package store

import (
	"time"

	"github.com/shopspring/decimal"
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
