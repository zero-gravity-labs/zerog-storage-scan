package rate

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/0glabs/0g-storage-scan/api/middlewares/metrics"
	"github.com/Conflux-Chain/go-conflux-util/http/middlewares"
	commonRate "github.com/Conflux-Chain/go-conflux-util/rate"
	commonHttp "github.com/Conflux-Chain/go-conflux-util/rate/http"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

const (
	GCScheduleInterval = 5 * time.Minute
)

type LimiterFactory struct {
	*commonHttp.Registry

	mu             sync.Mutex
	limitKeyLoader *LimitKeyLoader

	strategies    map[string]*Strategy // strategy name => *Strategy
	id2Strategies map[uint32]*Strategy // strategy id => *Strategy
}

func NewLimiterFactory(limitKeyLoader *LimitKeyLoader) *LimiterFactory {
	m := &LimiterFactory{
		limitKeyLoader: limitKeyLoader,
		strategies:     make(map[string]*Strategy),
		id2Strategies:  make(map[uint32]*Strategy),
	}

	m.Registry = commonHttp.NewRegistry(m)
	go m.ScheduleGC(GCScheduleInterval)

	return m
}

func (lf *LimiterFactory) GetGroupAndKey(ctx context.Context, resource string) (group, key string, err error) {
	apiKey, ok := middlewares.GetApiKeyFromContext(ctx)
	if !ok {
		// use default strategy if not authenticated
		return lf.genDefaultGroupAndKey(ctx, resource)
	}

	if limitKey, ok := lf.limitKeyLoader.Load(apiKey); ok && limitKey != nil {
		// use strategy with corresponding key info
		return lf.genKeyInfoGroupAndKey(ctx, resource, apiKey, limitKey)
	}

	return lf.genDefaultGroupAndKey(ctx, resource)
}

func (lf *LimiterFactory) Create(ctx context.Context, resource, group string) (commonRate.Limiter, error) {
	lf.mu.Lock()
	defer lf.mu.Unlock()

	stg, ok := lf.strategies[group]
	if !ok {
		return nil, errors.New("strategy not found")
	}

	opt, ok := stg.LimitOptions[resource]
	if !ok {
		return nil, errors.New("limit rule not found")
	}

	return lf.createWithOption(opt)
}

func (lf *LimiterFactory) genDefaultGroupAndKey(ctx context.Context, resource string) (group, key string, err error) {
	lf.mu.Lock()
	defer lf.mu.Unlock()

	stg, ok := lf.strategies[DefaultStrategy]
	if !ok { // no default strategy
		logrus.WithField("resource", resource).Info("Default strategy not configured")
		return
	}

	if _, ok := stg.LimitOptions[resource]; !ok {
		// limit rule not defined
		return
	}

	ip, _ := middlewares.GetRealIPFromContext(ctx)
	key = fmt.Sprintf("ip:%v", ip)

	return stg.Name, key, nil
}

func (lf *LimiterFactory) genKeyInfoGroupAndKey(ctx context.Context, resource, limitKey string, ki *KeyInfo) (
	group, key string, err error) {
	lf.mu.Lock()
	defer lf.mu.Unlock()

	stg, ok := lf.id2Strategies[ki.StrategyID]
	if !ok {
		logrus.WithFields(logrus.Fields{
			"limitKey": limitKey,
			"resource": resource,
			"keyInfo":  ki,
		}).Warn("Rate limit strategy not found")
		return
	}

	if _, ok := stg.LimitOptions[resource]; !ok {
		// limit rule not defined
		return
	}

	group = stg.Name

	switch ki.Type {
	case LimitTypeByIp: // limit by key-based IP
		ip, _ := middlewares.GetRealIPFromContext(ctx)
		key = fmt.Sprintf("key:%v/ip:%v", limitKey, ip)

	case LimitTypeByKey: // limit by key only
		key = fmt.Sprintf("key:%v", limitKey)

	default:
		err = errors.New("invalid limit type")
	}

	return group, key, err
}

func (lf *LimiterFactory) createWithOption(option interface{}) (l commonRate.Limiter, err error) {
	switch opt := option.(type) {
	case FixedWindowOption:
		l = commonRate.NewFixedWindow(opt.Interval, opt.Quota)
	case TokenBucketOption:
		l = commonRate.NewTokenBucket(int(opt.Rate), opt.Burst)
	default:
		err = errors.New("invalid limit option")
	}

	return l, err
}

type Config struct {
	// checksums for modification detection
	CheckSums  ConfigCheckSums
	Strategies map[uint32]*Strategy // limit strategies
}

// ConfigCheckSums config md5 checksum
type ConfigCheckSums struct {
	Strategies map[uint32][md5.Size]byte
}

func (lf *LimiterFactory) AutoReload(interval time.Duration, reloader func() (*Config, error)) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// last config fingerprints
	var cs ConfigCheckSums

	// load immediately at first
	if rconf, err := reloader(); err == nil {
		lf.reload(rconf, &cs)
		cs = rconf.CheckSums
	}

	// load periodically
	for range ticker.C {
		rconf, err := reloader()
		if err != nil {
			logrus.WithError(err).Error("Failed to load rate limit configs")
			continue
		}

		lf.reload(rconf, &cs)
		cs = rconf.CheckSums
	}
}

func (lf *LimiterFactory) reload(rc *Config, lastCs *ConfigCheckSums) {
	if rc == nil {
		return
	}

	lf.mu.Lock()
	defer lf.mu.Unlock()

	// remove strategy
	for sid, strategy := range lf.id2Strategies {
		if _, ok := rc.Strategies[sid]; !ok {
			lf.removeStrategy(strategy)
			logrus.WithField("strategy", strategy).Info("RateLimit strategy removed")
		}
	}

	// add or update strategy
	for sid, strategy := range rc.Strategies {
		s, ok := lf.id2Strategies[sid]
		if !ok { // add
			lf.addStrategy(strategy)
			logrus.WithField("strategy", strategy).Info("RateLimit strategy added")
			continue
		}

		if lastCs.Strategies[sid] != rc.CheckSums.Strategies[sid] { // update
			lf.updateStrategy(s, strategy)
			logrus.WithField("strategy", strategy).Info("RateLimit strategy updated")
		}
	}
}

func (lf *LimiterFactory) removeStrategy(s *Strategy) {
	// remove all limiters under this strategy
	for resource := range s.LimitOptions {
		lf.Remove(resource, s.Name)
		logrus.WithField("resource", resource).Info("RateLimit rule deleted")
	}

	delete(lf.strategies, s.Name)
	delete(lf.id2Strategies, s.ID)
}

func (lf *LimiterFactory) addStrategy(s *Strategy) {
	lf.strategies[s.Name] = s
	lf.id2Strategies[s.ID] = s
}

func (lf *LimiterFactory) updateStrategy(old, new *Strategy) {
	lf.strategies[new.Name] = new
	lf.id2Strategies[new.ID] = new

	if old.Name != new.Name {
		// strategy name changed? this shall be rare, but we also need to
		// delete all limiters with the groups of old strategy name.
		delete(lf.strategies, old.Name)
		for resource := range old.LimitOptions {
			lf.Remove(resource, old.Name)
			logrus.WithField("resource", resource).Info("RateLimit rule deleted")
		}

		return
	}

	// check the changes from old to new limit rule, and delete all old limiters
	// with resource whose limit rule has been altered.
	for resource, oldopt := range old.LimitOptions {
		newopt, ok := new.LimitOptions[resource]
		if !ok || !reflect.DeepEqual(oldopt, newopt) {
			lf.Remove(resource, old.Name)
			logrus.WithField("resource", resource).Info("RateLimit rule deleted")
		}
	}
}

// #################### strategy option ####################

type LimitAlgoType string

const (
	LimitAlgoFixedWindow LimitAlgoType = "fixed_window"
	LimitAlgoTokenBucket LimitAlgoType = "token_bucket"
)

type LimitType int

const (
	LimitTypeByKey LimitType = iota
	LimitTypeByIp
)

const (
	DefaultStrategy = "default"
)

// Strategy rate limit strategy
type Strategy struct {
	ID   uint32 // strategy ID
	Name string // strategy name

	LimitOptions map[string]interface{} // resource => limit option
}

func NewStrategy(id uint32, name string) *Strategy {
	return &Strategy{
		ID:           id,
		Name:         name,
		LimitOptions: make(map[string]interface{}),
	}
}

// UnmarshalJSON implements `json.Unmarshaler`
func (s *Strategy) UnmarshalJSON(data []byte) error {
	tmpRules := make(map[string]*LimitRule)
	if err := json.Unmarshal(data, &tmpRules); err != nil {
		return errors.WithMessage(err, "malformed json format")
	}

	for resource, rule := range tmpRules {
		s.LimitOptions[resource] = rule.Option
	}

	return nil
}

// FixedWindowOption limit option for fixed window
type FixedWindowOption struct {
	Interval time.Duration
	Quota    int
}

// UnmarshalJSON implements `json.Unmarshaler`
func (fwo *FixedWindowOption) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Interval string
		Quota    int
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	fwo.Quota = tmp.Quota

	interval, err := time.ParseDuration(tmp.Interval)
	if err == nil {
		fwo.Interval = interval
	}

	return err
}

// TokenBucketOption limit option for token bucket
type TokenBucketOption struct {
	Rate  rate.Limit
	Burst int
}

func NewTokenBucketOption(r, b int) TokenBucketOption {
	return TokenBucketOption{
		Rate:  rate.Limit(r),
		Burst: b,
	}
}

// LimitRule resource limit rule
type LimitRule struct {
	Algo   LimitAlgoType
	Option interface{}
}

func (r *LimitRule) UnmarshalJSON(data []byte) (err error) {
	var tmp struct {
		Algo   LimitAlgoType
		Option json.RawMessage
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	r.Algo = tmp.Algo

	switch tmp.Algo {
	case LimitAlgoFixedWindow:
		var opt FixedWindowOption
		if err = json.Unmarshal(tmp.Option, &opt); err == nil {
			r.Option = opt
		}
	case LimitAlgoTokenBucket:
		var opt TokenBucketOption
		if err = json.Unmarshal(tmp.Option, &opt); err == nil {
			r.Option = opt
		}
	default:
		return errors.New("invalid rate limit algorithm")
	}

	return err
}

// ################# single api middleware #################

func NewAPIRateMiddleware(f commonHttp.LimitFunc) middlewares.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			urlType, ok := r.Context().Value(metrics.CtxKeyURLType).(string)
			if !ok {
				return
			}

			resource := fmt.Sprintf("%v_qps", urlType)
			if err := f(r.Context(), resource); err != nil {
				http.Error(w, err.Error(), http.StatusTooManyRequests)
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}
