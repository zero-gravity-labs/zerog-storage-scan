package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	refreshCacheInterval = time.Second * 10
)

var (
	cache = Cache{
		topnAddresses:   make(map[string]map[time.Duration][]store.TopnAddress),
		topnMiners:      make(map[time.Duration][]store.TopnMiner),
		syncHeights:     LogSyncInfo{},
		minerRewardStat: MinerRewardStat{},
	}
)

type Cache struct {
	topnAddresses   map[string]map[time.Duration][]store.TopnAddress
	topnMiners      map[time.Duration][]store.TopnMiner
	syncHeights     LogSyncInfo
	minerRewardStat MinerRewardStat
	mu              sync.Mutex
}

func ScheduleCache(ctx context.Context, wg *sync.WaitGroup) {
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
			if err := cacheMinerRewardStat(); err != nil {
				logrus.WithError(err).Error("Failed to schedule cache miner reward stat")
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
	fields := []string{dataSizeTopn, storageFeeTopn, txsTopn, filesTopn}
	names := make([]string, 0)
	for _, field := range fields {
		names = append(names, fmt.Sprintf("%s.%s", store.StatTopnSubmitHeap, field))
	}
	configs, err := db.ConfigStore.BatchGet(names)
	if err != nil {
		return err
	}

	configCount := len(configs)
	if configCount == 0 { // not exist
		return nil
	}

	if configCount != len(fields) {
		return errors.New("Topn cache not match with topn fields")
	}

	for _, field := range fields {
		var addresses []store.TopnAddress
		c := configs[fmt.Sprintf("%s.%s", store.StatTopnSubmitHeap, field)]
		if err := json.Unmarshal([]byte(c.Value), &addresses); err != nil {
			return errors.WithMessagef(err, "Failed to unmarshal heap cache for %s", field)
		}

		topnAddresses[field][0] = addresses
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

func cacheMinerRewardStat() error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	avg, err := db.RewardStore.AvgRewardRecently(time.Hour * 24)
	if err != nil {
		return err
	}

	stat, err := db.RewardStatStore.LastByType(store.Day)
	if err != nil {
		return err
	}

	cache.minerRewardStat = MinerRewardStat{
		AvgReward24Hours: *avg,
		TotalReward:      stat.RewardTotal,
	}

	return nil
}
