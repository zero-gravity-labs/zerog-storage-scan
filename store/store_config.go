package store

import (
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	KeyLogSyncInfo = "LogSyncInfo"
)

type Config struct {
	Name  string `gorm:"size:32;primaryKey"`
	Value string `gorm:"size:512"`
}

func (Config) TableName() string {
	return "configs"
}

type ConfigStore struct {
	*mysql.Store
}

func newConfigStore(db *gorm.DB) *ConfigStore {
	return &ConfigStore{
		Store: mysql.NewStore(db),
	}
}

func (cs *ConfigStore) Upsert(name, value string) error {
	return cs.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&Config{
		Name:  name,
		Value: value,
	}).Error
}

func (cs *ConfigStore) Get(name string) (string, bool, error) {
	var cfg Config
	err := cs.DB.Where("name = ?", name).Take(&cfg).Error
	if err == nil {
		return cfg.Value, true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return cfg.Value, false, nil
	}

	return cfg.Value, false, err
}
