package cmd

import (
	"context"
	"sync"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zero-gravity-labs/zerog-storage-scan/stat"
)

var (
	statCmd = &cobra.Command{
		Use:   "stat",
		Short: "Start stat, including transactions and data size of storage",
		Run:   startStatService,
	}
)

func init() {
	rootCmd.AddCommand(statCmd)
}

func startStatService(*cobra.Command, []string) {
	logrus.Info("Start to stat transactions and data size of storage")
	cfg := stat.StatConfig{}
	viper.MustUnmarshalKey("stat", &cfg)

	dataCtx := MustInitDataContext()
	defer dataCtx.Close()

	startTime := stat.MustDefaultRangeStart(dataCtx.Eth)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	stSubmit := stat.MustNewStatSubmit(&cfg, dataCtx.DB, dataCtx.Eth, startTime)
	go stSubmit.DoStat(ctx, &wg)

	GracefulShutdown(&wg, cancel)
}
