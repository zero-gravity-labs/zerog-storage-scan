package api

import (
	"encoding/hex"
	"encoding/json"
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

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
		conds = append(conds, SenderID(addr.Id))
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
	for _, submit := range *submits {
		addrIds = append(addrIds, submit.SenderID)
	}
	addrMap, err := db.BatchGetAddresses(addrIds)
	if err != nil {
		return nil, err
	}

	storageTxs := make([]StorageTx, 0)
	for _, submit := range *submits {
		storageTx := StorageTx{
			TxSeq:     submit.SubmissionIndex,
			BlockNum:  submit.BlockNumber,
			TxHash:    submit.TxHash,
			RootHash:  submit.RootHash,
			Address:   addrMap[submit.SenderID].Address,
			Method:    "submit",
			Timestamp: submit.BlockTime.Unix(),
			DataSize:  submit.Length,
			BaseFee:   submit.Value,
		}
		if submit.Finalized {
			storageTx.Status = 1
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

	var submit store.Submit
	exist, err := db.Store.Exists(&submit, "submission_index = ?", param.TxSeq)
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, errors.Errorf("Biz error, txSeq %v", *param.TxSeq)
	}
	if !exist {
		return nil, errors.Errorf("Record not found, txSeq %v", *param.TxSeq)
	}

	addrIds := []uint64{submit.SenderID}
	addrMap, err := db.BatchGetAddresses(addrIds)
	if err != nil {
		return nil, err
	}

	result := TxBrief{
		TxSeq:    strconv.FormatUint(submit.SubmissionIndex, 10),
		From:     addrMap[submit.SenderID].Address,
		Method:   "submit",
		RootHash: submit.RootHash,
		DataSize: submit.Length,
		CostInfo: &CostInfo{
			TokenInfo: *chargeToken,
			BasicCost: submit.Value.String(),
		},
		BlockNumber: submit.BlockNumber,
		TxHash:      submit.TxHash,
		Timestamp:   uint64(submit.BlockTime.Unix()),
	}
	if submit.Finalized {
		result.Status = 1
	}

	hash := common.HexToHash(submit.TxHash)
	tx, err := sdk.Eth.TransactionByHash(hash)
	if err != nil {
		logrus.WithError(err).WithField("txSeq", param.TxSeq).Error("Failed to get transaction")
		return nil, errors.Errorf("Get tx error, txSeq %v", param.TxSeq)
	}
	rcpt, err := sdk.Eth.TransactionReceipt(hash)
	if err != nil {
		logrus.WithError(err).WithField("txSeq", param.TxSeq).Error("Failed to get receipt")
		return nil, errors.Errorf("Get receitp error, txSeq %v", param.TxSeq)
	}
	result.GasFee = tx.GasPrice.Uint64() * rcpt.GasUsed
	result.GasUsed = rcpt.GasUsed
	result.GasLimit = tx.Gas

	return result, nil
}

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

	var extra store.SubmitExtra
	if err := json.Unmarshal(submit.Extra, &extra); err != nil {
		logrus.WithError(err).Error("Failed to unmarshal submit extra")
		return nil, errors.Errorf("Unmarshal submit extra error, txSeq %v", *param.TxSeq)
	}

	var nodes []SubmissionNode
	for _, n := range extra.Submission.Nodes {
		nodes = append(nodes, SubmissionNode{
			Root:   "0x" + hex.EncodeToString(n.Root[:]),
			Height: n.Height,
		})
	}

	result := TxDetail{
		TxSeq:       strconv.FormatUint(submit.SubmissionIndex, 10),
		RootHash:    submit.RootHash,
		StartPos:    extra.StartPos.Uint64(),
		EndPos:      extra.StartPos.Uint64() + submit.Length,
		PieceCounts: uint64(len(nodes)),
		Pieces:      nodes,
	}

	return result, nil
}

func SenderID(si uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("sender_id = ?", si)
	}
}

func RootHash(rh string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("root_hash = ?", strings.ToLower(strings.TrimPrefix(rh, "0x")))
	}
}
