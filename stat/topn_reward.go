package stat

import (
	"container/heap"
	"encoding/json"
	"strconv"
	"sync"

	"gorm.io/gorm"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	batchInBns = 1000
)

type TopnReward struct {
	*BaseStat
	heap *topnMinerHeap
}

func MustNewTopnReward(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client) *AbsTopn[StatRange] {
	baseStat := &BaseStat{
		Config: cfg,
		DB:     db,
		Sdk:    sdk,
	}

	topnReward := &TopnReward{
		BaseStat: baseStat,
		heap:     newTopnMinerHeap(10, store.StatTopnRewardHeap, db),
	}
	topnReward.mustLoadLastPos()

	return &AbsTopn[StatRange]{
		Topn: topnReward,
	}
}

func (tr *TopnReward) mustLoadLastPos() {
	loaded, err := tr.loadLastPos(store.StatTopnRewardBn)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load stat pos from db")
	}

	// Reward bn is set config value if not loaded.
	if !loaded {
		tr.currentPos = tr.Config.BlockOnStatBegin
	}
}

func (tr *TopnReward) nextStatRange() (*StatRange, error) {
	minPos := tr.currentPos
	maxPos := tr.currentPos + uint64(batchInBns) - 1

	block, err := tr.Sdk.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if err != nil {
		return nil, err
	}
	maxPosFinalized := block.Number.Uint64()

	if maxPosFinalized < minPos {
		return nil, ErrMinPosNotFinalized
	}
	if maxPosFinalized < maxPos {
		maxPos = maxPosFinalized
	}

	return &StatRange{minPos, maxPos}, nil
}

func (tr *TopnReward) calculateStat(r StatRange) error {
	groupedRewards, err := tr.DB.RewardStore.GroupByMiner(r.minPos, r.maxPos)
	if err != nil {
		return err
	}

	minersUpdate := make([]store.Miner, 0)
	if len(groupedRewards) > 0 {
		ids := make([]uint64, 0)
		for _, reward := range groupedRewards {
			ids = append(ids, reward.MinerID)
		}

		minersDB, err := tr.DB.BatchGetMiners(ids)
		if err != nil {
			return err
		}

		for _, reward := range groupedRewards {
			m := minersDB[reward.MinerID]
			minersUpdate = append(minersUpdate, store.Miner{
				ID:        reward.MinerID,
				Amount:    m.Amount.Add(reward.Amount),
				UpdatedAt: reward.UpdatedAt,
			})
		}
	}

	if err := tr.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(minersUpdate) > 0 {
			if err := tr.DB.MinerStore.BatchUpsert(dbTx, minersUpdate); err != nil { //TODO
				return errors.WithMessage(err, "Failed to batch update miners for topn")
			}

			tr.heap.sort(minersUpdate)
			snapshot, err := tr.heap.snapshot()
			if err != nil {
				return err
			}

			if err := tr.DB.ConfigStore.Upsert(dbTx, tr.heap.configKey, snapshot); err != nil {
				return err
			}
		}

		if err := tr.DB.ConfigStore.Upsert(dbTx, store.StatTopnRewardBn,
			strconv.FormatUint(r.maxPos, 10)); err != nil {
			return errors.WithMessage(err, "Failed to batch update bn for topn")
		}
		return nil
	}); err != nil {
		return err
	}

	tr.currentPos = r.maxPos + 1

	return nil
}

// heap

type minerItem struct {
	store.Miner
	index int
}

type minerHeap []*minerItem

func (h minerHeap) Len() int { return len(h) }

func (h minerHeap) Less(i, j int) bool { return h[i].Amount.LessThan(h[j].Amount) }

func (h minerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *minerHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*minerItem)
	item.index = n
	*h = append(*h, item)
}

func (h *minerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

// sort

type topnMinerHeap struct {
	n         int
	mu        sync.Mutex
	DB        *store.MysqlStore
	configKey string

	minerHeap *minerHeap
}

func newTopnMinerHeap(n int, configKey string, db *store.MysqlStore) *topnMinerHeap {
	h := topnMinerHeap{
		n:         n,
		DB:        db,
		configKey: configKey,
		minerHeap: &minerHeap{},
	}

	h.mustLoadFromDB()

	return &h
}

func (t *topnMinerHeap) mustLoadFromDB() {
	loaded, err := t.loadFromDB()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load heap from db")
	}

	if !loaded {
		if err := t.init(); err != nil {
			logrus.WithError(err).Fatal("Failed to init heap")
		}
	}
}

func (t *topnMinerHeap) loadFromDB() (bool, error) {
	value, ok, err := t.DB.ConfigStore.Get(t.configKey)
	if err != nil {
		return false, errors.WithMessagef(err, "Failed to get heap cache")
	}

	if ok {
		var minerItems []minerItem
		if err := json.Unmarshal([]byte(value), &minerItems); err != nil {
			return false, errors.WithMessage(err, "Failed to unmarshal log heap cache")
		}
		for _, minerItem := range minerItems {
			heap.Push(t.minerHeap, &minerItem)
		}
	}

	return ok, nil
}

func (t *topnMinerHeap) init() error {
	miners, err := t.DB.MinerStore.Topn(0, t.n)
	if err != nil {
		return err
	}

	for _, miner := range miners {
		heap.Push(t.minerHeap, &minerItem{Miner: miner})
	}

	snapshot, err := t.snapshot()
	if err != nil {
		return err
	}

	if err := t.DB.ConfigStore.Upsert(nil, t.configKey, snapshot); err != nil {
		return err
	}

	return nil
}

func (t *topnMinerHeap) sort(miners []store.Miner) {
	t.mu.Lock()
	defer t.mu.Unlock()

	minerItemMap := t.deduplicate(*t.minerHeap, miners)
	h := &minerHeap{}
	for _, m := range minerItemMap {
		if h.Len() < t.n {
			heap.Push(h, m)
		} else if m.Amount.GreaterThan((*h)[0].Amount) {
			heap.Pop(h)
			heap.Push(h, m)
		}
	}
	t.minerHeap = h
}

func (t *topnMinerHeap) deduplicate(heap minerHeap, miners []store.Miner) map[uint64]*minerItem {
	minerMap := make(map[uint64]*minerItem)

	for _, m := range heap {
		minerMap[m.ID] = m
	}

	for _, m := range miners {
		minerMap[m.ID] = &minerItem{Miner: m}
	}

	return minerMap
}

func (t *topnMinerHeap) snapshot() (string, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	topnMiners := make([]store.Miner, t.minerHeap.Len())
	for t.minerHeap.Len() > 0 {
		item := heap.Pop(t.minerHeap).(*minerItem)
		topnMiners[t.minerHeap.Len()] = item.Miner
	}

	for _, miner := range topnMiners {
		t.minerHeap.Push(&minerItem{Miner: miner})
	}

	info, err := json.Marshal(topnMiners)
	if err != nil {
		return "", err
	}

	return string(info), nil
}
