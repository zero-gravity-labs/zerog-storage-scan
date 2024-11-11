package stat

import (
	"container/heap"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	batchInSubmitIds = 10000
)

type TopnSubmit struct {
	*BaseStat
	heap *topnSubmitHeap
}

func MustNewTopnSubmit(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client) *AbsTopn[StatRange] {
	baseStat := &BaseStat{
		Config: cfg,
		DB:     db,
		Sdk:    sdk,
	}

	topnSubmit := &TopnSubmit{
		BaseStat: baseStat,
		heap:     newTopnSubmitHeap(10, store.StatTopnSubmitHeap, db),
	}
	topnSubmit.mustLoadLastPos()

	return &AbsTopn[StatRange]{
		Topn: topnSubmit,
	}
}

func (ts *TopnSubmit) mustLoadLastPos() {
	loaded, err := ts.loadLastPos(store.StatTopnSubmitId)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load stat pos from db")
	}

	// Submission index is set zero if not loaded.
	if !loaded {
		ts.currentPos = 0
	}
}

func (ts *TopnSubmit) nextStatRange() (*StatRange, error) {
	minPos := ts.currentPos
	maxPos := ts.currentPos + uint64(batchInSubmitIds) - 1

	block, err := ts.Sdk.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if err != nil {
		return nil, err
	}

	maxPosFinalized, exists, err := ts.DB.SubmitStore.MaxSubmissionIndexFinalized(block.Number.Uint64())
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrMaxPosFinalizedNotSync
	}

	if maxPosFinalized < minPos {
		return nil, ErrMinPosNotFinalized
	}
	if maxPosFinalized < maxPos {
		maxPos = maxPosFinalized
	}

	return &StatRange{minPos, maxPos}, nil
}

func (ts *TopnSubmit) calculateStat(r StatRange) error {
	groupedSubmits, err := ts.DB.SubmitStore.GroupBySender(r.minPos, r.maxPos)
	if err != nil {
		return err
	}

	addressesUpdate := make([]store.Address, 0)
	if len(groupedSubmits) > 0 {
		ids := make([]uint64, 0)
		for _, submit := range groupedSubmits {
			ids = append(ids, submit.SenderID)
		}

		addressesDB, err := ts.DB.BatchGetAddresses(ids)
		if err != nil {
			return err
		}

		for _, submit := range groupedSubmits {
			a := addressesDB[submit.SenderID]
			addressesUpdate = append(addressesUpdate, store.Address{
				ID:         submit.SenderID,
				DataSize:   a.DataSize + submit.DataSize,
				StorageFee: a.StorageFee.Add(submit.StorageFee),
				Txs:        a.Txs + submit.Txs,
				Files:      a.Files + submit.Files,
				UpdatedAt:  submit.UpdatedAt,
			})
		}
	}

	if err := ts.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(groupedSubmits) > 0 {
			if err := ts.DB.AddressStore.BatchUpsert(dbTx, addressesUpdate); err != nil {
				return errors.WithMessage(err, "Failed to batch update submits for topn")
			}

			ts.heap.sort(addressesUpdate)
			snapshot, err := ts.heap.snapshot()
			if err != nil {
				return err
			}

			if err := ts.DB.ConfigStore.Upsert(dbTx, ts.heap.configKey, snapshot); err != nil {
				return err
			}
		}

		if err := ts.DB.ConfigStore.Upsert(dbTx, store.StatTopnSubmitId,
			strconv.FormatUint(r.maxPos, 10)); err != nil {
			return errors.WithMessage(err, "Failed to update submit id for topn")
		}

		return nil
	}); err != nil {
		return err
	}

	ts.currentPos = r.maxPos + 1

	return nil
}

// heap

type addressItem struct {
	store.Address
	index int
}

type addressHeap []*addressItem

func (h addressHeap) Len() int { return len(h) }

func (h addressHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *addressHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*addressItem)
	item.index = n
	*h = append(*h, item)
}

func (h *addressHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

type dataSizeHeap struct {
	addressHeap
}

func (dsh dataSizeHeap) Less(i, j int) bool {
	return dsh.addressHeap[i].DataSize < dsh.addressHeap[j].DataSize
}

type storageFeeHeap struct {
	addressHeap
}

func (sfh storageFeeHeap) Less(i, j int) bool {
	return sfh.addressHeap[i].StorageFee.LessThan(sfh.addressHeap[j].StorageFee)
}

type txsHeap struct {
	addressHeap
}

func (th txsHeap) Less(i, j int) bool {
	return th.addressHeap[i].Txs < th.addressHeap[j].Txs
}

type filesHeap struct {
	addressHeap
}

func (fh filesHeap) Less(i, j int) bool {
	return fh.addressHeap[i].Files < fh.addressHeap[j].Files
}

// sort

type TopnField string

const (
	DataSize   TopnField = "data_size"
	StorageFee TopnField = "storage_fee"
	Txs        TopnField = "txs"
	Files      TopnField = "files"
)

type topnSubmitHeap struct {
	n         int
	mu        sync.Mutex
	DB        *store.MysqlStore
	configKey string

	dataSizeHeap   *dataSizeHeap
	storageFeeHeap *storageFeeHeap
	txsHeap        *txsHeap
	filesHeap      *filesHeap
}

func newTopnSubmitHeap(n int, configKey string, db *store.MysqlStore) *topnSubmitHeap {
	h := topnSubmitHeap{
		n:              n,
		DB:             db,
		configKey:      configKey,
		dataSizeHeap:   &dataSizeHeap{},
		storageFeeHeap: &storageFeeHeap{},
		txsHeap:        &txsHeap{},
		filesHeap:      &filesHeap{},
	}

	h.mustLoadFromDB()

	return &h
}

func (t *topnSubmitHeap) mustLoadFromDB() {
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

func (t *topnSubmitHeap) loadFromDB() (bool, error) {
	value, ok, err := t.DB.ConfigStore.Get(t.configKey)
	if err != nil {
		return false, errors.WithMessagef(err, "Failed to get heap cache")
	}

	if ok {
		var heaps map[TopnField][]addressItem
		if err := json.Unmarshal([]byte(value), &heaps); err != nil {
			return false, errors.WithMessage(err, "Failed to unmarshal log heap cache")
		}

		heapMap := map[TopnField]heap.Interface{
			DataSize:   t.dataSizeHeap,
			StorageFee: t.storageFeeHeap,
			Txs:        t.txsHeap,
			Files:      t.filesHeap,
		}

		for filed, h := range heapMap {
			for _, address := range heaps[filed] {
				heap.Push(h, &address)
			}
		}
	}

	return ok, nil
}

func (t *topnSubmitHeap) init() error {
	heapMap := map[TopnField]heap.Interface{
		DataSize:   t.dataSizeHeap,
		StorageFee: t.storageFeeHeap,
		Txs:        t.txsHeap,
		Files:      t.filesHeap,
	}

	for filed, h := range heapMap {
		addresses, err := t.DB.AddressStore.Topn(string(filed), 0, t.n)
		if err != nil {
			return err
		}

		for _, address := range addresses {
			heap.Push(h, &addressItem{Address: address})
		}
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

func (t *topnSubmitHeap) sort(addresses []store.Address) {
	t.mu.Lock()
	defer t.mu.Unlock()

	ds := t.deduplicate(t.dataSizeHeap.addressHeap, addresses)
	dsh := &dataSizeHeap{}
	for _, a := range ds {
		if dsh.Len() < t.n {
			heap.Push(dsh, a)
		} else if a.DataSize > dsh.addressHeap[0].DataSize {
			heap.Pop(dsh)
			heap.Push(dsh, a)
		}
	}
	t.dataSizeHeap = dsh

	sf := t.deduplicate(t.storageFeeHeap.addressHeap, addresses)
	sfh := &storageFeeHeap{}
	for _, a := range sf {
		if sfh.Len() < t.n {
			heap.Push(sfh, a)
		} else if a.StorageFee.GreaterThan(sfh.addressHeap[0].StorageFee) {
			heap.Pop(sfh)
			heap.Push(sfh, a)
		}
	}
	t.storageFeeHeap = sfh

	ts := t.deduplicate(t.txsHeap.addressHeap, addresses)
	tsh := &txsHeap{}
	for _, a := range ts {
		if tsh.Len() < t.n {
			heap.Push(tsh, a)
		} else if a.Txs > tsh.addressHeap[0].Txs {
			heap.Pop(tsh)
			heap.Push(tsh, a)
		}
	}
	t.txsHeap = tsh

	fs := t.deduplicate(t.filesHeap.addressHeap, addresses)
	fsh := &filesHeap{}
	for _, a := range fs {
		if fsh.Len() < t.n {
			heap.Push(fsh, a)
		} else if a.Files > fsh.addressHeap[0].Files {
			heap.Pop(fsh)
			heap.Push(fsh, a)
		}
	}
	t.filesHeap = fsh
}

func (t *topnSubmitHeap) deduplicate(heap addressHeap, addresses []store.Address) map[uint64]*addressItem {
	dsm := make(map[uint64]*addressItem)

	for _, a := range heap {
		dsm[a.ID] = a
	}

	for _, a := range addresses {
		dsm[a.ID] = &addressItem{Address: a}
	}

	return dsm
}

func (t *topnSubmitHeap) snapshot() (string, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	topn := map[TopnField]interface{}{
		DataSize:   t.dataSizeHeap.addressHeap,
		StorageFee: t.storageFeeHeap.addressHeap,
		Txs:        t.txsHeap.addressHeap,
		Files:      t.filesHeap.addressHeap,
	}

	info, err := json.Marshal(topn)
	if err != nil {
		return "", err
	}

	return string(info), nil
}
