package sync

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zero-gravity-labs/zerog-storage-client/node"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
)

var (
	ErrNoRootHashToSync = errors.New("No root hash to sync")
	ErrFileInfoNotReady = errors.New("File info not ready")
)

type StorageSyncer struct {
	l2Sdk *node.Client
	db    *store.MysqlStore
}

func MustNewStorageSyncer(l2Sdk *node.Client, db *store.MysqlStore) *StorageSyncer {
	return &StorageSyncer{
		l2Sdk: l2Sdk,
		db:    db,
	}
}

func (s *StorageSyncer) Sync(ctx context.Context) {
	logrus.Info("Storage syncer starting to sync data")
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if err := s.syncRootHash(); err != nil {
			if !errors.Is(err, ErrNoRootHashToSync) {
				logrus.WithError(err).Error("Sync root hash")
			}
			time.Sleep(time.Second * 10)
		}
	}
}

// TODO add finality filed when add code for recalculate root hash
func (s *StorageSyncer) syncRootHash() error {
	submit, err := s.db.SubmitStore.FirstWithoutRootHash()
	if err != nil {
		return err
	}
	if submit == nil {
		return ErrNoRootHashToSync
	}

	info, err := s.l2Sdk.ZeroGStorage().GetFileInfoByTxSeq(submit.SubmissionIndex)
	if err != nil {
		return err
	}
	if info == nil {
		return ErrFileInfoNotReady
	}

	updateSubmit := store.Submit{
		SubmissionIndex: submit.SubmissionIndex,
		RootHash:        info.Tx.DataMerkleRoot.String()[2:],
	}
	if err := s.db.SubmitStore.UpdateByPrimaryKey(&updateSubmit); err != nil {
		return err
	}

	return nil
}
