package api

import (
	"encoding/json"
	"strconv"

	"github.com/0glabs/0g-storage-scan/store"
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func listSubmits(addressID *uint64, rootHash *string, idDesc bool, skip, limit int) (int64, []store.Submit, error) {
	if addressID == nil {
		return db.SubmitStore.List(rootHash, idDesc, skip, limit)
	}

	total, addrSubmits, err := db.AddressSubmitStore.List(addressID, rootHash, idDesc, skip, limit)
	if err != nil {
		return 0, nil, err
	}

	submits := make([]store.Submit, 0)
	for _, as := range addrSubmits {
		submits = append(submits, store.Submit{
			SubmissionIndex: as.SubmissionIndex,
			RootHash:        as.RootHash,
			SenderID:        as.SenderID,
			Length:          as.Length,
			BlockNumber:     as.BlockNumber,
			BlockTime:       as.BlockTime,
			TxHash:          as.TxHash,
			Status:          as.Status,
			TotalSegNum:     as.TotalSegNum,
			UploadedSegNum:  as.UploadedSegNum,
			Fee:             as.Fee,
		})
	}

	return total, submits, nil
}

func convertTxList(total int64, submits []store.Submit) (*StorageTxList, error) {
	addrIDs := make([]uint64, 0)
	for _, submit := range submits {
		addrIDs = append(addrIDs, submit.SenderID)
	}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, err
	}

	storageTxs := make([]StorageTxInfo, 0)
	for _, submit := range submits {
		storageTx := StorageTxInfo{
			TxSeq:            submit.SubmissionIndex,
			BlockNumber:      submit.BlockNumber,
			TxHash:           submit.TxHash,
			RootHash:         submit.RootHash,
			From:             addrMap[submit.SenderID].Address,
			Method:           "submit",
			Status:           submit.Status,
			Segments:         submit.TotalSegNum,
			UploadedSegments: submit.UploadedSegNum,
			Timestamp:        submit.BlockTime.Unix(),
			DataSize:         submit.Length,
			StorageFee:       submit.Fee,
		}
		storageTxs = append(storageTxs, storageTx)
	}

	return &StorageTxList{
		Total: total,
		List:  storageTxs,
	}, nil
}

func listAddressStorageTx(c *gin.Context) (interface{}, error) {
	address := c.Param("address")
	if address == "" {
		logrus.Error("Failed to parse nil address")
		return nil, errors.Errorf("Biz error, nil address %v", address)
	}
	addr, exist, err := db.AddressStore.Get(address)
	if err != nil {
		return nil, commonApi.ErrInternal(err)
	}
	if !exist {
		return TxList{}, nil
	}
	addrIDPtr := &addr.ID

	var param listAddressStorageTxParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	total, submits, err := listSubmits(addrIDPtr, param.RootHash, param.isDesc(), param.Skip, param.Limit)
	if err != nil {
		return nil, err
	}

	return convertTxList(total, submits)
}

func listStorageTx(c *gin.Context) (interface{}, error) {
	var param listStorageTxParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	total, submits, err := listSubmits(nil, nil, param.isDesc(), param.Skip, param.Limit)
	if err != nil {
		return nil, err
	}

	return convertTxList(total, submits)
}

func getStorageTx(c *gin.Context) (interface{}, error) {
	txSeqParam := c.Param("txSeq")
	txSeq, err := strconv.ParseUint(txSeqParam, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse txSeq")
		return nil, errors.Errorf("Biz error, invalid txSeq %v", txSeq)
	}

	var submit store.Submit
	exist, err := db.Store.Exists(&submit, "submission_index = ?", txSeq)
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, errors.Errorf("Biz error, txSeq %v", txSeq)
	}
	if !exist {
		return nil, errors.Errorf("Record not found, txSeq %v", txSeq)
	}

	addrIDs := []uint64{submit.SenderID}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, err
	}

	result := StorageTxDetail{
		TxSeq:       strconv.FormatUint(submit.SubmissionIndex, 10),
		From:        addrMap[submit.SenderID].Address,
		Method:      "submit",
		RootHash:    submit.RootHash,
		DataSize:    submit.Length,
		Status:      submit.Status,
		StorageFee:  submit.Fee,
		BlockNumber: submit.BlockNumber,
		TxHash:      submit.TxHash,
		Timestamp:   uint64(submit.BlockTime.Unix()),
	}

	var extra store.SubmitExtra
	if err := json.Unmarshal(submit.Extra, &extra); err != nil {
		logrus.WithError(err).Error("Failed to unmarshal submit extra")
		return nil, errors.Errorf("Unmarshal submit extra error, txSeq %v", txSeq)
	}
	result.StartPosition = extra.StartPos.Uint64()
	result.EndPosition = extra.StartPos.Uint64() + submit.Length
	result.Segments = submit.TotalSegNum

	hash := common.HexToHash(submit.TxHash)
	tx, err := sdk.Eth.TransactionByHash(hash)
	if err != nil {
		logrus.WithError(err).WithField("txSeq", txSeq).Error("Failed to get transaction")
		return nil, errors.Errorf("Get tx error, txSeq %v", txSeq)
	}
	rcpt, err := sdk.Eth.TransactionReceipt(hash)
	if err != nil {
		logrus.WithError(err).WithField("txSeq", txSeq).Error("Failed to get receipt")
		return nil, errors.Errorf("Get receitp error, txSeq %v", txSeq)
	}
	result.GasFee = tx.GasPrice.Uint64() * rcpt.GasUsed
	result.GasUsed = rcpt.GasUsed
	result.GasLimit = tx.Gas

	return result, nil
}
