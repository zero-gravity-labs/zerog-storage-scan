package api

import (
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// TO add logic when refactor submit db domain
func listTx(c *gin.Context) (interface{}, error) {
	var param listTxParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	dbRaw := db.DB.Model(&store.Submit{})
	var conds []func(db *gorm.DB) *gorm.DB
	if param.Address != "" {
		addr, exist, err := db.AddressStore.Get(param.Address)
		if err != nil {
			return nil, commonApi.ErrInternal(err)
		}
		if !exist {
			return TxList{}, nil
		}
		conds = append(conds, SenderId(addr.Id))
	}
	if param.RootHash != "" {
		conds = append(conds, RootHash(param.RootHash))
	}
	dbRaw.Scopes(conds...)

	submits := new([]store.Submit)
	total, err := db.List(dbRaw, true, param.Skip, param.Limit, submits)
	if err != nil {
		return nil, err
	}

	addrIds := make([]uint64, 0)
	txHashes := make([]string, 0)
	for _, submit := range *submits {
		addrIds = append(addrIds, submit.SenderID)
		txHashes = append(txHashes, submit.TxHash)
	}
	addrMap, err := db.BatchGetAddresses(addrIds)
	if err != nil {
		return nil, err
	}

	storageTxs := make([]StorageTx, 0)
	for _, submit := range *submits {
		storageTx := StorageTx{
			TxSeq:    submit.SubmissionIndex,
			BlockNum: submit.BlockNumber,
			TxHash:   "0x" + submit.TxHash,
			RootHash: "0x" + submit.RootHash,
			Address:  addrMap[submit.SenderID].Address,
			Method:   "submit",
		}
		storageTxs = append(storageTxs, storageTx)
	}

	result := TxList{
		Total: total,
		List:  storageTxs,
	}
	return result, nil
}

func getTxBrief(c *gin.Context) (interface{}, error) {
	var param queryTxParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	var submit struct {
		SubmissionIndex  uint64
		SubmissionLength uint64
		SenderId         uint64
		RootHash         string
		BlockNumber      uint64
		Hash             string
		CreatedAt        *time.Time
		Value            *decimal.Decimal
		Status           uint64
		GasFee           uint64
		GasUsed          uint64
		GasLimit         uint64
	}

	err := db.DB.Raw(`select s.submission_index, s.submission_length, s.sender_id, s.root_hash, t.block_number, 
       t.hash, t.created_at, t.status, t.gas_fee, t.gas_used, t.gas_limit, ts.value from submits s 
       left join txs t on s.tx_hash = t.hash 
       left join erc20_transfers ts on t.hash = ts.tx_hash 
       where s.submission_index =?`, param.TxSeq).Take(&submit).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Errorf("Record not found, txSeq %v", *param.TxSeq)
	}
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, errors.Errorf("Biz error, txSeq %v", *param.TxSeq)
	}

	addrIds := []uint64{submit.SenderId}
	addrMap, err := db.BatchGetAddresses(addrIds)
	if err != nil {
		return nil, err
	}

	result := TxBrief{
		TxSeq:    strconv.FormatUint(submit.SubmissionIndex, 10),
		From:     addrMap[submit.SenderId].Address,
		Method:   "submit",
		RootHash: "0x" + submit.RootHash,
		DataSize: submit.SubmissionLength,
		CostInfo: &CostInfo{
			TokenInfo: *chargeToken,
			BasicCost: submit.Value.String(),
		},
		BlockNumber: submit.BlockNumber,
		TxHash:      "0x" + submit.Hash,
		Timestamp:   uint64(submit.CreatedAt.Unix()),
		Status:      submit.Status,
		GasFee:      submit.GasFee,
		GasUsed:     submit.GasUsed,
		GasLimit:    submit.GasLimit,
	}

	return result, nil
}

// TODO add StartPos, EndPos and PieceCounts when refactor api module
func getTxDetail(c *gin.Context) (interface{}, error) {
	var param queryTxParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	var submit store.Submit
	exist, err := db.Exists(&submit, "submission_index = ?", param.TxSeq)
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, errors.Errorf("Biz error, txSeq %v", *param.TxSeq)
	}
	if !exist {
		return nil, errors.Errorf("Record not found, txSeq %v", *param.TxSeq)
	}

	nodes, err := getSubmitEvent("0x" + submit.TxHash)
	if err != nil {
		return nil, err
	}

	result := TxDetail{
		TxSeq:    strconv.FormatUint(submit.SubmissionIndex, 10),
		RootHash: "0x" + submit.RootHash,
		//StartPos:    submit.StartPos,
		//EndPos:      submit.StartPos + submit.Length,
		//PieceCounts: submit.Nodes,
		Pieces: nodes,
	}

	return result, nil
}

func SenderId(si uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("sender_id = ?", si)
	}
}

func RootHash(rh string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("root_hash = ?", strings.ToLower(strings.TrimPrefix(rh, "0x")))
	}
}
