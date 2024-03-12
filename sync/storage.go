package sync

import (
	"context"
	"time"

	"github.com/0glabs/0g-storage-client/node"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoFileInfoToSync         = errors.New("No file info to sync")
	BatchGetSubmitsNotFinalized = 1000
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

func (ss *StorageSyncer) Sync(ctx context.Context) {
	logrus.Info("Storage syncer starting to sync data")
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if err := ss.syncFileInfo(); err != nil {
			if !errors.Is(err, ErrNoFileInfoToSync) {
				logrus.WithError(err).Error("Sync file info")
			}
			time.Sleep(time.Second * 10)
		}
	}
}

func (ss *StorageSyncer) syncFileInfo() error {
	submits, err := ss.db.SubmitStore.BatchGetNotFinalized(BatchGetSubmitsNotFinalized)
	if err != nil {
		return err
	}
	if len(submits) == 0 {
		return ErrNoFileInfoToSync
	}

	for _, s := range submits {
		info, err := ss.l2Sdk.ZeroGStorage().GetFileInfoByTxSeq(s.SubmissionIndex)
		if err != nil {
			return err
		}
		if info == nil {
			continue
		}

		submit := store.Submit{
			SubmissionIndex: s.SubmissionIndex,
			UploadedSegNum:  info.UploadedSegNum,
		}
		if !info.Finalized {
			if info.UploadedSegNum == 0 {
				submit.Status = uint8(store.NotUploaded)
			} else {
				submit.Status = uint8(store.Uploading)
			}
		} else {
			submit.Status = uint8(store.Uploaded)
		}

		addressSubmit := store.AddressSubmit{
			SenderID:        s.SenderID,
			SubmissionIndex: s.SubmissionIndex,
			UploadedSegNum:  info.UploadedSegNum,
			Status:          submit.Status,
		}

		if err := ss.db.UpdateSubmitByPrimaryKey(&submit, &addressSubmit); err != nil {
			return err
		}
	}

	return nil
}
