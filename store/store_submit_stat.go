package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SubmitStat struct {
	ID        uint64    `json:"-"`
	StatType  string    `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime  time.Time `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`
	FileCount uint64    `gorm:"not null;default:0" json:"fileCount"` // Number of files in a specific time interval
	FileTotal uint64    `gorm:"not null;default:0" json:"fileTotal"` // Total number of files by a certain time
	DataSize  uint64    `gorm:"not null;default:0" json:"dataSize"`  // Size of storage data in a specific time interval
	DataTotal uint64    `gorm:"not null;default:0" json:"dataTotal"` // Total Size of storage data by a certain time
	// TODO not enough for blockchain value of decimals 18.
	BaseFee      uint64 `gorm:"not null;default:0" json:"baseFee"`      // The base fee for storage
	BaseFeeTotal uint64 `gorm:"not null;default:0" json:"baseFeeTotal"` // The total base fee for storage
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
	BaseFee   uint64
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
