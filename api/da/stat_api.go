package da

import (
	scanApi "github.com/0glabs/0g-storage-scan/api"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func listDADataStat(c *gin.Context) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid stat query param"))
	}

	total, records, err := db.DASubmitStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get da submit stat list"))
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]DADataStat, 0)
	for _, r := range records {
		list = append(list, DADataStat{
			StatTime:        r.StatTime,
			BlobNew:         r.BlobNew,
			BlobTotal:       r.BlobTotal,
			DataSizeNew:     r.DataSizeNew,
			DataSizeTotal:   r.DataSizeTotal,
			StorageFeeNew:   r.StorageFeeNew,
			StorageFeeTotal: r.StorageFeeTotal,
		})
	}
	result["list"] = list

	return result, nil
}

func listDAClientStat(c *gin.Context) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid stat query param"))
	}

	total, records, err := db.DAClientStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get da client stat list"))
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]DAClientStat, 0)
	for _, r := range records {
		list = append(list, DAClientStat{
			StatTime:     r.StatTime,
			ClientNew:    r.ClientNew,
			ClientActive: r.ClientActive,
			ClientTotal:  r.ClientTotal,
		})
	}
	result["list"] = list

	return result, nil
}

func listDASignerStat(c *gin.Context) (interface{}, error) {
	var statP statParam
	if err := c.ShouldBind(&statP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid stat query param"))
	}

	total, records, err := db.DASignerStatStore.List(&statP.IntervalType, statP.MinTimestamp, statP.MaxTimestamp,
		statP.isDesc(), statP.Skip, statP.Limit)
	if err != nil {
		return nil, scanApi.ErrDatabase(errors.WithMessage(err, "Failed to get da signer stat list"))
	}

	result := make(map[string]interface{})
	result["total"] = total

	list := make([]DASignerStat, 0)
	for _, r := range records {
		list = append(list, DASignerStat{
			StatTime:     r.StatTime,
			SignerNew:    r.SignerNew,
			SignerActive: r.SignerActive,
			SignerTotal:  r.SignerTotal,
		})
	}
	result["list"] = list

	return result, nil
}
