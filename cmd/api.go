package cmd

import (
	"context"
	"sync"

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

	storage.MustInit(dataCtx.DB, dataCtx.Eth, dataCtx.StorageConfig)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	go storage.ScheduleCache(ctx, &wg)

	mws := httpMiddlewares(dataCtx)
	api.MustServeFromViper(storage.Register, mws...)

	GracefulShutdown(&wg, cancel)
}
