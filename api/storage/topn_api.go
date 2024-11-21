package storage

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	dataSizeTopn   = "data_size"
	storageFeeTopn = "storage_fee"
	txsTopn        = "txs"
	filesTopn      = "files"

	maxRecords = 10

	refreshCacheInterval = time.Second * 10
)

var (
	spanTypes = map[string]time.Duration{
		"24h": time.Hour * 24,
		"3d":  time.Hour * 24 * 3,
		"7d":  time.Hour * 24 * 7,
	}
	cache = Cache{
		topnAddresses: make(map[string]map[time.Duration][]store.TopnAddress),
		topnMiners:    make(map[time.Duration][]store.TopnMiner),
		syncHeights:   LogSyncInfo{},
	}
)

func topnDataSize(c *gin.Context) (interface{}, error) {
	return topnByType(c, dataSizeTopn)
}

func topnStorageFee(c *gin.Context) (interface{}, error) {
	return topnByType(c, storageFeeTopn)
}

func topnTxs(c *gin.Context) (interface{}, error) {
	return topnByType(c, txsTopn)
}

func topnFiles(c *gin.Context) (interface{}, error) {
	return topnByType(c, filesTopn)
}

func topnByType(c *gin.Context, t string) (interface{}, error) {
	var topnP topnParam
	if err := c.ShouldBind(&topnP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid topn query param"))
	}

	statSpan := spanTypes[topnP.SpanType]
	records := cache.topnAddresses[t][statSpan]

	result := make(map[string]interface{})
	switch t {
	case dataSizeTopn:
		list := make([]DataTopn, 0)
		for _, r := range records {
			list = append(list, DataTopn{
				Address:  r.Address,
				DataSize: r.DataSize,
			})
		}
		result["list"] = list
	case storageFeeTopn:
		list := make([]FeeTopn, 0)
		for _, r := range records {
			list = append(list, FeeTopn{
				Address:    r.Address,
				StorageFee: r.StorageFee,
			})
		}
		result["list"] = list
	case txsTopn:
		list := make([]TxsTopn, 0)
		for _, r := range records {
			list = append(list, TxsTopn{
				Address: r.Address,
				Txs:     r.Txs,
			})
		}
		result["list"] = list
	case filesTopn:
		list := make([]FilesTopn, 0)
		for _, r := range records {
			list = append(list, FilesTopn{
				Address: r.Address,
				Files:   r.Files,
			})
		}
		result["list"] = list
	default:
		return nil, api.ErrValidation(errors.Errorf("Invalid topn type %v", t))
	}

	return result, nil
}

func topnReward(c *gin.Context) (interface{}, error) {
	var topnP topnParam
	if err := c.ShouldBind(&topnP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid topn query param"))
	}

	statSpan := spanTypes[topnP.SpanType]
	miners := cache.topnMiners[statSpan]
	if len(miners) == 0 {
		return map[string]interface{}{"list": []RewardTopn{}}, nil
	}

	list := make([]RewardTopn, 0)
	for _, m := range miners {
		list = append(list, RewardTopn{
			Address: m.Address,
			Amount:  m.Amount,
		})
	}

	return map[string]interface{}{"list": list}, nil
}

type Cache struct {
	topnAddresses map[string]map[time.Duration][]store.TopnAddress
	topnMiners    map[time.Duration][]store.TopnMiner
	syncHeights   LogSyncInfo
	mu            sync.Mutex
}

func ScheduleCache(ctx context.Context, wg *sync.WaitGroup /*, cacheCh chan<- *Cache*/) {
	wg.Add(1)
	defer wg.Done()

	ticker := time.NewTicker(refreshCacheInterval)
	defer ticker.Stop()

	logrus.Info("Schedule cache starting")
	for {
		select {
		case <-ctx.Done():
			logrus.Info("Schedule cache shutdown ok")
			return
		case <-ticker.C:
			if err := cacheTopn(); err != nil {
				logrus.WithError(err).Error("Failed to schedule cache topn")
			}
			if err := cacheSyncHeights(); err != nil {
				logrus.WithError(err).Error("Failed to schedule cache sync heights")
			}
		}
	}
}

func cacheTopn() error {
	statSpan := make([]time.Duration, 0)
	for _, duration := range spanTypes {
		statSpan = append(statSpan, duration)
	}

	topnAddresses := make(map[string]map[time.Duration][]store.TopnAddress) // field => duration => addresses
	topnMiners := make(map[time.Duration][]store.TopnMiner)                 // duration => miners

	if err := loadTopnSubmits(statSpan, topnAddresses); err != nil {
		return err
	}
	if err := loadTopnRewards(statSpan, topnMiners); err != nil {
		return err
	}
	if err := loadTopnSubmitsOverall(topnAddresses); err != nil {
		return err
	}
	if err := loadTopnRewardsOverall(topnMiners); err != nil {
		return err
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.topnAddresses = topnAddresses
	cache.topnMiners = topnMiners

	return nil
}

func loadTopnSubmits(durations []time.Duration, topnAddresses map[string]map[time.Duration][]store.TopnAddress) error {
	for _, duration := range durations {
		for _, field := range []string{dataSizeTopn, storageFeeTopn, txsTopn, filesTopn} {
			addresses, err := db.SubmitTopnStatStore.Topn(field, duration, maxRecords)
			if err != nil {
				return errors.WithMessage(err, "Failed to cache topn submit")
			}

			if _, ok := topnAddresses[field]; !ok {
				topnAddresses[field] = make(map[time.Duration][]store.TopnAddress)
			}

			topnAddresses[field][duration] = addresses
		}
	}

	return nil
}

func loadTopnRewards(durations []time.Duration, topnMiners map[time.Duration][]store.TopnMiner) error {
	for _, duration := range durations {
		miners, err := db.RewardTopnStatStore.Topn(duration, maxRecords)
		if err != nil {
			return errors.WithMessage(err, "Failed to cache topn reward")
		}

		topnMiners[duration] = miners
	}

	return nil
}

func loadTopnSubmitsOverall(topnAddresses map[string]map[time.Duration][]store.TopnAddress) error {
	value, ok, err := db.ConfigStore.Get(store.StatTopnSubmitHeap)
	if err != nil {
		return errors.WithMessagef(err, "Failed to get submit heap")
	}

	if ok {
		var heaps map[string][]store.TopnAddress // field => addresses
		if err := json.Unmarshal([]byte(value), &heaps); err != nil {
			return errors.WithMessage(err, "Failed to unmarshal submit heap")
		}

		for field, addresses := range heaps {
			topnAddresses[field][0] = addresses
		}
	}

	return nil
}

func loadTopnRewardsOverall(topnMiners map[time.Duration][]store.TopnMiner) error {
	value2, ok, err := db.ConfigStore.Get(store.StatTopnRewardHeap)
	if err != nil {
		return errors.WithMessagef(err, "Failed to load reward heap")
	}

	if ok {
		var minerSlice []store.Miner
		if err := json.Unmarshal([]byte(value2), &minerSlice); err != nil {
			return errors.WithMessage(err, "Failed to unmarshal reward heap")
		}

		addrIDs := make([]uint64, 0)
		for _, miner := range minerSlice {
			addrIDs = append(addrIDs, miner.ID)
		}
		addrMap, err := db.BatchGetAddresses(addrIDs)
		if err != nil {
			return errors.WithMessage(err, "Failed to batch get addresses")
		}

		miners := make([]store.TopnMiner, 0)
		for _, miner := range minerSlice {
			miners = append(miners, store.TopnMiner{
				Address: addrMap[miner.ID].Address,
				Amount:  miner.Amount,
			})
		}

		topnMiners[0] = miners
	}

	return nil
}

func cacheSyncHeights() error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	value, ok, err := db.ConfigStore.Get(store.SyncHeightNode)
	if err != nil {
		return errors.WithMessage(err, "Failed to get node sync height")
	}
	if !ok {
		return errors.New("No matching record found(node sync height)")
	}
	nodeSyncHeight, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "Failed to parse node sync height")
	}

	scanSyncHeight, ok, err := db.BlockStore.MaxBlock()
	if err != nil {
		return errors.WithMessage(err, "Failed to get scan sync height")
	}
	if !ok {
		return errors.New("No matching record found(scan sync height)")
	}

	cache.syncHeights = LogSyncInfo{
		Layer1LogSyncHeight: nodeSyncHeight,
		LogSyncHeight:       scanSyncHeight,
	}

	return nil
}
