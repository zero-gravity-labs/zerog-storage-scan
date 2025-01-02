package store

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/0glabs/0g-storage-scan/api/middlewares/rate"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	RateLimitStrategyConfKeyPrefix   = "ratelimit.strategy."
	rateLimitStrategySqlMatchPattern = RateLimitStrategyConfKeyPrefix + "%"

	FileExpireSeconds = "file.expire.seconds"

	SyncHeightNode    = "sync.height.node"
	SyncPatchSubmitId = "sync.patch.submit.id"

	StatTopnSubmitId        = "stat.topn.submit.id"
	StatTopnSubmitHeap      = "stat.topn.submit.heap"
	StatTopnRewardBn        = "stat.topn.reward.bn"
	StatTopnRewardHeap      = "stat.topn.reward.heap"
	StatTopnSubmitTime      = "stat.topn.submit.time"
	StatTopnRewardTime      = "stat.topn.reward.time"
	StatFileExpiredSubmitId = "stat.file.expired.submit.id"
	StatFileExpiredTotal    = "stat.file.expired.total"
	StatFilePrunedTotal     = "stat.file.pruned.total"
)

type Config struct {
	ID        uint32
	Name      string `gorm:"unique;size:128;not null"` // config name
	Value     string `gorm:"type:mediumText;not null"` // config value
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

func (cs *ConfigStore) Upsert(dbTx *gorm.DB, name, value string) error {
	db := cs.DB
	if dbTx != nil {
		db = dbTx
	}

	return db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&Config{
		Name:  name,
		Value: value,
	}).Error
}

func (cs *ConfigStore) BatchUpsert(dbTx *gorm.DB, configs []Config) error {
	db := cs.DB
	if dbTx != nil {
		db = dbTx
	}

	var placeholders string
	var params []interface{}
	size := len(configs)
	for i, c := range configs {
		placeholders += "(?,?,?,?)"
		if i != size-1 {
			placeholders += ",\n\t\t\t"
		}
		params = append(params, []interface{}{c.Name, c.Value, time.Now(), c.UpdatedAt}...)
	}

	sql := fmt.Sprintf(`
		insert into 
    		configs(name, value, created_at, updated_at)
		values
			%s
		on duplicate key update
			value = values(value),
			updated_at = values(updated_at)
	`, placeholders)

	if err := db.Exec(sql, params...).Error; err != nil {
		return err
	}

	return nil
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

func (cs *ConfigStore) BatchGet(names []string) (map[string]Config, error) {
	configs := new([]Config)
	if err := cs.DB.Raw("select * from configs where `name` in ?", names).Scan(configs).Error; err != nil {
		return nil, err
	}

	m := make(map[string]Config)
	for _, config := range *configs {
		m[config.Name] = config
	}

	return m, nil
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
