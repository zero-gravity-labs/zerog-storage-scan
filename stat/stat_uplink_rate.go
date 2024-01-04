package stat

import (
	"context"
	"fmt"
	viperutil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type UplinkRateStat struct {
	Db       *store.MysqlStore
	interval *time.Duration
}

func MustNewUplinkRateStat(db *store.MysqlStore) *UplinkRateStat {
	var stat struct {
		StatIntervalDataUplinkRate time.Duration `default:"1s"`
	}
	viperutil.MustUnmarshalKey("stat", &stat)

	return &UplinkRateStat{
		Db:       db,
		interval: &stat.StatIntervalDataUplinkRate,
	}
}

func (urs *UplinkRateStat) DoStat(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	logrus.Info("Stat starting")

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		endTime := time.Now()
		startTime := endTime.Add(-*urs.interval)
		intervalInSec := endTime.Sub(startTime).Seconds()

		_, dataSize, err := urs.Db.SubmitStore.Count(&startTime, &endTime)
		if err != nil {
			logrus.WithError(err).Error("Stat data size error.")
			time.Sleep(10 * time.Second)
			continue
		}

		tps := float64(dataSize) / intervalInSec

		if err := urs.Db.ConfigStore.Upsert(store.CfgDataUplinkRate, fmt.Sprintf("%f", tps)); err != nil {
			logrus.WithError(err).Error("Update data uplink rate error.")
			time.Sleep(10 * time.Second)
			continue
		}

		time.Sleep(time.Second)
	}
}
