package sync

import (
	"context"
	"time"

	"github.com/0glabs/0g-storage-client/node"
	"github.com/0glabs/0g-storage-scan/rpc"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoFileInfoToSync          = errors.New("No file info to sync")
	BatchGetSubmitsLatest        = 100
	checkStatusIntervalNormal    = time.Second
	checkStatusIntervalException = time.Second * 10
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

func (ss *StorageSyncer) Sync(ctx context.Context, syncFunc func(ctx2 context.Context) error) {
	logrus.Info("Storage syncer starting to sync data.")
	for {
		if interrupted(ctx) {
			return
		}

		if err := syncFunc(ctx); err != nil {
			if !errors.Is(err, ErrNoFileInfoToSync) {
				logrus.WithError(err).Error("Failed to sync storage data")
			}
			time.Sleep(time.Second * 10)
		}
	}
}

func (ss *StorageSyncer) SyncLatest(ctx context.Context) error {
	if interrupted(ctx) {
		return nil
	}

	submits, err := ss.db.SubmitStore.QueryDesc(BatchGetSubmitsLatest)
	if err != nil {
		return err
	}

	if len(submits) == 0 {
		return ErrNoFileInfoToSync
	}

	unfinalized := make([]store.Submit, 0)
	for _, submit := range submits {
		if submit.Status < uint8(rpc.Pruned) {
			unfinalized = append(unfinalized, submit)
		}
	}

	if len(unfinalized) == 0 {
		return ErrNoFileInfoToSync
	}

	if _, err := ss.db.UpdateFileInfos(ctx, unfinalized, ss.l2Sdks); err != nil {
		return err
	}

	return nil
}

func (ss *StorageSyncer) CheckStatus(ctx context.Context) {
	ticker := time.NewTicker(checkStatusIntervalNormal)
	defer ticker.Stop()

	logrus.Info("Storage syncer starting to alert.")
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ss.checkStatusOnce(ctx, ticker)
		}
	}
}

func (ss *StorageSyncer) checkStatusOnce(ctx context.Context, ticker *time.Ticker) {
	for index, l2Sdk := range ss.l2Sdks {
		_, err := l2Sdk.GetStatus(ctx)

		if ss.alertChannel != "" {
			e := rpc.AlertErr(ctx, "StorageNodeRPCError", ss.alertChannel, err, ss.healthReport,
				ss.storageRpcHealths[index], l2Sdk.URL())

			if e != nil {
				ticker.Reset(checkStatusIntervalException)
				logrus.WithError(err).Error("Failed to alert storage status")
			} else {
				ticker.Reset(checkStatusIntervalNormal)
			}
		}
	}
}
