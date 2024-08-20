package cmd

import (
	"github.com/0glabs/0g-storage-scan/api/storage"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/spf13/cobra"
)

var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "run rest api server",
		Run:   startAPIService,
	}
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startAPIService(*cobra.Command, []string) {
	dataCtx := MustInitDataContext()
	defer dataCtx.Close()

	storage.MustInit(dataCtx.Eth, dataCtx.DB)

	mws := httpMiddlewares(dataCtx)

	api.MustServeFromViper(storage.Register, mws...)
}
