package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DASubmitStat struct {
	ID       uint64    `json:"-"`
	StatType string    `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime time.Time `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`

	BlobNew         uint64          `gorm:"not null;default:0" json:"blobNew"`                          // Number of blobs in a specific time interval
	BlobTotal       uint64          `gorm:"not null;default:0" json:"blobTotal"`                        // Total number of blobs by a certain time
	DataSizeNew     uint64          `gorm:"type:decimal(65);not null;default:0" json:"dataSizeNew"`     // The data size of storage
	DataSizeTotal   uint64          `gorm:"type:decimal(65);not null;default:0" json:"dataSizeTotal"`   // The total data size of storage
	StorageFeeNew   decimal.Decimal `gorm:"type:decimal(65);not null;default:0" json:"storageFeeNew"`   // The fee for storage
	StorageFeeTotal decimal.Decimal `gorm:"type:decimal(65);not null;default:0" json:"storageFeeTotal"` // The total fee for storage
}

func (DASubmitStat) TableName() string {
	return "da_submit_stats"
}

type DASubmitStatStore struct {
	*mysql.Store
}

func newDASubmitStatStore(db *gorm.DB) *DASubmitStatStore {
	return &DASubmitStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *DASubmitStatStore) LastByType(statType string) (*DASubmitStat, error) {
	var submitStat DASubmitStat
	err := t.Store.DB.Where("stat_type = ?", statType).Order("stat_time desc").Last(&submitStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &submitStat, nil
}

type DASubmitStatResult struct {
	Blobs      uint64
	StorageFee decimal.Decimal
}

func (t *DASubmitStatStore) Sum(startTime, endTime time.Time, statType string) (*DASubmitStatResult, error) {
	nilTime := time.Time{}
	if startTime == nilTime && endTime == nilTime {
		return nil, errors.New("At least provide one parameter for startTime and endTime")
	}

	db := t.DB.Model(&DASubmitStat{}).Select(`IFNULL(sum(blob_new), 0) as blobs, 
		IFNULL(sum(storage_fee_new), 0) as storage_fee`)
	if startTime != nilTime && endTime != nilTime {
		db = db.Where("stat_type = ? and stat_time >= ? and stat_time < ?", statType, startTime, endTime)
	}
	if startTime != nilTime && endTime == nilTime {
		db = db.Where("stat_type = ? and stat_time >= ?", statType, startTime)
	}
	if startTime == nilTime && endTime != nilTime {
		db = db.Where("stat_type = ? and stat_time < ?", statType, endTime)
	}

	var sum DASubmitStatResult
	err := db.Find(&sum).Error
	if err != nil {
		return nil, err
	}

	return &sum, nil
}

func (t *DASubmitStatStore) Add(dbTx *gorm.DB, submitStat []*DASubmitStat) error {
	return dbTx.CreateInBatches(submitStat, batchSizeInsert).Error
}

func (t *DASubmitStatStore) Del(dbTx *gorm.DB, submitStat *DASubmitStat) error {
	return dbTx.Where("stat_type = ? and stat_time = ?", submitStat.StatType, submitStat.StatTime).Delete(&DASubmitStat{}).Error
}

func (t *DASubmitStatStore) List(intervalType *string, minTimestamp, maxTimestamp *int, desc bool, skip, limit int) (int64,
	[]DASubmitStat, error) {
	var conds []func(db *gorm.DB) *gorm.DB

	if intervalType != nil {
		intervalType := IntervalTypes[*intervalType]
		conds = append(conds, StatType(intervalType))
	}

	if minTimestamp != nil {
		conds = append(conds, MinTimestamp(*minTimestamp))
	}

	if maxTimestamp != nil {
		conds = append(conds, MaxTimestamp(*maxTimestamp))
	}

	dbRaw := t.DB.Model(&DASubmitStat{})
	dbRaw.Scopes(conds...)

	list := new([]DASubmitStat)
	total, err := t.Store.List(dbRaw, desc, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
