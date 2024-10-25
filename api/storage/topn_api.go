package storage

import (
	"time"

	scanApi "github.com/0glabs/0g-storage-scan/api"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	dataSizeTopn   = "data_size"
	storageFeeTopn = "storage_fee"
	txsTopn        = "txs"
	filesTopn      = "files"

	maxRecords = 10
)

var (
	spanTypes = map[string]time.Duration{
		"24h": time.Hour * 24,
		"3d":  time.Hour * 24 * 3,
		"7d":  time.Hour * 24 * 7,
	}
)

func topnDataSize(c *gin.Context) (interface{}, error) {
	return topnByType(c, dataSizeTopn)
}

func topnStorageFee(c *gin.Context) (interface{}, error) {
	return topnByType(c, storageFeeTopn)
}

func topnTxs(c *gin.Context) (interface{}, error) {
	return topnByType(c, txsTopn)
}

func topnFiles(c *gin.Context) (interface{}, error) {
	return topnByType(c, filesTopn)
}

func topnByType(c *gin.Context, t string) (interface{}, error) {
	var topnP topnParam
	if err := c.ShouldBind(&topnP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid topn query param"))
	}

	statSpan := spanTypes[topnP.SpanType]
	records, err := db.AddressStore.Topn(t, statSpan, maxRecords)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get submit topn list"))
	}

	result := make(map[string]interface{})
	switch t {
	case dataSizeTopn:
		list := make([]DataTopn, 0)
		for _, r := range records {
			list = append(list, DataTopn{
				Address:  r.Address,
				DataSize: r.DataSize,
			})
		}
		result["list"] = list
	case storageFeeTopn:
		list := make([]FeeTopn, 0)
		for _, r := range records {
			list = append(list, FeeTopn{
				Address:    r.Address,
				StorageFee: r.StorageFee,
			})
		}
		result["list"] = list
	case txsTopn:
		list := make([]TxsTopn, 0)
		for _, r := range records {
			list = append(list, TxsTopn{
				Address: r.Address,
				Txs:     r.Txs,
			})
		}
		result["list"] = list
	case filesTopn:
		list := make([]FilesTopn, 0)
		for _, r := range records {
			list = append(list, FilesTopn{
				Address: r.Address,
				Files:   r.Files,
			})
		}
		result["list"] = list
	default:
		return nil, api.ErrValidation(errors.Errorf("Invalid topn type %v", t))
	}

	return result, nil
}
