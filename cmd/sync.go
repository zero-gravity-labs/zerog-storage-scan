package cmd

import (
	"context"
	"sync"

	nhSync "github.com/0glabs/0g-storage-scan/sync"
	viperUtil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	viperUtil.MustUnmarshalKey("sync", &conf)

	cs := nhSync.MustNewCatchupSyncer(dataCtx.Eth, dataCtx.DB, conf, dataCtx.EthCfg.AlertChannel, dataCtx.EthCfg.HealthReport)
	ss := nhSync.MustNewStorageSyncer(dataCtx.L2Sdks, dataCtx.DB, dataCtx.L2SdkCfg.AlertChannel, dataCtx.L2SdkCfg.HealthReport)
	syncer := nhSync.MustNewSyncer(dataCtx.Eth, dataCtx.DB, conf, cs, ss)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go syncer.Sync(ctx, &wg)

	GracefulShutdown(&wg, cancel)
}
