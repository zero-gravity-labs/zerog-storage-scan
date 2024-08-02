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

	return convertDATxs(total, submits)
}

func convertDATxs(total int64, submits []store.DASubmit) (*DATxList, error) {
	daTxs := make([]DATxInfo, 0)

	for _, submit := range submits {
		var status uint8
		if submit.Verified {
			status = 1
		}

		daTxs = append(daTxs, DATxInfo{
			BlockNumber: submit.BlockNumber,
			TxHash:      submit.TxHash,
			Timestamp:   submit.BlockTime.Unix(),
			Epoch:       submit.Epoch,
			QuorumID:    submit.QuorumID,
			RootHash:    submit.RootHash,
			Status:      status,
		})
	}

	return &DATxList{
		Total: total,
		List:  daTxs,
	}, nil
}
