package store

import (
	"crypto/md5"
	"encoding/json"
	"time"

	"github.com/0glabs/0g-storage-scan/api/middlewares/rate"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	KeyLogSyncInfo                   = "LogSyncInfo"
	RateLimitStrategyConfKeyPrefix   = "ratelimit.strategy."
	rateLimitStrategySqlMatchPattern = RateLimitStrategyConfKeyPrefix + "%"
)

type Config struct {
	ID        uint32
	Name      string `gorm:"unique;size:128;not null"` // config name
	Value     string `gorm:"size:16250;not null"`      // config value
	CreatedAt time.Time
	UpdatedAt time.Time
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

func (cs *ConfigStore) LoadRateLimitConfigs() (*rate.Config, error) {
	rlStrategies, csStrategies, err := cs.LoadRateLimitStrategyConfigs()
	if err != nil {
		return nil, err
	}

	return &rate.Config{
		CheckSums: rate.ConfigCheckSums{
			Strategies: csStrategies,
		},
		Strategies: rlStrategies,
	}, nil
}

func (cs *ConfigStore) LoadRateLimitStrategyConfigs() (map[uint32]*rate.Strategy, map[uint32][md5.Size]byte, error) {
	var cfgs []Config
	if err := cs.DB.Where("name LIKE ?", rateLimitStrategySqlMatchPattern).Find(&cfgs).Error; err != nil {
		return nil, nil, err
	}

	if len(cfgs) == 0 {
		return nil, nil, nil
	}

	strategies := make(map[uint32]*rate.Strategy)
	checksums := make(map[uint32][md5.Size]byte)

	// decode ratelimit strategy from config item
	for _, v := range cfgs {
		strategy, err := cs.decodeRateLimitStrategy(v)
		if err != nil {
			logrus.WithField("cfg", v).WithError(err).Warn("Invalid rate limit strategy config")
			continue
		}

		strategies[v.ID] = strategy
		checksums[v.ID] = md5.Sum([]byte(v.Value))
	}

	return strategies, checksums, nil
}

func (cs *ConfigStore) decodeRateLimitStrategy(cfg Config) (*rate.Strategy, error) {
	// eg., ratelimit.strategy.whitelist
	name := cfg.Name[len(RateLimitStrategyConfKeyPrefix):]
	if len(name) == 0 {
		return nil, errors.New("strategy name is too short")
	}

	data := []byte(cfg.Value)
	stg := rate.NewStrategy(cfg.ID, name)

	if err := json.Unmarshal(data, stg); err != nil {
		return nil, err
	}

	return stg, nil
}
