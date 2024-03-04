package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	Min    = "1m"
	TenMin = "10m"
	Hour   = "1h"
	Day    = "1d"
)

var (
	Intervals = map[string]time.Duration{
		Min:    time.Minute,
		TenMin: time.Minute * 10,
		Hour:   time.Hour,
		Day:    time.Hour * 24,
	}

	IntervalTypes = map[string]string{
		"min":   Min,
		"10min": TenMin,
		"hour":  Hour,
		"day":   Day,
	}
)

type SubmitStat struct {
	ID           uint64          `json:"-"`
	StatType     string          `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime     time.Time       `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`
	FileCount    uint64          `gorm:"not null;default:0" json:"fileCount"`                     // Number of files in a specific time interval
	FileTotal    uint64          `gorm:"not null;default:0" json:"fileTotal"`                     // Total number of files by a certain time
	DataSize     uint64          `gorm:"not null;default:0" json:"dataSize"`                      // Size of storage data in a specific time interval
	DataTotal    uint64          `gorm:"not null;default:0" json:"dataTotal"`                     // Total Size of storage data by a certain time
	BaseFee      decimal.Decimal `gorm:"type:decimal(65);not null;default:0" json:"baseFee"`      // The base fee for storage
	BaseFeeTotal decimal.Decimal `gorm:"type:decimal(65);not null;default:0" json:"baseFeeTotal"` // The total base fee for storage
}

func (SubmitStat) TableName() string {
	return "submit_stats"
}

type SubmitStatStore struct {
	*mysql.Store
}

func newSubmitStatStore(db *gorm.DB) *SubmitStatStore {
	return &SubmitStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *SubmitStatStore) LastByType(statType string) (*SubmitStat, error) {
	var submitStat SubmitStat
	err := t.Store.DB.Where("stat_type = ?", statType).Order("stat_time asc").Last(&submitStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &submitStat, nil
}

type SubmitStatResult struct {
	FileCount uint64
	DataSize  uint64
	BaseFee   decimal.Decimal
}

func (t *SubmitStatStore) Sum(startTime, endTime *time.Time, statType string) (*SubmitStatResult, error) {
	if startTime == nil && endTime == nil {
		return nil, errors.New("At least provide one parameter for startTime and endTime")
	}

	db := t.DB.Model(&SubmitStat{}).Select(`IFNULL(sum(file_count), 0) as file_count, 
		IFNULL(sum(data_size), 0) as data_size, IFNULL(sum(base_fee), 0) as base_fee`)
	if startTime != nil && endTime != nil {
		db = db.Where("stat_type = ? and stat_time >= ? and stat_time < ?", statType, startTime, endTime)
	}
	if startTime != nil && endTime == nil {
		db = db.Where("stat_type = ? and stat_time >= ?", statType, startTime)
	}
	if startTime == nil && endTime != nil {
		db = db.Where("stat_type = ? and stat_time < ?", statType, endTime)
	}

	var sum SubmitStatResult
	err := db.Find(&sum).Error
	if err != nil {
		return nil, err
	}

	return &sum, nil
}

func (t *SubmitStatStore) Add(dbTx *gorm.DB, submitStat []*SubmitStat) error {
	return dbTx.CreateInBatches(submitStat, batchSizeInsert).Error
}

func (t *SubmitStatStore) Del(dbTx *gorm.DB, submitStat *SubmitStat) error {
	return dbTx.Where("stat_type = ? and stat_time = ?", submitStat.StatType, submitStat.StatTime).Delete(&SubmitStat{}).Error
}

func (t *SubmitStatStore) List(intervalType *string, minTimestamp, maxTimestamp *int, desc bool, skip, limit int) (int64,
	[]SubmitStat, error) {
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

	dbRaw := t.DB.Model(&SubmitStat{})
	dbRaw.Scopes(conds...)

	list := new([]SubmitStat)
	total, err := t.Store.List(dbRaw, desc, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
