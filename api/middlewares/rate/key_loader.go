package rate

import (
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/sirupsen/logrus"
)

const (
	LimitKeySize = 5000
	LimitKeyTTL  = time.Minute
)

type KeyInfo struct {
	StrategyID uint32    // bound strategy ID
	LimitKey   string    // limit key
	Type       LimitType // limit type
}

type KeyFilter struct {
	StrategyIDs []uint32 // strategy IDs
	LimitKeys   []string // limit key set
	Limit       int      // result limit size (<= 0 means none)
}

type keyLoadFunc func(filter *KeyFilter) ([]*KeyInfo, error)

type LimitKeyLoader struct {
	mu       sync.Mutex
	loadFunc keyLoadFunc
	// limit key cache: limit key => *RateLimit (nil if missing)
	keyCache *TtlLruCache
}

func NewLimitKeyLoader(loadFunc keyLoadFunc) *LimitKeyLoader {
	kl := &LimitKeyLoader{
		loadFunc: loadFunc,
		keyCache: NewTtlLruCache(
			LimitKeySize, LimitKeyTTL,
		),
	}

	kl.warmUp()

	return kl
}

func (l *LimitKeyLoader) Load(key string) (*KeyInfo, bool) {
	if limitKey, ok := l.loadCache(key); ok { // cache hit
		return limitKey, ok
	}
	return l.storeCache(key)
}

func (l *LimitKeyLoader) loadCache(key string) (*KeyInfo, bool) {
	cv, expired, found := l.keyCache.Get(key)
	if found && !expired { // found in cache
		return cv.(*KeyInfo), true
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	cv, expired, found = l.keyCache.Get(key)
	if found && !expired { // double check
		return cv.(*KeyInfo), true
	}

	if found && expired {
		// extend lifespan for expired cache kv temporarily for performance
		l.keyCache.Add(key, cv.(*KeyInfo))
	}

	return nil, false
}

func (l *LimitKeyLoader) storeCache(key string) (*KeyInfo, bool) {
	// load key info from db
	limitKey, err := l.loadDB(key)
	if err != nil {
		l.mu.Lock()
		defer l.mu.Unlock()

		// for db error, we cache nil for the key by which no expiry cache value existed
		// so that db pressure can be mitigated by reducing too many subsequent queries.
		if _, _, found := l.keyCache.Get(key); !found {
			l.keyCache.Add(key, nil)
		}

		logrus.WithField("key", key).
			WithError(err).
			Error("Limit key loader failed to load limit key info")
		return nil, false
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// cache limit key
	l.keyCache.Add(key, limitKey)
	return limitKey, true
}

func (l *LimitKeyLoader) loadDB(key string) (*KeyInfo, error) {
	limitKeys, err := l.loadFunc(&KeyFilter{LimitKeys: []string{key}})
	if err == nil && len(limitKeys) > 0 {
		return limitKeys[0], nil
	}

	return nil, err
}

func (l *LimitKeyLoader) warmUp() {
	kis, err := l.loadFunc(&KeyFilter{Limit: LimitKeySize * 3 / 4})
	if err != nil {
		logrus.WithError(err).Warn("Failed to load limit key to warm up cache")
		return
	}

	for i := range kis {
		l.keyCache.Add(kis[i].LimitKey, kis[i])
	}

	logrus.WithField("totalKeys", len(kis)).Info("Limit key loaded to cache")
}

// CachedValue is used to hold ttl value
type CachedValue struct {
	value     interface{}
	expiresAt time.Time
}

// TtlLruCache naive implementation of LRU cache with fixed TTL expiration duration.
// This cache uses a lazy eviction policy, by which the expired entry will be purged when
// it's being looked up.
type TtlLruCache struct {
	mu  sync.Mutex
	lru *lru.Cache
	ttl time.Duration
}

func NewTtlLruCache(size int, ttl time.Duration) *TtlLruCache {
	cache, _ := lru.New(size)
	return &TtlLruCache{lru: cache, ttl: ttl}
}

// Add adds a value to the cache. Returns true if an eviction occurred.
func (c *TtlLruCache) Add(key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	ev := &CachedValue{
		value:     value,
		expiresAt: time.Now().Add(c.ttl),
	}

	return c.lru.Add(key, ev)
}

// GetOrPurge looks up a key's value from the cache. Will purge the entry and return nil if the entry expired.
func (c *TtlLruCache) GetOrPurge(key interface{}) (v interface{}, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cv, ok := c.lru.Get(key) // not found
	if !ok {
		return nil, false
	}

	ev := cv.(*CachedValue)
	if ev.expiresAt.Before(time.Now()) { // expired
		c.lru.Remove(key)
		return ev.value, false
	}

	return ev.value, true
}

// Get looks up a key's value from the cache without expiration action.
func (c *TtlLruCache) Get(key interface{}) (v interface{}, expired, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cv, ok := c.lru.Get(key) // not found
	if !ok {
		return nil, false, false
	}

	ev := cv.(*CachedValue)

	if ev.expiresAt.Before(time.Now()) { // expired
		return ev.value, true, true
	}

	return ev.value, false, true
}
