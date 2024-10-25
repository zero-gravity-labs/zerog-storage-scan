package stat

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrMaxPosFinalizedNotSync = errors.New("MaxPosFinalized not synced")
	ErrMinPosNotFinalized     = errors.New("MinPos not finalized")
)

type Topn[T any] interface {
	nextStatRange() (*T, error)
	calculateStat(T) error
}

type AbsTopn[T any] struct {
	Topn[T]
}

func (as *AbsTopn[T]) DoStat(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	logrus.Info("Stat starting")

	for {
		r, err := as.nextStatRange()

		if as.interrupted(ctx) {
			break
		}

		if err != nil {
			if !errors.Is(err, ErrMaxPosFinalizedNotSync) &&
				!errors.Is(err, ErrMinPosNotFinalized) {
				logrus.WithError(err).Error("Acquire next stat ranges for stats")
			}
			time.Sleep(time.Second)
			continue
		}

		err = as.calculateStat(*r)
		if err != nil {
			logrus.WithError(err).Error("Do stat")
			time.Sleep(time.Second * 10)
			continue
		}
	}
}

func (as *AbsTopn[T]) interrupted(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}
