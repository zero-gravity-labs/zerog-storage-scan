package stat

import (
	"math/big"
	"strconv"
	"time"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StatSubmitExpired struct {
	*BaseStat
	fileExpireSeconds *big.Int
	fileExpiredTotal  uint64
}

func MustNewStatSubmitExpired(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client) *AbsTopn[StatRange] {
	baseStat := &BaseStat{
		Config: cfg,
		DB:     db,
		Sdk:    sdk,
	}

	statSubmitExpired := &StatSubmitExpired{
		BaseStat: baseStat,
	}
	statSubmitExpired.mustLoadLastPos()
	statSubmitExpired.mustLoadFileExpireSeconds()
	statSubmitExpired.mustLoadTotalExpiredFiles()

	return &AbsTopn[StatRange]{
		Topn: statSubmitExpired,
	}
}

func (ts *StatSubmitExpired) mustLoadLastPos() {
	loaded, err := ts.loadLastPos(store.StatFileExpiredSubmitId)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load stat pos from db")
	}

	// Submission index is set zero if not loaded.
	if !loaded {
		ts.currentPos = 0
	}
}

func (ts *StatSubmitExpired) mustLoadFileExpireSeconds() {
	value, exist, err := ts.DB.ConfigStore.Get(store.FileExpireSeconds)
	if err != nil || !exist {
		logrus.WithError(err).Fatal("Failed to get file expiration from DB")
	}

	seconds, success := new(big.Int).SetString(value, 10)
	if !success {
		logrus.WithError(err).Fatal("Failed to parse file expiration from DB")
	}

	ts.fileExpireSeconds = seconds
}

func (ts *StatSubmitExpired) mustLoadTotalExpiredFiles() {
	value, exist, err := ts.DB.ConfigStore.Get(store.StatFileExpiredTotal)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get total expired files from DB")
	}
	if !exist {
		ts.fileExpiredTotal = uint64(0)
		return
	}

	total, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to parse total expired files from DB")
	}

	ts.fileExpiredTotal = total
}

func (ts *StatSubmitExpired) nextStatRange() (*StatRange, error) {
	minPos := ts.currentPos
	maxPos := ts.currentPos + uint64(batchInSubmitIds) - 1

	block, err := ts.Sdk.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if err != nil {
		return nil, err
	}

	timestamp := block.Timestamp - ts.fileExpireSeconds.Uint64()
	maxExpireTime := time.Unix(int64(timestamp), 0)

	maxPosExpired, exists, err := ts.DB.SubmitStore.MaxSubmissionIndexExpired(maxExpireTime)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrMaxPosFinalizedNotSync
	}

	if maxPosExpired < minPos {
		return nil, ErrMinPosNotFinalized
	}
	if maxPosExpired < maxPos {
		maxPos = maxPosExpired
	}

	return &StatRange{minPos, maxPos}, nil
}

func (ts *StatSubmitExpired) calculateStat(r StatRange) error {
	groupedSubmits, err := ts.DB.SubmitStore.GroupBySender(r.minPos, r.maxPos)
	if err != nil {
		return err
	}

	addressesUpdate := make([]store.Address, 0)
	if len(groupedSubmits) > 0 {
		for _, submit := range groupedSubmits {
			addressesUpdate = append(addressesUpdate, store.Address{
				ID:           submit.SenderID,
				ExpiredFiles: submit.Files,
			})
			ts.fileExpiredTotal += submit.Files
		}
	}

	if err := ts.DB.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(groupedSubmits) > 0 {
			if err := ts.DB.AddressStore.BatchDeltaUpsertExpiredFiles(dbTx, addressesUpdate); err != nil {
				return errors.WithMessage(err, "Failed to batch update submits for stat expired files")
			}

			if err := ts.DB.ConfigStore.Upsert(dbTx, store.StatFileExpiredTotal,
				strconv.FormatUint(ts.fileExpiredTotal, 10)); err != nil {
				return errors.WithMessage(err, "Failed to update total count for stat expired files")
			}
		}

		if err := ts.DB.ConfigStore.Upsert(dbTx, store.StatFileExpiredSubmitId,
			strconv.FormatUint(r.maxPos, 10)); err != nil {
			return errors.WithMessage(err, "Failed to update submit id for stat expired files")
		}

		return nil
	}); err != nil {
		return err
	}

	ts.currentPos = r.maxPos + 1

	return nil
}
