package cmd

import (
	"context"
	viperutil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	nhSync "github.com/zero-gravity-labs/zerog-storage-scan/sync"
	"sync"
)

var (
	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Start sync, including block/submitLog",
		Run:   startSyncService,
	}
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

func startSyncService(*cobra.Command, []string) {
	logrus.Info("Start to sync evm space blockchain data into database")
	dataCtx := MustInitDataContext()
	defer dataCtx.Close()

	var conf nhSync.SyncConfig
	viperutil.MustUnmarshalKey("sync", &conf)

	cs := nhSync.MustNewCatchupSyncer(dataCtx.Eth, dataCtx.DB, conf)
	ss := nhSync.MustNewStorageSyncer(dataCtx.L2Sdk, dataCtx.DB)
	syncer := nhSync.MustNewSyncer(dataCtx.Eth, dataCtx.DB, conf, cs, ss)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go syncer.Sync(ctx, &wg)

	GracefulShutdown(&wg, cancel)
}
