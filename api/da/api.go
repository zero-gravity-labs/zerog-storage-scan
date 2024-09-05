package da

import (
	scanApi "github.com/0glabs/0g-storage-scan/api"
	nhContract "github.com/0glabs/0g-storage-scan/contract"
	"github.com/0glabs/0g-storage-scan/docs"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	viperUtil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/gin-gonic/gin"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const BasePath = "/api"

var (
	sdk         *web3go.Client
	db          *store.MysqlStore
	chargeToken *TokenInfo
)

func MustInit(client *web3go.Client, store *store.MysqlStore) {
	sdk = client
	db = store

	var charge struct {
		Erc20TokenAddress string
		Symbol            string
		Decimals          uint8
	}
	viperUtil.MustUnmarshalKey("charge", &charge)

	if charge.Erc20TokenAddress != "" {
		name, symbol, decimals, err := nhContract.TokenInfo(client, charge.Erc20TokenAddress)
		if err != nil {
			logrus.WithError(err).Fatal("Get erc20 token info")
		}
		chargeToken = &TokenInfo{
			Address:  charge.Erc20TokenAddress,
			Name:     name,
			Symbol:   symbol,
			Decimals: decimals,
		}
	} else {
		chargeToken = &TokenInfo{
			Symbol:   charge.Symbol,
			Decimals: charge.Decimals,
		}
		chargeToken.Native = true
	}

	var flow struct {
		Address              string
		SubmitEventSignature string
	}
	viperUtil.MustUnmarshalKey("flow", &flow)
}

// @title			0G DA Scan API
// @version		1.0
// @description: Use any http client to fetch data from the 0G DA Scan
// @description.markdown
func init() {
	docs.SwaggerInfoda.BasePath = BasePath
}

func Register(router *gin.Engine) {
	apiRoute := router.Group(BasePath)

	daStatsRoute := apiRoute.Group("/stats")
	daStatsRoute.GET("storage", listDADataStatsHandler)
	daStatsRoute.GET("client", listDAClientStatsHandler)
	daStatsRoute.GET("signer", listDASignerStatsHandler)

	daTxsRoute := apiRoute.Group("/txs")
	daTxsRoute.GET("", listDATxsHandler)
	daTxsRoute.GET(":blockNumber/:epoch/:quorumID/:dataRoot", getDATxHandler)

	rewardsRoute := apiRoute.Group("/rewards")
	rewardsRoute.GET("", listRewardsHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName("da")))
}

// listTxsHandler godoc
//
//	@Summary		DA transaction list
//	@Description	Query DA transactions
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			skip		query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit		query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(100)	default(10)
//	@Param			rootHash	query		string	false	"The merkle root hash of the uploaded file"
//	@Param			txHash		query		string	false	"The layer1 tx hash of the submission"
//	@Success		200			{object}	api.BusinessError{Data=DATxList}
//	@Failure		600			{object}	api.BusinessError
//	@Router			/txs [get]
func listDATxsHandler(c *gin.Context) {
	api.Wrap(listDATxs)(c)
}

// getDATxHandler godoc
//
//	@Summary		DA transaction information
//	@Description	Query DA transaction by blockNumber, epoch, quorumId, dataRoot
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			blockNumber	path		string	true	"Block number at which the file is uploaded"
//	@Param			epoch		path		string	true	"The consecutive blocks in 0g chain is divided into groups of EpochBlocks and each group is an epoch"
//	@Param			quorumID	path		string	true	"Quorum id in an epoch"
//	@Param			dataRoot	path		string	true	"Data root"
//	@Success		200			{object}	api.BusinessError{Data=DATxInfo}
//	@Failure		600			{object}	api.BusinessError
//	@Router			/txs/{blockNumber}/{epoch}/{quorumID}/{dataRoot} [get]
func getDATxHandler(c *gin.Context) {
	api.Wrap(getDATx)(c)
}

// listDADataStatsHandler godoc
//
//	@Summary		DA data storage statistics
//	@Description	Query DA data storage statistics, including incremental and full data, and support querying at hourly or daily time intervals
//	@Tags			statistic
//	@Accept			json
//	@Produce		json
//	@Param			skip			query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit			query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(2000)	default(10)
//	@Param			minTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			maxTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			intervalType	query		string	false	"Statistics interval"	Enums(hour, day)	default(day)
//	@Param			sort			query		string	false	"Sort by timestamp"		Enums(asc, desc)	default(desc)
//	@Success		200				{object}	api.BusinessError{Data=DADataStatList}
//	@Failure		600				{object}	api.BusinessError
//	@Router			/stats/storage [get]
func listDADataStatsHandler(c *gin.Context) {
	api.Wrap(listDADataStat)(c)
}

// listDAClientStatsHandler godoc
//
//	@Summary		DA client statistics
//	@Description	Query DA client statistics, including incremental, active and full data, and support querying at hourly or daily time intervals
//	@Tags			statistic
//	@Accept			json
//	@Produce		json
//	@Param			skip			query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit			query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(2000)	default(10)
//	@Param			minTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			maxTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			intervalType	query		string	false	"Statistics interval"	Enums(hour, day)	default(day)
//	@Param			sort			query		string	false	"Sort by timestamp"		Enums(asc, desc)	default(desc)
//	@Success		200				{object}	api.BusinessError{Data=DAClientStatList}
//	@Failure		600				{object}	api.BusinessError
//	@Router			/stats/client [get]
func listDAClientStatsHandler(c *gin.Context) {
	api.Wrap(listDAClientStat)(c)
}

// listDAClientStatsHandler godoc
//
//	@Summary		DA signer statistics
//	@Description	Query DA signer statistics, including incremental, active and full data, and support querying at hourly or daily time intervals
//	@Tags			statistic
//	@Accept			json
//	@Produce		json
//	@Param			skip			query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit			query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(2000)	default(10)
//	@Param			minTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			maxTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			intervalType	query		string	false	"Statistics interval"	Enums(hour, day)	default(day)
//	@Param			sort			query		string	false	"Sort by timestamp"		Enums(asc, desc)	default(desc)
//	@Success		200				{object}	api.BusinessError{Data=DASignerStatList}
//	@Failure		600				{object}	api.BusinessError
//	@Router			/stats/signer [get]
func listDASignerStatsHandler(c *gin.Context) {
	api.Wrap(listDASignerStat)(c)
}

// listRewardsHandler godoc
//
//	@Summary		DA reward list
//	@Description	Query DA rewards
//	@Tags			reward
//	@Accept			json
//	@Produce		json
//	@Param			skip	query		int	false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit	query		int	false	"The number of records displayed on the page"								minimum(1)	maximum(100)	default(10)
//	@Success		200		{object}	api.BusinessError{Data=RewardList}
//	@Failure		600		{object}	api.BusinessError
//	@Router			/rewards [get]
func listRewardsHandler(c *gin.Context) {
	api.Wrap(listDARewards)(c)
}

func getAddressInfo(c *gin.Context) (*AddressInfo, error) {
	address := c.Param("address")
	if address == "" {
		logrus.Error("Failed to parse nil address")
		return nil, api.ErrValidation(errors.Errorf("Address is '%v'", address))
	}

	addressInfo, exist, err := db.AddressStore.Get(address)
	if err != nil {
		return nil, api.ErrInternal(err)
	}
	if !exist {
		return nil, scanApi.ErrNoMatchingRecords(errors.Errorf("Blockchain account, address %v", address))
	}

	return &AddressInfo{address, addressInfo.ID}, nil
}
