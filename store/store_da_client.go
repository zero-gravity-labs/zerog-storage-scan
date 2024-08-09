package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DAClient struct {
	ID              uint64
	FirstUploadTime time.Time `gorm:"not null"`
}

func (DAClient) TableName() string {
	return "da_clients"
}

type DAClientStore struct {
	*mysql.Store
}

func newDAClientStore(db *gorm.DB) *DAClientStore {
	return &DAClientStore{
		Store: mysql.NewStore(db),
	}
}

func (ms *DAClientStore) Add(id uint64, firstUploadTime time.Time) (uint64, error) {
	var cli DAClient
	existed, err := ms.Store.Exists(&cli, "id = ?", id)
	if err != nil {
		return 0, err
	}
	if existed {
		return cli.ID, nil
	}

	cli = DAClient{
		ID:              id,
		FirstUploadTime: firstUploadTime,
	}

	if err := ms.DB.Create(&cli).Error; err != nil {
		return 0, err
	}

	return cli.ID, nil
}

func (ms *DAClientStore) Count(startTime, endTime time.Time) (uint64, error) {
	db := ms.DB.Model(&DAClient{})
	nilTime := time.Time{}
	if startTime != nilTime && endTime != nilTime {
		db = db.Where("first_upload_time >= ? and first_upload_time < ?", startTime, endTime)
	}
	if startTime == nilTime && endTime != nilTime {
		db = db.Where("first_upload_time < ?", endTime)
	}

	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint64(count), nil
}

type DAClientStat struct {
	ID       uint64    `json:"-"`
	StatType string    `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime time.Time `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`

	ClientNew    uint64 `gorm:"not null;default:0" json:"clientNew"`    // Number of da client in a specific time interval
	ClientActive uint64 `gorm:"not null;default:0" json:"clientActive"` // Number of active da client in a specific time interval
	ClientTotal  uint64 `gorm:"not null;default:0" json:"clientTotal"`  // Total number of da client by a certain time
}

func (DAClientStat) TableName() string {
	return "da_client_stats"
}

type DAClientStatStore struct {
	*mysql.Store
}

func newDAClientStatStore(db *gorm.DB) *DAClientStatStore {
	return &DAClientStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *DAClientStatStore) LastByType(statType string) (*DAClientStat, error) {
	var daClientStat DAClientStat
	err := t.Store.DB.Where("stat_type = ?", statType).Order("stat_time desc").Last(&daClientStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &daClientStat, nil
}

func (t *DAClientStatStore) Add(dbTx *gorm.DB, daClientStats []*DAClientStat) error {
	return dbTx.CreateInBatches(daClientStats, batchSizeInsert).Error
}

func (t *DAClientStatStore) Del(dbTx *gorm.DB, daClientStat *DAClientStat) error {
	return dbTx.Where("stat_type = ? and stat_time = ?", daClientStat.StatType, daClientStat.StatTime).Delete(&DAClientStat{}).Error
}

func (t *DAClientStatStore) List(intervalType *string, minTimestamp, maxTimestamp *int, desc bool, skip, limit int) (int64,
	[]DAClientStat, error) {
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

	dbRaw := t.DB.Model(&DAClientStat{})
	dbRaw.Scopes(conds...)

	list := new([]DAClientStat)
	total, err := t.Store.List(dbRaw, desc, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
