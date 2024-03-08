package cmd

import (
	nhApi "github.com/0glabs/0g-storage-scan/api"
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

	nhApi.MustInit(dataCtx.Eth, dataCtx.DB)

	api.MustServeFromViper(nhApi.RegisterRouter)
}
