package cmd

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/0glabs/0g-storage-client/node"
	"github.com/0glabs/0g-storage-scan/api/middlewares/metrics"
	nhRate "github.com/0glabs/0g-storage-scan/api/middlewares/rate"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/Conflux-Chain/go-conflux-util/http/middlewares"
	"github.com/Conflux-Chain/go-conflux-util/rate/http"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/Conflux-Chain/go-conflux-util/viper"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go"
	"github.com/sirupsen/logrus"
)

// DataContext context to hold sdk clients for blockchain interoperation.
type DataContext struct {
	Eth      *web3go.Client
	L2Sdks   []*node.ZgsClient
	DB       *store.MysqlStore
	EthCfg   SdkConfig
	L2SdkCfg L2SdkConfig
}

type SdkConfig struct {
	URL             string
	Retry           int
	RetryInterval   time.Duration `default:"1s"`
	RequestTimeout  time.Duration `default:"3s"`
	MaxConnsPerHost int           `default:"1024"`
	AlertChannel    string
	HealthReport    health.TimedCounterConfig
}

type L2SdkConfig struct {
	URLs            []string
	Retry           int
	RetryInterval   time.Duration `default:"1s"`
	RequestTimeout  time.Duration `default:"3s"`
	MaxConnsPerHost int           `default:"1024"`
	AlertChannel    string
	HealthReport    health.TimedCounterConfig
}

var migrationModels = []interface{}{
	&store.Address{},
	&store.Block{},
	&store.Config{},
	&store.Submit{},
	&store.AddressSubmit{},
	&store.SubmitStat{},
	&store.SubmitTopnStat{},
	&store.Reward{},
	&store.RewardStat{},
	&store.RewardTopnStat{},
	&store.AddressReward{},
	&store.AddressStat{},
	&store.Miner{},
	&store.MinerStat{},
	&store.MinerRegister{},
	&store.FlowEpoch{},
	&store.DASigner{},
	&store.DASignerStat{},
	&store.DASubmit{},
	&store.DASubmitStat{},
	&store.DAReward{},
	&store.DAClient{},
	&store.DAClientStat{},
	&store.RateLimit{},
}

func MustInitDataContext() DataContext {
	cfg := mysql.MustNewConfigFromViper()
	db := cfg.MustOpenOrCreate()
	if err := db.AutoMigrate(migrationModels...); err != nil {
		logrus.WithError(err).Fatalln("failed to migrate database")
	}

	sdkCfg := SdkConfig{}
	viper.MustUnmarshalKey("eth", &sdkCfg)
	opt := web3go.ClientOption{}
	opt.WithRetry(sdkCfg.Retry, sdkCfg.RetryInterval).
		WithTimout(sdkCfg.RequestTimeout).
		WithMaxConnectionPerHost(sdkCfg.MaxConnsPerHost)
	eth := web3go.MustNewClientWithOption(sdkCfg.URL, opt)

	l2SdkCfg := L2SdkConfig{}
	viper.MustUnmarshalKey("storage", &l2SdkCfg)
	opt2 := providers.Option{}
	opt2.WithRetry(l2SdkCfg.Retry, l2SdkCfg.RetryInterval).
		WithTimout(l2SdkCfg.RequestTimeout).
		WithMaxConnectionPerHost(l2SdkCfg.MaxConnsPerHost)
	l2Sdks := node.MustNewZgsClients(l2SdkCfg.URLs, opt2)

	return DataContext{
		DB:       store.MustNewStore(db, cfg),
		L2Sdks:   l2Sdks,
		Eth:      eth,
		EthCfg:   sdkCfg,
		L2SdkCfg: l2SdkCfg,
	}
}

func (ctx *DataContext) Close() {
	if ctx.DB != nil {
		ctx.DB.Close()
	}

	if ctx.Eth != nil {
		ctx.Eth.Close()
	}
}

func GracefulShutdown(wg *sync.WaitGroup, cancel context.CancelFunc) {
	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	// Wait for SIGTERM to be captured
	<-termChan
	logrus.Info("SIGTERM/SIGINT received, shutdown process initiated")

	// Cancel to notify active goroutines to clean up.
	cancel()

	logrus.Info("Waiting for shutdown...")
	wg.Wait()

	logrus.Info("Shutdown gracefully")
}

func httpMiddlewares(dataCtx DataContext) []middlewares.Middleware {
	mws := make([]middlewares.Middleware, 0)
	mws = append(mws, middlewares.RealIP)
	mws = append(mws, metrics.URLType)
	mws = append(mws, middlewares.NewApiKeyMiddleware(middlewares.ApiKeyOption{ParamName: "apikey"}))

	limiterFactory := nhRate.NewLimiterFactory(nhRate.NewLimitKeyLoader(dataCtx.DB.ListLimitKeyInfos))
	go limiterFactory.AutoReload(10*time.Second, dataCtx.DB.LoadRateLimitConfigs)
	mws = append(mws, nhRate.NewAPIRateMiddleware(limiterFactory.Limit))
	mws = append(mws, http.NewHttpMiddleware(limiterFactory.Limit, "api_all_qps"))
	mws = append(mws, http.NewHttpMiddleware(limiterFactory.Limit, "api_all_daily"))

	mws = append(mws, metrics.Metrics())

	return mws
}
