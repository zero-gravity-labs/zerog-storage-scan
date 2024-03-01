package store

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/zero-gravity-labs/zerog-storage-client/contract"
	"gorm.io/gorm"
)

type Submit struct {
	SubmissionIndex uint64 `gorm:"primaryKey;autoIncrement:false"`
	RootHash        string `gorm:"size:66;index:idx_root"`
	Sender          string `gorm:"-"`
	SenderID        uint64 `gorm:"not null"`
	Length          uint64 `gorm:"not null"`

	BlockNumber uint64    `gorm:"not null"`
	BlockTime   time.Time `gorm:"not null"`
	TxHash      string    `gorm:"size:66;not null"`

	Status uint64          `gorm:"not null;default:0"`
	Fee    decimal.Decimal `gorm:"type:decimal(65);not null"`

	Extra []byte `gorm:"type:mediumText"` // json field
}

type SubmitExtra struct {
	Identity   common.Hash         `json:"identity"`
	StartPos   *big.Int            `json:"startPos"`
	Submission contract.Submission `json:"submission"`
}

func NewSubmit(blockTime time.Time, log *types.Log, filter *contract.FlowFilterer) (*Submit, error) {
	flowSubmit, err := filter.ParseSubmit(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	extra, err := json.Marshal(SubmitExtra{
		Identity:   flowSubmit.Identity,
		StartPos:   flowSubmit.StartPos,
		Submission: flowSubmit.Submission,
	})
	if err != nil {
		return nil, err
	}

	submit := &Submit{
		SubmissionIndex: flowSubmit.SubmissionIndex.Uint64(),
		Sender:          flowSubmit.Sender.String(),
		Length:          flowSubmit.Submission.Length.Uint64(),
		BlockNumber:     log.BlockNumber,
		BlockTime:       blockTime,
		TxHash:          log.TxHash.String(),
		Fee:             decimal.NewFromBigInt(big.NewInt(0), 0),
		Extra:           extra,
	}

	return submit, nil
}

func (Submit) TableName() string {
	return "submits"
}

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

type SubmitStore struct {
	*mysql.Store
}

func newSubmitStore(db *gorm.DB) *SubmitStore {
	return &SubmitStore{
		Store: mysql.NewStore(db),
	}
}

func (ss *SubmitStore) Add(dbTx *gorm.DB, submits []*Submit) error {
	addressSubmits := make([]AddressSubmit, 0)
	for _, submit := range submits {
		addressSubmit := AddressSubmit{
			SenderID:        submit.SenderID,
			SubmissionIndex: submit.SubmissionIndex,
			RootHash:        submit.RootHash,
			Length:          submit.Length,
			BlockNumber:     submit.BlockNumber,
			BlockTime:       submit.BlockTime,
			TxHash:          submit.TxHash,
			Fee:             submit.Fee,
			Status:          submit.Status,
		}
		addressSubmits = append(addressSubmits, addressSubmit)
	}

	if err := dbTx.CreateInBatches(submits, batchSizeInsert).Error; err != nil {
		return err
	}

	return dbTx.CreateInBatches(addressSubmits, batchSizeInsert).Error
}

func (ss *SubmitStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&Submit{}).Error
}

func (ss *SubmitStore) Count(startTime, endTime time.Time) (*SubmitStatResult, error) {
	var result SubmitStatResult
	err := ss.DB.Model(&Submit{}).Select(`count(id) as file_count, IFNULL(sum(length), 0) as data_size, 
		IFNULL(sum(fee), 0) as base_fee`).Where("block_time >= ? and block_time < ?", startTime, endTime).
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ss *SubmitStore) FirstWithoutRootHash() (*Submit, error) {
	var submit Submit
	err := ss.DB.Where("root_hash = ?", "").First(&submit).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &submit, nil
}

func (ss *SubmitStore) UpdateByPrimaryKey(submit *Submit) error {
	if err := ss.DB.Model(&submit).Where("submission_index=?", submit.SubmissionIndex).
		Updates(submit).Error; err != nil {
		return err
	}

	return nil
}
