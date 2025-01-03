package sync

import (
	"context"
	"strconv"
	"time"

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
	db               *store.MysqlStore
	storageConfig    rpc.StorageConfig
	alertChannel     string
	healthReport     health.TimedCounterConfig
	storageRpcHealth health.TimedCounter
}

func MustNewStorageSyncer(db *store.MysqlStore, storageConfig rpc.StorageConfig, alertChannel string,
	healthReport health.TimedCounterConfig) *StorageSyncer {
	return &StorageSyncer{
		db:               db,
		storageConfig:    storageConfig,
		alertChannel:     alertChannel,
		healthReport:     healthReport,
		storageRpcHealth: health.TimedCounter{},
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

	if _, err := ss.db.UpdateFileInfos(ctx, unfinalized, ss.storageConfig); err != nil {
		ticker.Reset(intervalException)
	}
}

func (ss *StorageSyncer) NodeSyncHeight(ctx context.Context, ticker *time.Ticker) {
	nodeStatus, err := rpc.GetNodeStatus(ss.storageConfig)
	if err == nil {
		height := nodeStatus.LogSyncHeight
		err := ss.db.ConfigStore.Upsert(nil, store.SyncHeightNode, strconv.FormatUint(height, 10))
		if err != nil {
			logrus.WithError(err).Error("Failed to upsert storage node sync height")
		}
	}

	if ss.alertChannel != "" {
		e := rpc.AlertErr(ctx, "StorageNodeRPCError", ss.alertChannel, err, ss.healthReport,
			&ss.storageRpcHealth, ss.storageConfig.Indexer)

		if e != nil {
			ticker.Reset(intervalException)
			logrus.WithError(err).Error("Failed to alert storage node status")
		} else {
			ticker.Reset(intervalNormal)
		}
	}
}
