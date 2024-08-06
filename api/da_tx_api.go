package api

import (
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/gin-gonic/gin"
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
