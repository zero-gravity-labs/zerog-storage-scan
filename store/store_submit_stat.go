package store

import (
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type SubmitStat struct {
	ID        uint64     `gorm:"primaryKey" json:"-"`
	StatTime  *time.Time `gorm:"not null;index:idx_statTime_statType,unique,priority:1" json:"statTime"`
	StatType  string     `gorm:"type:char(3);not null;index:idx_statTime_statType,unique,priority:2" json:"-"`
	FileCount uint64     `gorm:"not null;default:0" json:"fileCount"` // Number of files in a specific time interval
	FileTotal uint64     `gorm:"not null;default:0" json:"fileTotal"` // Total number of files by a certain time
	DataSize  uint64     `gorm:"not null;default:0" json:"dataSize"`  // Size of storage data in a specific time interval
	DataTotal uint64     `gorm:"not null;default:0" json:"dataTotal"` // Total Size of storage data by a certain time
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
	err := t.Store.DB.Where("stat_type = ?", statType).Last(&submitStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &submitStat, nil
}

func (t *SubmitStatStore) Sum(startTime, endTime *time.Time, statType string) (uint64, uint64, error) {
	if startTime == nil && endTime == nil {
		return 0, 0, errors.New("At least provide one parameter for startTime and endTime")
	}

	db := t.DB.Model(&SubmitStat{}).Select("IFNULL(sum(file_count), 0) as file_sum, IFNULL(sum(data_size), 0) as data_sum")
	if startTime != nil && endTime != nil {
		db = db.Where("stat_time >= ? and stat_time < ? and stat_type = ?", startTime, endTime, statType)
	}
	if startTime != nil && endTime == nil {
		db = db.Where("stat_time >= ? and stat_type = ?", startTime, statType)
	}
	if startTime == nil && endTime != nil {
		db = db.Where("stat_time < ? and stat_type = ?", endTime, statType)
	}

	var sum struct {
		FileSum int64
		DataSum int64
	}

	err := db.Find(&sum).Error
	if err != nil {
		return 0, 0, err
	}

	return uint64(sum.FileSum), uint64(sum.DataSum), nil
}

func (t *SubmitStatStore) Add(dbTx *gorm.DB, submitStat []*SubmitStat) error {
	return dbTx.CreateInBatches(submitStat, batchSizeInsert).Error
}

func (t *SubmitStatStore) Del(dbTx *gorm.DB, submitStat *SubmitStat) error {
	return dbTx.Where("stat_time = ? and stat_type = ?", submitStat.StatTime, submitStat.StatType).Delete(&SubmitStat{}).Error
}
