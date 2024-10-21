package store

import (
	"time"

	nhRate "github.com/0glabs/0g-storage-scan/api/middlewares/rate"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"gorm.io/gorm"
)

type RateLimit struct {
	ID         uint32
	LimitKey   string `gorm:"unique;size:128;not null"`
	LimitType  int    `gorm:"default:0;not null"`
	StrategyID uint32 `gorm:"index"`
	Remark     string `gorm:"size:128"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (RateLimit) TableName() string {
	return "rate_limits"
}

type RateLimitStore struct {
	*mysql.Store
}

func newRateLimitStore(db *gorm.DB) *RateLimitStore {
	return &RateLimitStore{
		Store: mysql.NewStore(db),
	}
}

func (r *RateLimitStore) Add(
	strategyID uint32,
	limitType int,
	limitKey string,
	remark string,
) error {
	ratelimit := &RateLimit{
		StrategyID: strategyID,
		LimitType:  int(limitType),
		LimitKey:   limitKey,
		Remark:     remark,
	}

	return r.DB.Create(ratelimit).Error
}

func (r *RateLimitStore) Delete(limitKey string) (bool, error) {
	res := r.DB.Delete(&RateLimit{}, "limit_key = ?", limitKey)
	return res.RowsAffected > 0, res.Error
}

func (r *RateLimitStore) Get(limitKey string) (RateLimit, bool, error) {
	var rateLimit RateLimit
	exist, err := r.Store.Exists(&rateLimit, "limit_key = ?", limitKey)
	return rateLimit, exist, err
}

func (r *RateLimitStore) List(limitKeys []string, strategyIDs []uint32, limit int) (res []*RateLimit, err error) {
	db := r.DB

	if len(limitKeys) > 0 {
		db = db.Where("limit_key in (?)", limitKeys)
	}

	if len(strategyIDs) > 0 {
		db = db.Where("strategy_id in (?)", strategyIDs)
	}

	if limit > 0 {
		db = db.Limit(limit)
	}

	var rateLimits []*RateLimit

	err = db.FindInBatches(&rateLimits, 1000, func(tx *gorm.DB, batch int) error {
		res = append(res, rateLimits...)
		return nil
	}).Error

	return res, err
}

func (r *RateLimitStore) ListLimitKeyInfos(filter *nhRate.KeyFilter) (res []*nhRate.KeyInfo, err error) {
	rateLimits, err := r.List(filter.LimitKeys, filter.StrategyIDs, filter.Limit)
	if err != nil {
		return nil, err
	}

	for _, limit := range rateLimits {
		res = append(res, &nhRate.KeyInfo{
			Type:       nhRate.LimitType(limit.LimitType),
			LimitKey:   limit.LimitKey,
			StrategyID: limit.StrategyID,
		})
	}

	return res, nil
}
