package stat

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	batchInSubmitIds = 10000
	maxSubmits       = 100
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
		heap:     newTopnSubmitHeap(maxSubmits, store.StatTopnSubmitHeap, db),
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
				Address:    a.Address,
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

			if err := ts.heap.store(dbTx); err != nil {
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
	heapMap := map[TopnField]heap.Interface{
		DataSize:   t.dataSizeHeap,
		StorageFee: t.storageFeeHeap,
		Txs:        t.txsHeap,
		Files:      t.filesHeap,
	}
	fields := make([]TopnField, 0)
	for field := range heapMap {
		fields = append(fields, field)
	}

	names := make([]string, 0)
	for _, field := range fields {
		names = append(names, fmt.Sprintf("%s.%s", t.configKey, string(field)))
	}
	configs, err := t.DB.ConfigStore.BatchGet(names)
	if err != nil {
		return false, err
	}

	configCount := len(configs)
	if configCount == 0 { // not exist
		return false, nil
	}

	if configCount != len(fields) {
		return false, errors.New("Topn cache not match with topn fields")
	}

	for _, field := range fields {
		var addresses []addressItem
		c := configs[fmt.Sprintf("%s.%s", t.configKey, string(field))]
		if err := json.Unmarshal([]byte(c.Value), &addresses); err != nil {
			return false, errors.WithMessagef(err, "Failed to unmarshal heap cache for %s", field)
		}

		h := heapMap[field]
		for _, address := range addresses {
			heap.Push(h, &address)
		}
	}

	return true, nil
}

func (t *topnSubmitHeap) init() error {
	heapMap := map[TopnField]heap.Interface{
		DataSize:   t.dataSizeHeap,
		StorageFee: t.storageFeeHeap,
		Txs:        t.txsHeap,
		Files:      t.filesHeap,
	}

	for field, h := range heapMap {
		addresses, err := t.DB.AddressStore.Topn(string(field), 0, t.n)
		if err != nil {
			return err
		}

		for _, address := range addresses {
			heap.Push(h, &addressItem{Address: address})
		}
	}

	if err := t.store(nil); err != nil {
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

func (t *topnSubmitHeap) store(dbTx *gorm.DB) error {
	snapshots, err := t.snapshot()
	if err != nil {
		return err
	}

	configs := make([]store.Config, 0)
	for field, s := range snapshots {
		configs = append(configs, store.Config{
			Name:      fmt.Sprintf("%s.%s", t.configKey, string(field)),
			Value:     s,
			UpdatedAt: time.Now(),
		})
	}

	if err := t.DB.ConfigStore.BatchUpsert(dbTx, configs); err != nil {
		return err
	}

	return nil
}

func (t *topnSubmitHeap) snapshot() (map[TopnField]string, error) {
	heapMap := map[TopnField]heap.Interface{
		DataSize:   t.dataSizeHeap,
		StorageFee: t.storageFeeHeap,
		Txs:        t.txsHeap,
		Files:      t.filesHeap,
	}

	topnCopy := make(map[TopnField][]store.Address)
	for field, h := range heapMap {
		addresses := make([]store.Address, h.Len())
		for h.Len() > 0 {
			item := heap.Pop(h).(*addressItem)
			addresses[h.Len()] = item.Address
		}
		topnCopy[field] = addresses
	}

	snapshots := make(map[TopnField]string)
	for field, addresses := range topnCopy {
		snapshot, err := json.Marshal(addresses)
		if err != nil {
			return nil, err
		}
		snapshots[field] = string(snapshot)

		h := heapMap[field]
		for _, address := range addresses {
			h.Push(&addressItem{Address: address})
		}
	}

	return snapshots, nil
}
