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
	ErrNoFileInfoToSync               = errors.New("No file info to sync")
	BatchGetSubmitsNotFinalized       = 100
	BatchGetSubmitsNotFinalizedLatest = 100
)

type StorageSyncer struct {
	l2Sdks            []*node.Client
	db                *store.MysqlStore
	alertChannel      string
	healthReport      health.TimedCounterConfig
	storageRpcHealths []*health.TimedCounter
}

func MustNewStorageSyncer(l2Sdks []*node.Client, db *store.MysqlStore, alertChannel string,
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

func (ss *StorageSyncer) SyncOverall(ctx context.Context) error {
	lastSubmissionIndex := uint64(0)

	for {
		submits, err := ss.db.SubmitStore.QueryUnfinalizedByAsc(&lastSubmissionIndex, BatchGetSubmitsNotFinalized)
		if err != nil {
			return err
		}
		if len(submits) == 0 {
			return ErrNoFileInfoToSync
		}

		if interrupted(ctx) {
			return nil
		}

		if _, err := ss.db.UpdateFileInfos(ctx, submits, ss.l2Sdks); err != nil {
			return err
		}

		lastSubmissionIndex = submits[len(submits)-1].SubmissionIndex + 1
	}
}

func (ss *StorageSyncer) SyncLatest(ctx context.Context) error {
	submits, err := ss.db.SubmitStore.QueryUnfinalizedLatestByDesc(BatchGetSubmitsNotFinalizedLatest)
	if err != nil {
		return err
	}
	if len(submits) == 0 {
		return ErrNoFileInfoToSync
	}

	if interrupted(ctx) {
		return nil
	}

	if _, err := ss.db.UpdateFileInfos(ctx, submits, ss.l2Sdks); err != nil {
		return err
	}

	return nil
}

func (ss *StorageSyncer) CheckStatus(ctx context.Context) error {
	for index, l2Sdk := range ss.l2Sdks {
		_, err := l2Sdk.ZeroGStorage().GetStatus()

		if ss.alertChannel != "" {
			if e := rpc.AlertErr(ctx, "StorageNodeRPCError", ss.alertChannel, err, ss.healthReport,
				ss.storageRpcHealths[index], l2Sdk.URL()); e != nil {
				return e
			}
		}
	}

	return nil
}
