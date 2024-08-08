package api

import (
	"strconv"

	"github.com/0glabs/0g-storage-scan/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func listDATxs(c *gin.Context) (interface{}, error) {
	var param listDATxParam
	if err := c.ShouldBind(&param); err != nil {
		return nil, err
	}

	total, submits, err := db.DASubmitStore.List(param.RootHash, param.TxHash, param.isDesc(), param.Skip, param.Limit)
	if err != nil {
		return nil, err
	}
	if len(submits) == 0 {
		return &DATxList{
			Total: total,
			List:  make([]DATxInfo, 0),
		}, nil
	}

	return convertDATxs(total, submits)
}

func getDATx(c *gin.Context) (interface{}, error) {
	blockNumberParam := c.Param("blockNumber")
	blockNumber, err := strconv.ParseUint(blockNumberParam, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse blockNumber")
		return nil, errors.Errorf("Biz error, invalid blockNumber %v", blockNumberParam)
	}
	epochParam := c.Param("epoch")
	epoch, err := strconv.ParseUint(epochParam, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse epoch")
		return nil, errors.Errorf("Biz error, invalid epoch %v", epochParam)
	}
	quorumIDParam := c.Param("quorumID")
	quorumID, err := strconv.ParseUint(quorumIDParam, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse quorumID")
		return nil, errors.Errorf("Biz error, invalid quorumID %v", quorumIDParam)
	}
	dataRoot := c.Param("dataRoot")

	var submit store.DASubmit
	exist, err := db.Store.Exists(&submit, "block_number=? and epoch=? and quorum_id=? and root_hash=?",
		blockNumber, epoch, quorumID, dataRoot)
	if err != nil {
		logrus.WithError(err).Error("Failed to query databases")
		return nil, errors.Errorf("Biz error, blockNumber %v blockNumber %v blockNumber %v blockNumber %v",
			blockNumber, epoch, quorumID, dataRoot)
	}
	if !exist {
		return nil, errors.Errorf("Record not found, blockNumber %v blockNumber %v blockNumber %v blockNumber %v",
			blockNumber, epoch, quorumID, dataRoot)
	}

	addrIDs := []uint64{submit.SenderID}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, err
	}

	var status uint8
	method := "submit"
	if submit.Verified {
		status = 1
		method = "submitVerified"
	}
	return DATxInfo{
		BlockNumber: submit.BlockNumber,
		TxHash:      submit.TxHash,
		Timestamp:   submit.BlockTime.Unix(),
		From:        addrMap[submit.SenderID].Address,
		Method:      method,

		Epoch:      submit.Epoch,
		QuorumID:   submit.QuorumID,
		RootHash:   submit.RootHash,
		StorageFee: submit.BlobPrice,
		Status:     status,
	}, nil
}

func convertDATxs(total int64, submits []store.DASubmit) (*DATxList, error) {
	addrIDs := make([]uint64, 0)
	for _, submit := range submits {
		addrIDs = append(addrIDs, submit.SenderID)
	}
	addrMap, err := db.BatchGetAddresses(addrIDs)
	if err != nil {
		return nil, err
	}

	daTxs := make([]DATxInfo, 0)
	for _, submit := range submits {
		var status uint8
		method := "submit"
		if submit.Verified {
			status = 1
			method = "submitVerified"
		}

		daTxs = append(daTxs, DATxInfo{
			BlockNumber: submit.BlockNumber,
			TxHash:      submit.TxHash,
			Timestamp:   submit.BlockTime.Unix(),
			From:        addrMap[submit.SenderID].Address,
			Method:      method,

			Epoch:      submit.Epoch,
			QuorumID:   submit.QuorumID,
			RootHash:   submit.RootHash,
			StorageFee: submit.BlobPrice,
			Status:     status,
		})
	}

	return &DATxList{
		Total: total,
		List:  daTxs,
	}, nil
}
