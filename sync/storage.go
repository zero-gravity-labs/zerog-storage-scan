package sync

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/0glabs/0g-storage-client/node"
	"github.com/0glabs/0g-storage-scan/rpc"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/sirupsen/logrus"
)

var (
	BatchGetSubmitsLatest = 100
	intervalNormal        = time.Second
	intervalException     = time.Second * 10
)

type StorageSyncer struct {
	l2Sdks            []*node.ZgsClient
	db                *store.MysqlStore
	alertChannel      string
	healthReport      health.TimedCounterConfig
	storageRpcHealths []*health.TimedCounter
}

func MustNewStorageSyncer(l2Sdks []*node.ZgsClient, db *store.MysqlStore, alertChannel string,
	healthReport health.TimedCounterConfig) *StorageSyncer {

	storageRpcHealths := make([]*health.TimedCounter, len(l2Sdks))
	for i := 0; i < len(l2Sdks); i++ {
		storageRpcHealths[i] = &health.TimedCounter{}
	}

	return &StorageSyncer{
		l2Sdks:            l2Sdks,
		db:                db,
		alertChannel:      alertChannel,
		healthReport:      healthReport,
		storageRpcHealths: storageRpcHealths,
	}
}

func (ss *StorageSyncer) Sync(ctx context.Context, f func(ctx context.Context, ticker *time.Ticker)) {
	ticker := time.NewTicker(intervalNormal)
	defer ticker.Stop()

	logrus.Info("Storage syncer starting to sync data.")
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			f(ctx, ticker)
		}
	}
}

func (ss *StorageSyncer) LatestFiles(ctx context.Context, ticker *time.Ticker) {
	if interrupted(ctx) {
		return
	}

	submits, err := ss.db.SubmitStore.QueryDesc(BatchGetSubmitsLatest)
	if err != nil {
		ticker.Reset(intervalException)
	}

	if len(submits) == 0 {
		return
	}

	unfinalized := make([]store.Submit, 0)
	for _, submit := range submits {
		if submit.Status < uint8(rpc.Pruned) {
			unfinalized = append(unfinalized, submit)
		}
	}

	if len(unfinalized) == 0 {
		return
	}

	if _, err := ss.db.UpdateFileInfos(ctx, unfinalized, ss.l2Sdks); err != nil {
		ticker.Reset(intervalException)
	}
}

func (ss *StorageSyncer) NodeSyncHeight(ctx context.Context, ticker *time.Ticker) {
	heights := make([]uint64, len(ss.l2Sdks))

	for index, l2Sdk := range ss.l2Sdks {
		if interrupted(ctx) {
			return
		}

		nodeStatus, err := l2Sdk.GetStatus(ctx)
		if err == nil {
			heights[index] = nodeStatus.LogSyncHeight
		}

		if ss.alertChannel != "" {
			e := rpc.AlertErr(ctx, "StorageNodeRPCError", ss.alertChannel, err, ss.healthReport,
				ss.storageRpcHealths[index], l2Sdk.URL())

			if e != nil {
				ticker.Reset(intervalException)
				logrus.WithError(err).Error("Failed to alert storage node status")
			} else {
				ticker.Reset(intervalException)
			}
		}
	}

	sort.Slice(heights, func(i, j int) bool { return heights[i] > heights[j] })

	if heights[0] > 0 {
		err := ss.db.ConfigStore.Upsert(nil, store.SyncHeightNode, strconv.FormatUint(heights[0], 10))
		if err != nil {
			logrus.WithError(err).Error("Failed to upsert storage node sync height")
		}
	}
}
