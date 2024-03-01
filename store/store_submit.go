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

// TODO one table for overall submissions and another table for submissions of specific account.

type Submit struct {
	ID          uint64    `gorm:"index:idx_sender_id,priority:2"`
	BlockNumber uint64    `gorm:"not null"`
	BlockTime   time.Time `gorm:"not null"`
	TxHash      string    `gorm:"size:66;not null"`

	// TODO use as primary key?
	SubmissionIndex uint64 `gorm:"not null;index:idx_seq"`
	RootHash        string `gorm:"size:66;index:idx_root"`
	Sender          string `gorm:"-"`
	SenderID        uint64 `gorm:"not null;index:idx_sender_id,priority:1"`
	Length          uint64 `gorm:"not null"`
	// TODO supports more status, including L1 and L2 status
	Finalized bool `gorm:"default:false"`
	// TODO change to Fee. how about use decimal(64,18) as sql type?
	Value decimal.Decimal `gorm:"type:varchar(78);not null"`

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
		BlockNumber:     log.BlockNumber,
		BlockTime:       blockTime,
		TxHash:          log.TxHash.String(),
		SubmissionIndex: flowSubmit.SubmissionIndex.Uint64(),
		Sender:          flowSubmit.Sender.String(),
		Length:          flowSubmit.Submission.Length.Uint64(),
		Value:           decimal.NewFromBigInt(big.NewInt(0), 0),
		Extra:           extra,
	}

	return submit, nil
}

func (Submit) TableName() string {
	return "submits"
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
	return dbTx.CreateInBatches(submits, batchSizeInsert).Error
}

func (ss *SubmitStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&Submit{}).Error
}

func (ss *SubmitStore) Count(startTime, endTime time.Time) (*SubmitStatResult, error) {
	var result SubmitStatResult
	err := ss.DB.Model(&Submit{}).Select(`count(id) as file_count, IFNULL(sum(length), 0) as data_size, 
		IFNULL(sum(value), 0) as base_fee`).Where("block_time >= ? and block_time < ?", startTime, endTime).
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

func (ss *SubmitStore) Update(submit *Submit) error {
	if err := ss.DB.Model(&submit).Updates(submit).Error; err != nil {
		return err
	}

	return nil
}
