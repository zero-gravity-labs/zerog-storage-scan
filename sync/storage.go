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
	BatchGetSubmitsNotFinalized = 10000
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
	skip := 0
	for {
		submits, err := ss.db.SubmitStore.BatchGetNotFinalized(skip, BatchGetSubmitsNotFinalized)
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
				submit.UploadedSegNum = submit.TotalSegNum // Field `uploadedSegNum` is set 0 by rpc when `finalized` is true
			}

			addressSubmit := store.AddressSubmit{
				SenderID:        s.SenderID,
				SubmissionIndex: s.SubmissionIndex,
				UploadedSegNum:  submit.UploadedSegNum,
				Status:          submit.Status,
			}

			if err := ss.db.UpdateSubmitByPrimaryKey(&submit, &addressSubmit); err != nil {
				return err
			}
		}

		if len(submits) < BatchGetSubmitsNotFinalized {
			return nil
		} else {
			skip += BatchGetSubmitsNotFinalized
		}
	}
}
