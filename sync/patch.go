package sync

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoFileInfoToPatch   = errors.New("No file info to patch")
	BatchGetSubmitsToPatch = 100
)

type PatchSyncer struct {
	sdk             *web3go.Client
	db              *store.MysqlStore
	currentSubmitId uint64
}

func MustNewPatchSyncer(sdk *web3go.Client, db *store.MysqlStore) *PatchSyncer {
	syncer := PatchSyncer{
		sdk: sdk,
		db:  db,
	}
	syncer.mustLoadLastSubmitId()

	return &syncer
}

func (ps *PatchSyncer) mustLoadLastSubmitId() {
	loaded, err := ps.loadLastSubmitId()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load submit id from db")
	}

	// Submit id is set 0 value if not loaded.
	if !loaded {
		ps.currentSubmitId = 0
	}
}

func (ps *PatchSyncer) loadLastSubmitId() (loaded bool, err error) {
	value, ok, err := ps.db.ConfigStore.Get(store.SyncPatchSubmitId)
	if err != nil {
		return false, errors.WithMessagef(err, "Failed to get submit id")
	}

	if ok {
		submitId, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return false, errors.WithMessagef(err, "Invalid submit id %s", value)
		}
		ps.currentSubmitId = submitId + 1
	}

	return ok, nil
}

func (ps *PatchSyncer) Sync(ctx context.Context) {
	logrus.Info("Patch syncer starting to sync data.")
	for {
		if interrupted(ctx) {
			return
		}

		if err := ps.syncOnce(ctx); err != nil {
			if !errors.Is(err, ErrNoFileInfoToPatch) {
				logrus.WithError(err).Error("Failed to patch data")
			}
			time.Sleep(time.Second * 10)
		}
	}
}

func (ps *PatchSyncer) syncOnce(ctx context.Context) error {
	for {
		submits, err := ps.db.SubmitStore.QueryOverallByAsc(&ps.currentSubmitId, BatchGetSubmitsToPatch)
		if err != nil {
			return err
		}
		if len(submits) == 0 {
			return ErrNoFileInfoToPatch
		}

		for _, submit := range submits {
			if interrupted(ctx) {
				return nil
			}

			hash := common.HexToHash(submit.TxHash)
			tx, err := ps.sdk.Eth.TransactionByHash(hash)
			if err != nil {
				return errors.WithMessage(err, "Failed to get tx")
			}
			rcpt, err := ps.sdk.Eth.TransactionReceipt(hash)
			if err != nil {
				return errors.WithMessage(err, "Failed to get rcpt")
			}
			if tx == nil || rcpt == nil { // Not patch pruned tx
				continue
			}

			var extra store.SubmitExtra
			if err = json.Unmarshal(submit.Extra, &extra); err != nil {
				return errors.WithMessage(err, "Failed to unmarshal extra")
			}

			extra.GasPrice = tx.GasPrice.Uint64()
			extra.GasLimit = tx.Gas
			extra.GasUsed = rcpt.GasUsed
			updateExtra, err := json.Marshal(extra)
			if err != nil {
				return errors.WithMessage(err, "Failed to marshal extra")
			}

			update := store.Submit{
				SubmissionIndex: submit.SubmissionIndex,
				Extra:           updateExtra,
			}
			if err := ps.db.SubmitStore.UpdateByPrimaryKey(nil, &update); err != nil {
				return err
			}
		}

		lastSubmitId := submits[len(submits)-1].SubmissionIndex
		if err := ps.db.ConfigStore.Upsert(nil, store.SyncPatchSubmitId,
			strconv.FormatUint(lastSubmitId, 10)); err != nil {
			return err
		}

		ps.currentSubmitId = lastSubmitId + 1
	}
}
