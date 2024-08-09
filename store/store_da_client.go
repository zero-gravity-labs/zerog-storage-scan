package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
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
