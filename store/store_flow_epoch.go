package store

import (
	"time"

	"github.com/0glabs/0g-storage-client/contract"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"gorm.io/gorm"
)

type FlowEpoch struct {
	SubmissionIndex uint64 `gorm:"primaryKey;autoIncrement:false"`

	Sender          string `gorm:"-"`
	SenderID        uint64 `gorm:"not null"`
	Index           uint64 `gorm:"not null"`
	StartMerkleRoot string `gorm:"size:66;not null"`
	FlowLength      uint64 `gorm:"not null"`
	Context         string `gorm:"size:66;not null"`

	BlockNumber uint64    `gorm:"not null;index:idx_bn"`
	BlockTime   time.Time `gorm:"not null;index:idx_bt"`
	TxHash      string    `gorm:"size:66;not null;index:idx_txHash,length:10"`
}

func NewFlowEpoch(blockTime time.Time, log types.Log, filter *contract.FlowFilterer) (*FlowEpoch, error) {
	flowEpoch, err := filter.ParseNewEpoch(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	epoch := &FlowEpoch{
		SubmissionIndex: flowEpoch.SubmissionIndex.Uint64(),
		Sender:          flowEpoch.Sender.String(),
		Index:           flowEpoch.Index.Uint64(),
		StartMerkleRoot: common.Hash(flowEpoch.StartMerkleRoot[:]).String(),
		FlowLength:      flowEpoch.FlowLength.Uint64(),
		Context:         common.Hash(flowEpoch.Context[:]).String(),

		BlockNumber: log.BlockNumber,
		BlockTime:   blockTime,
		TxHash:      log.TxHash.String(),
	}

	return epoch, nil
}

func (FlowEpoch) TableName() string {
	return "flow_epochs"
}

type FlowEpochStore struct {
	*mysql.Store
}

func newFlowEpochStore(db *gorm.DB) *FlowEpochStore {
	return &FlowEpochStore{
		Store: mysql.NewStore(db),
	}
}

func (ss *FlowEpochStore) Add(dbTx *gorm.DB, flowEpochs []FlowEpoch) error {
	return dbTx.CreateInBatches(flowEpochs, batchSizeInsert).Error
}

func (ss *FlowEpochStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&FlowEpoch{}).Error
}
