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
	BatchGetSubmitsNotFinalized       = 10000
	BatchGetSubmitsNotFinalizedLatest = 100
)

type StorageSyncer struct {
	l2Sdks            []*node.Client
	db                *store.MysqlStore
	alertChannel      string
	healthReport      health.TimedCounterConfig
	storageRpcHealths []health.TimedCounter
}

func MustNewStorageSyncer(l2Sdks []*node.Client, db *store.MysqlStore, alertChannel string,
	healthReport health.TimedCounterConfig) *StorageSyncer {
	return &StorageSyncer{
		l2Sdks:            l2Sdks,
		db:                db,
		alertChannel:      alertChannel,
		healthReport:      healthReport,
		storageRpcHealths: make([]health.TimedCounter, len(l2Sdks)),
	}
}

func (ss *StorageSyncer) Sync(ctx context.Context, syncFunc func(ctx2 context.Context) error) {
	logrus.Info("Storage syncer starting to sync data.")
	for {
		select {
		case <-ctx.Done():
			return
		default:
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
		submits, err := ss.db.SubmitStore.GetUnfinalizedOverall(lastSubmissionIndex, BatchGetSubmitsNotFinalized)
		if err != nil {
			return err
		}
		if len(submits) == 0 {
			return ErrNoFileInfoToSync
		}

		if _, err := rpc.RefreshFileInfos(ctx, submits, ss.l2Sdks, ss.db); err != nil {
			return err
		}

		lastSubmissionIndex = submits[len(submits)-1].SubmissionIndex + 1
	}
}

func (ss *StorageSyncer) SyncLatest(ctx context.Context) error {
	submits, err := ss.db.SubmitStore.GetUnfinalizedLatest(BatchGetSubmitsNotFinalizedLatest)
	if err != nil {
		return err
	}
	if len(submits) == 0 {
		return ErrNoFileInfoToSync
	}

	if _, err := rpc.RefreshFileInfos(ctx, submits, ss.l2Sdks, ss.db); err != nil {
		return err
	}

	return nil
}
