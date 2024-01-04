package store

import (
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type CostStat struct {
	ID             uint64     `gorm:"primaryKey" json:"-"`
	StatTime       *time.Time `gorm:"not null;index:idx_statTime_statType,unique,priority:1" json:"statTime"`
	StatType       string     `gorm:"type:char(3);not null;index:idx_statTime_statType,unique,priority:2" json:"-"`
	BasicCost      uint64     `gorm:"not null;default:0" json:"basicCost"`      // The basic cost for storage
	BasicCostTotal uint64     `gorm:"not null;default:0" json:"basicCostTotal"` // The total basic cost for storage
}

func (CostStat) TableName() string {
	return "cost_stats"
}

type CostStatStore struct {
	*mysql.Store
}

func newCostStatStore(db *gorm.DB) *CostStatStore {
	return &CostStatStore{
		Store: mysql.NewStore(db),
	}
}

func (c *CostStatStore) LastByType(statType string) (*CostStat, error) {
	var costStat CostStat
	err := c.Store.DB.Where("stat_type = ?", statType).Last(&costStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &costStat, nil
}

func (c *CostStatStore) Sum(startTime, endTime *time.Time, statType string) (uint64, error) {
	if startTime == nil && endTime == nil {
		return 0, errors.New("At least provide one parameter for startTime and endTime")
	}

	db := c.DB.Model(&CostStat{}).Select("IFNULL(sum(basic_cost), 0) as basic_cost_sum")
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
		BasicCostSum int64
	}

	err := db.Find(&sum).Error
	if err != nil {
		return 0, err
	}

	return uint64(sum.BasicCostSum), nil
}

func (c *CostStatStore) Add(dbTx *gorm.DB, costStats []*CostStat) error {
	return dbTx.CreateInBatches(costStats, batchSizeInsert).Error
}

func (c *CostStatStore) Del(dbTx *gorm.DB, costStat *CostStat) error {
	return dbTx.Where("stat_time = ? and stat_type = ?", costStat.StatTime, costStat.StatType).Delete(&CostStat{}).Error
}
