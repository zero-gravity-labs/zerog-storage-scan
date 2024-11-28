package cmd

import (
	"github.com/0glabs/0g-storage-scan/api/da"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/http/middlewares"
	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/spf13/cobra"
)

var (
	daAPICmd = &cobra.Command{
		Use:   "da_api",
		Short: "run rest da api server",
		Run:   startDAAPIService,
	}
)

func init() {
	rootCmd.AddCommand(daAPICmd)
}

func startDAAPIService(*cobra.Command, []string) {
	dataCtx := MustInitDataContext()
	defer dataCtx.Close()

	da.MustInit(dataCtx.Eth, dataCtx.DB)

	mws := httpMiddlewares(dataCtx)
	api.MustServeFromViper(da.Register, mws...)
	MustServeFromViper(da.Register, "daApi", mws...)
}

func MustServeFromViper(factory api.RouteFactory, cfgKey string, middlewares ...middlewares.Middleware) {
	var config api.Config
	viper.MustUnmarshalKey(cfgKey, &config)
	api.MustServe(config, factory, middlewares...)
}
