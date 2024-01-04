package store

import (
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type TxStat struct {
	ID       uint64     `gorm:"primaryKey" json:"-"`
	StatTime *time.Time `gorm:"not null;index:idx_statTime_statType,unique,priority:1" json:"statTime"`
	StatType string     `gorm:"type:char(3);not null;index:idx_statTime_statType,unique,priority:2" json:"-"`
	TxCount  uint64     `gorm:"not null;default:0" json:"txCount"`
	TxTotal  uint64     `gorm:"not null;default:0" json:"txTotal"`
}

func (TxStat) TableName() string {
	return "tx_stats"
}

type TxStatStore struct {
	*mysql.Store
}

func newTxStatStore(db *gorm.DB) *TxStatStore {
	return &TxStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *TxStatStore) LastByType(statType string) (*TxStat, error) {
	var txStat TxStat
	err := t.Store.DB.Where("stat_type = ?", statType).Last(&txStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &txStat, nil
}

func (t *TxStatStore) Sum(startTime, endTime *time.Time, statType string) (uint64, error) {
	var sum int64

	db := t.DB.Model(&TxStat{}).Select("IFNULL(sum(tx_count), 0) as sum")

	if startTime != nil && endTime != nil {
		db = db.Where("stat_time >= ? and stat_time < ? and stat_type = ?", startTime, endTime, statType)
	}
	if startTime != nil && endTime == nil {
		db = db.Where("stat_time >= ? and stat_type = ?", startTime, statType)
	}
	if startTime == nil && endTime != nil {
		db = db.Where("stat_time < ? and stat_type = ?", endTime, statType)
	}

	err := db.Find(&sum).Error
	if err != nil {
		return 0, err
	}

	return uint64(sum), nil
}

func (t *TxStatStore) Add(dbTx *gorm.DB, txStat []*TxStat) error {
	return dbTx.CreateInBatches(txStat, batchSizeInsert).Error
}

func (t *TxStatStore) Del(dbTx *gorm.DB, txStat *TxStat) error {
	return dbTx.Where("stat_time = ? and stat_type = ?", txStat.StatTime, txStat.StatType).Delete(&TxStat{}).Error
}
