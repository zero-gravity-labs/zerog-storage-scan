package store

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/0glabs/0g-storage-client/contract"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Status int

const (
	NotFinalized Status = iota
	Finalized
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

func NewSubmit(blockTime time.Time, log types.Log, filter *contract.FlowFilterer) (*Submit, error) {
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
		RootHash:        flowSubmit.Submission.Root().String(),
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
	err := ss.DB.Model(&Submit{}).Select(`count(submission_index) as file_count, 
		IFNULL(sum(length), 0) as data_size, IFNULL(sum(fee), 0) as base_fee, count(distinct tx_hash) as tx_count`).
		Where("block_time >= ? and block_time < ?", startTime, endTime).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ss *SubmitStore) UpdateByPrimaryKey(dbTx *gorm.DB, s *Submit) error {
	db := ss.DB
	if dbTx != nil {
		db = dbTx
	}

	if err := db.Model(&s).Where("submission_index=?", s.SubmissionIndex).
		Updates(s).Error; err != nil {
		return err
	}

	return nil
}

func (ss *SubmitStore) List(rootHash *string, idDesc bool, skip, limit int) (int64, []Submit, error) {
	dbRaw := ss.DB.Model(&Submit{})
	var conds []func(db *gorm.DB) *gorm.DB
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

	list := new([]Submit)
	total, err := ss.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}

func (ss *SubmitStore) BatchGetNotFinalized(batch int) ([]Submit, error) {
	submits := new([]Submit)
	if err := ss.DB.Raw("select submission_index, sender_id from submits where status = ? limit ?",
		NotFinalized, batch).Scan(submits).Error; err != nil {
		return nil, err
	}

	return *submits, nil
}
