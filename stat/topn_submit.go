package stat

import (
	"strconv"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/sirupsen/logrus"
)

var (
	batchInSubmitIds = 10000
)

type TopnSubmit struct {
	*BaseStat
}

func MustNewTopnSubmit(cfg *StatConfig, db *store.MysqlStore, sdk *web3go.Client) *AbsTopn[StatRange] {
	baseStat := &BaseStat{
		Config: cfg,
		DB:     db,
		Sdk:    sdk,
	}

	topnSubmit := &TopnSubmit{
		BaseStat: baseStat,
	}
	topnSubmit.mustLoadLastPos()

	return &AbsTopn[StatRange]{
		Topn: topnSubmit,
	}
}

func (ts *TopnSubmit) mustLoadLastPos() {
	loaded, err := ts.loadLastPos(store.StatTopnSubmitId)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load stat pos from db")
	}

	// Submission index is set zero if not loaded.
	if !loaded {
		ts.currentPos = 0
	}
}

func (ts *TopnSubmit) nextStatRange() (*StatRange, error) {
	minPos := ts.currentPos
	maxPos := ts.currentPos + uint64(batchInSubmitIds) - 1

	block, err := ts.Sdk.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if err != nil {
		return nil, err
	}

	maxPosFinalized, exists, err := ts.DB.SubmitStore.MaxSubmissionIndexFinalized(block.Number.Uint64())
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrMaxPosFinalizedNotSync
	}

	if maxPosFinalized < minPos {
		return nil, ErrMinPosNotFinalized
	}
	if maxPosFinalized < maxPos {
		maxPos = maxPosFinalized
	}

	return &StatRange{minPos, maxPos}, nil
}

func (ts *TopnSubmit) calculateStat(r StatRange) error {
	groupedSubmits, err := ts.DB.SubmitStore.GroupBySender(r.minPos, r.maxPos)
	if err != nil {
		return err
	}

	for _, submit := range groupedSubmits {
		a := store.Address{
			ID:         submit.SenderID,
			DataSize:   submit.DataSize,
			StorageFee: submit.StorageFee,
			Txs:        submit.Txs,
			Files:      submit.Files,
			UpdatedAt:  submit.UpdatedAt,
		}
		err := ts.DB.AddressStore.IncreaseStatByPrimaryKey(nil, &a)
		if err != nil {
			return err
		}
	}

	if err := ts.DB.ConfigStore.Upsert(store.StatTopnSubmitId, strconv.FormatUint(r.maxPos, 10)); err != nil {
		return err
	}

	ts.currentPos = r.maxPos + 1

	return nil
}
