package metrics

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/metrics"
	ethereumMetrics "github.com/ethereum/go-ethereum/metrics"
)

type Metrics struct {
	API   APIMetrics
	Sync  SyncMetrics
	Store StoreMetrics
}

var Registry Metrics

type APIMetrics struct{}

func (*APIMetrics) UpdateDuration(path string, status, code int, start time.Time) {
	var isSuccess, isInternalError bool
	if isSuccess = status == http.StatusOK && code == api.ErrCodeSuccess; !isSuccess {
		isInternalError = status == api.ErrCodeInternal
	}

	// Overall rate statistics
	metrics.GetOrRegisterTimeWindowPercentageDefault(100, "openapi/rate/success").Mark(isSuccess)
	metrics.GetOrRegisterTimeWindowPercentageDefault(0, "openapi/rate/internalErr").Mark(isInternalError)
	metrics.GetOrRegisterTimeWindowPercentageDefault(0, "openapi/rate/nonInternalErr").Mark(!isSuccess && !isInternalError)

	// API rate statistics
	metrics.GetOrRegisterTimeWindowPercentageDefault(100, "openapi/rate/success/%v", path).Mark(isSuccess)
	metrics.GetOrRegisterTimeWindowPercentageDefault(0, "openapi/rate/internalErr/%v", path).Mark(isInternalError)
	metrics.GetOrRegisterTimeWindowPercentageDefault(0, "openapi/rate/nonInternalErr/%v", path).Mark(!isSuccess && !isInternalError)

	// Update QPS & Latency
	metrics.GetOrRegisterTimer("openapi/duration/all").UpdateSince(start)
	metrics.GetOrRegisterTimer("openapi/duration/%v", path).UpdateSince(start)
}

type SyncMetrics struct{}

func (*SyncMetrics) SyncOnceQps(err error) ethereumMetrics.Timer {
	if IsInterfaceValNil(err) {
		return metrics.GetOrRegisterTimer("sync/once/success")
	}

	return metrics.GetOrRegisterTimer("sync/once/failure")
}

func (*SyncMetrics) SyncOnceSize() ethereumMetrics.Histogram {
	return metrics.GetOrRegisterHistogram("sync/once/size")
}

func (*SyncMetrics) QueryEthData(rpcMethod string) ethereumMetrics.Timer {
	return metrics.GetOrRegisterTimer(fmt.Sprintf("sync/%v/fullnode", rpcMethod))
}

func (*SyncMetrics) QueryEthDataAvailability(rpcMethod string) metrics.Percentage {
	return metrics.GetOrRegisterTimeWindowPercentageDefault(0, "sync/%v/fullnode/availability", rpcMethod)
}

// Store metrics
type StoreMetrics struct{}

func (*StoreMetrics) Push() ethereumMetrics.Timer {
	return metrics.GetOrRegisterTimer("store/push")
}

func (*StoreMetrics) Pop() ethereumMetrics.Timer {
	return metrics.GetOrRegisterTimer("store/pop")
}

// Helper function to check if interface value is nil, since "i == nil" checks nil interface case only.
// Refer to https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1 for more details.
func IsInterfaceValNil(i interface{}) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func:
		return reflect.ValueOf(i).IsNil()
	}

	return false
}
