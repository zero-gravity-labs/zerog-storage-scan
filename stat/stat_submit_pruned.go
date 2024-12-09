package stat

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/0glabs/0g-storage-scan/rpc"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrNoFileInfoToStat   = errors.New("No file info to stat")
	BatchGetSubmitsToStat = 10000
)

type StatSubmitPruned struct {
	db              *store.MysqlStore
	filePrunedTotal uint64
}

func MustNewStatSubmitPruned(db *store.MysqlStore) *StatSubmitPruned {
	statSubmitPruned := &StatSubmitPruned{
		db: db,
	}
	statSubmitPruned.mustLoadTotalPrunedFiles()

	return statSubmitPruned
}

func (s *StatSubmitPruned) mustLoadTotalPrunedFiles() {
	value, exist, err := s.db.ConfigStore.Get(store.StatFilePrunedTotal)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get total pruned files from DB")
	}
	if !exist {
		s.filePrunedTotal = uint64(0)
		return
	}

	total, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to parse total pruned files from DB")
	}

	s.filePrunedTotal = total
}

func (s *StatSubmitPruned) DoStat(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	logrus.Info("Stat starting")

	for {
		if s.interrupted(ctx) {
			return
		}

		if err := s.stat(ctx); err != nil {
			if !errors.Is(err, ErrNoFileInfoToStat) {
				logrus.WithError(err).Error("Failed to stat pruned files")
			}
			time.Sleep(time.Second * 10)
		}
	}
}

func (s *StatSubmitPruned) stat(ctx context.Context) error {
	maxSubmitId, err := s.db.SubmitStore.MaxSubmissionIndex()
	if err != nil {
		return err
	}

	nextSubmitId := uint64(0)
	for {
		if s.interrupted(ctx) {
			return nil
		}

		if maxSubmitId < nextSubmitId {
			return ErrNoFileInfoToStat
		}

		endSubmitId := nextSubmitId + uint64(BatchGetSubmitsToStat) - 1
		submits, err := s.db.SubmitStore.QueryFinalizedAscWithCursor(&nextSubmitId, &endSubmitId)
		if err != nil {
			return err
		}
		if len(submits) == 0 {
			nextSubmitId = endSubmitId + 1
			continue
		}

		submitIdMapping := make(map[uint64][]uint64) // addressId => []submitId
		for _, submit := range submits {
			senderID := submit.SenderID
			submitID := submit.SubmissionIndex
			submitIdMapping[senderID] = append(submitIdMapping[senderID], submitID)
		}

		addresses := make([]store.Address, 0)
		submitIds := make([]uint64, 0)
		for addressId, subIds := range submitIdMapping {
			addresses = append(addresses, store.Address{ID: addressId, PrunedFiles: uint64(len(subIds))})
			submitIds = append(submitIds, subIds...)
			s.filePrunedTotal += uint64(len(subIds))
		}

		if err := s.db.DB.Transaction(func(dbTx *gorm.DB) error {
			if err := s.db.AddressStore.BatchDeltaUpsertPrunedFiles(dbTx, addresses); err != nil {
				return errors.WithMessage(err, "Failed to update addresses for stat pruned files")
			}
			if err := s.db.SubmitStore.UpdateByPrimaryKeys(dbTx, &store.Submit{Status: uint8(rpc.PrunedCounted)},
				submitIds); err != nil {
				return errors.WithMessage(err, "Failed to update submits for stat pruned files")
			}
			if err := s.db.ConfigStore.Upsert(dbTx, store.StatFilePrunedTotal,
				strconv.FormatUint(s.filePrunedTotal, 10)); err != nil {
				return errors.WithMessage(err, "Failed to update total count for stat pruned files")
			}
			return nil
		}); err != nil {
			return err
		}

		nextSubmitId = endSubmitId + 1
	}
}

func (s *StatSubmitPruned) interrupted(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}
