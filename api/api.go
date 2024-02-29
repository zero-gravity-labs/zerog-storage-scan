package api

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	viperutil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/gin-gonic/gin"
	"github.com/openweb3/web3go"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	nhContract "github.com/zero-gravity-labs/zerog-storage-scan/contract"
	"github.com/zero-gravity-labs/zerog-storage-scan/docs"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
)

const BasePath = "/api"

var (
	sdk           *web3go.Client
	db            *store.MysqlStore
	chargeToken   *TokenInfo
	flowAddr      string
	flowSubmitSig string
)

func MustInit(client *web3go.Client, store *store.MysqlStore) {
	sdk = client
	db = store

	var charge struct {
		Erc20TokenAddress string
	}
	viperutil.MustUnmarshalKey("charge", &charge)

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

	var flow struct {
		Address              string
		SubmitEventSignature string
	}
	viperutil.MustUnmarshalKey("flow", &flow)
	flowAddr = flow.Address
	flowSubmitSig = flow.SubmitEventSignature
}

// @title		ZeroGStorage Scan API
// @version		1.0
// @description	Use any http client to fetch data from the ZeroGStorage Scan.
func init() {
	docs.SwaggerInfo.BasePath = BasePath
}

// dashboardHandler godoc
//
//	@Summary		Statistics dashboard
//	@Description	Query statistics dashboard includes `average uplink rate` and `storage base fee`
//	@Tags			statistic
//	@Produce		json
//	@Success		200	{object}	api.BusinessError{Data=Dashboard}
//	@Failure		600	{object}	api.BusinessError
//	@Router			/statistic/dashboard [get]
func dashboardHandler(c *gin.Context) {
	api.Wrap(dashboard)(c)
}

// listTxStatHandler godoc
//
//	@Summary		Transaction statistics
//	@Description	Query transaction statistics, including incremental and full data, and support querying at hourly or daily time intervals
//	@Tags			statistic
//	@Accept			json
//	@Produce		json
//	@Param			skip			query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit			query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(2000)	default(10)
//	@Param			minTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			maxTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			intervalType	query		string	false	"Statistics interval"	Enums(hour, day)	default(day)
//	@Param			sort			query		string	false	"Sort by timestamp"		Enums(asc, desc)	default(desc)
//	@Success		200				{object}	api.BusinessError{Data=TxStatList}
//	@Failure		600				{object}	api.BusinessError
//	@Router			/statistic/transaction/list [get]
func listTxStatHandler(c *gin.Context) {
	api.Wrap(listTxStat)(c)
}

// listDataStatHandler godoc
//
//	@Summary		Data storage statistics
//	@Description	Query data storage statistics, including incremental and full data, and support querying at hourly or daily time intervals
//	@Tags			statistic
//	@Accept			json
//	@Produce		json
//	@Param			skip			query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit			query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(2000)	default(10)
//	@Param			minTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			maxTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			intervalType	query		string	false	"Statistics interval"	Enums(hour, day)	default(day)
//	@Param			sort			query		string	false	"Sort by timestamp"		Enums(asc, desc)	default(desc)
//	@Success		200				{object}	api.BusinessError{Data=DataStatList}
//	@Failure		600				{object}	api.BusinessError
//	@Router			/statistic/storage/list [get]
func listDataStatHandler(c *gin.Context) {
	api.Wrap(listDataStat)(c)
}

// listFeeStatHandler godoc
//
//	@Summary		fee statistics
//	@Description	Query fee statistics, including incremental and full data, and support querying at hourly or daily time intervals
//	@Tags			statistic
//	@Accept			json
//	@Produce		json
//	@Param			skip			query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit			query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(2000)	default(10)
//	@Param			minTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			maxTimestamp	query		int		false	"Timestamp in seconds"
//	@Param			intervalType	query		string	false	"Statistics interval"	Enums(hour, day)	default(day)
//	@Param			sort			query		string	false	"Sort by timestamp"		Enums(asc, desc)	default(desc)
//	@Success		200				{object}	api.BusinessError{Data=FeeStatList}
//	@Failure		600				{object}	api.BusinessError
//	@Router			/statistic/fee/list [get]
func listFeeStatHandler(c *gin.Context) {
	api.Wrap(listFeeStat)(c)
}

// listTxHandler godoc
//
//	@Summary		Layer2 transaction list
//	@Description	Query layer2 transactions, support address and root hash filter
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			skip		query		int		false	"The number of skipped records, usually it's pageSize * (pageNumber - 1)"	minimum(0)	default(0)
//	@Param			limit		query		int		false	"The number of records displayed on the page"								minimum(1)	maximum(100)	default(10)
//	@Param			address		query		string	false	"The submitter address of the uploaded file"
//	@Param			rootHash	query		string	false	"The merkle root hash of the uploaded file"
//	@Success		200			{object}	api.BusinessError{Data=TxList}
//	@Failure		600			{object}	api.BusinessError
//	@Router			/transaction/list [get]
func listTxHandler(c *gin.Context) {
	api.Wrap(listTx)(c)
}

// getTxBriefHandler godoc
//
//	@Summary		Layer2 transaction overview
//	@Description	Query layer2 transaction overview by txSeq
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			txSeq	query		string	true	"Lay2 transaction sequence number"
//	@Success		200		{object}	api.BusinessError{Data=TxBrief}
//	@Failure		600		{object}	api.BusinessError
//	@Router			/transaction/brief [get]
func getTxBriefHandler(c *gin.Context) {
	api.Wrap(getTxBrief)(c)
}

// getTxDetailHandler godoc
//
//	@Summary		Layer2 transaction advanced info
//	@Description	Query layer2 transaction advanced info by txSeq
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			txSeq	query		string	true	"Lay2 transaction sequence number"
//	@Success		200		{object}	api.BusinessError{Data=TxDetail}
//	@Failure		600		{object}	api.BusinessError
//	@Router			/transaction/detail [get]
func getTxDetailHandler(c *gin.Context) {
	api.Wrap(getTxDetail)(c)
}

func RegisterRouter(router *gin.Engine) {
	apiRoute := router.Group(BasePath)

	statRoute := apiRoute.Group("/statistic")
	statRoute.GET("dashboard", dashboardHandler)
	statRoute.GET("transaction/list", listTxStatHandler)
	statRoute.GET("storage/list", listDataStatHandler)
	statRoute.GET("fee/list", listFeeStatHandler)

	txRoute := apiRoute.Group("/transaction")
	txRoute.GET("list", listTxHandler)
	txRoute.GET("brief", getTxBriefHandler)
	txRoute.GET("detail", getTxDetailHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
