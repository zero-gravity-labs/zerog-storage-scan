package stat

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	viperutil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zero-gravity-labs/zerog-storage-client/node"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	"gorm.io/gorm"
)

type LogSyncInfoStat struct {
	db       *store.MysqlStore
	l2Sdk    *node.Client
	interval time.Duration
}

func MustNewSyncStatusStat(db *store.MysqlStore, l2Sdk *node.Client) *LogSyncInfoStat {
	var stat struct {
		StatIntervalSyncStatus time.Duration `default:"1s"`
	}
	viperutil.MustUnmarshalKey("stat", &stat)

	return &LogSyncInfoStat{
		db:       db,
		l2Sdk:    l2Sdk,
		interval: stat.StatIntervalSyncStatus,
	}
}

func (s *LogSyncInfoStat) DoStat(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	logrus.Info("Stat log sync info starting")

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		var submit store.Submit
		err := s.db.Store.DB.Last(&submit).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithError(err).Info("No submit record to update log sync info")
			time.Sleep(10 * time.Second)
			continue
		}
		if err != nil {
			logrus.WithError(err).Error("Failed to get submit record to update log sync info")
			time.Sleep(10 * time.Second)
			continue
		}

		zgStatus, err := s.l2Sdk.ZeroGStorage().GetStatus()
		if err != nil {
			logrus.WithError(err).Error("Failed to get zg status to update log sync info")
			time.Sleep(10 * time.Second)
			continue
		}

		status := LogSyncInfo{
			LogSyncHeight:   submit.BlockNumber,
			L2LogSyncHeight: zgStatus.LogSyncHeight,
		}
		statusBytes, err := json.Marshal(status)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal log sync info")
			continue
		}

		if err := s.db.ConfigStore.Upsert(store.KeyLogSyncInfo, string(statusBytes)); err != nil {
			logrus.WithError(err).Error("Update log sync info error.")
			time.Sleep(10 * time.Second)
			continue
		}

		time.Sleep(s.interval)
	}
}
