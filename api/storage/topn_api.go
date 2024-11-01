package storage

import (
	"context"
	"sync"
	"time"

	scanApi "github.com/0glabs/0g-storage-scan/api"
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
		topnAddresses: make(map[string]map[time.Duration][]store.Address),
		topnMiners:    make(map[time.Duration][]store.Miner),
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

	addrIDs := make([]uint64, 0)
	for _, miner := range miners {
		addrIDs = append(addrIDs, miner.ID)
	}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, scanApi.ErrBatchGetAddress(err)
	}

	list := make([]RewardTopn, 0)
	for _, m := range miners {
		list = append(list, RewardTopn{
			Address: addrMap[m.ID].Address,
			Amount:  m.Amount,
		})
	}

	return map[string]interface{}{"list": list}, nil
}

type Cache struct {
	topnAddresses map[string]map[time.Duration][]store.Address
	topnMiners    map[time.Duration][]store.Miner
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
				logrus.WithError(err).Error("Failed to schedule cache")
			}
		}
	}
}

func cacheTopn() error {
	statSpan := make([]time.Duration, 0)
	for _, duration := range spanTypes {
		statSpan = append(statSpan, duration)
	}
	statSpan = append(statSpan, 0)

	topnAddresses := make(map[string]map[time.Duration][]store.Address)
	topnMiners := make(map[time.Duration][]store.Miner)

	for _, duration := range statSpan {
		for _, order := range []string{dataSizeTopn, storageFeeTopn, txsTopn, filesTopn} {
			addresses, err := db.AddressStore.Topn(order, duration, maxRecords)
			if err != nil {
				return errors.WithMessage(err, "Failed to cache topn submit")
			}
			if _, ok := topnAddresses[order]; !ok {
				topnAddresses[order] = make(map[time.Duration][]store.Address)
			}
			topnAddresses[order][duration] = addresses
		}

		miners, err := db.MinerStore.Topn(duration, maxRecords)
		if err != nil {
			return errors.WithMessage(err, "Failed to cache topn reward")
		}
		topnMiners[duration] = miners
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.topnAddresses = topnAddresses
	cache.topnMiners = topnMiners

	return nil
}
