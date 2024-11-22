package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/0glabs/0g-storage-client/node"
	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/Conflux-Chain/go-conflux-util/parallel"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Status uint8

const (
	NotUploaded Status = iota
	PartialUploaded
	Uploaded
	Pruned
)

var (
	BatchGetSubmitsByGoroutines = 10
)

type FileInfoParam struct {
	SubmissionIndex uint64
	Status          uint8
}

type FileInfoResult struct {
	Data    *FileInfo
	Err     error
	Latency time.Duration
}

type FileInfoExecutor struct {
	storageConfig StorageConfig
	rpcParams     []FileInfoParam
	rpcResults    map[uint64]*FileInfoResult
}

// ParallelDo implements the parallel.Interface
func (executor *FileInfoExecutor) ParallelDo(ctx context.Context, routine, task int) (*FileInfoResult, error) {
	rpcParam := executor.rpcParams[task]
	var result FileInfoResult
	result.Data, result.Err = executor.getFileInfo(ctx, executor.storageConfig, rpcParam, task)

	return &result, nil
}

// ParallelCollect implements the parallel.Interface
func (executor *FileInfoExecutor) ParallelCollect(ctx context.Context, result *parallel.Result[*FileInfoResult]) error {
	rpcParam := executor.rpcParams[result.Task]
	executor.rpcResults[rpcParam.SubmissionIndex] = result.Value

	return nil
}

type FileInfo struct {
	FileInfoParam
	UploadedSegNum uint64
}

type StorageConfig struct {
	Indexer         string
	Retry           int
	RetryInterval   time.Duration `default:"1s"`
	RequestTimeout  time.Duration `default:"3s"`
	MaxConnsPerHost int           `default:"1024"`
	AlertChannel    string
	HealthReport    health.TimedCounterConfig
}

// getFileInfo implements the rpcFunc interface
func (executor *FileInfoExecutor) getFileInfo(ctx context.Context, storageConfig StorageConfig,
	rpcParam FileInfoParam, task int) (*FileInfo, error) {
	fileInfo := FileInfo{rpcParam, 0}
	updated := false

	info, err := GetFileInfoByTxSeq(storageConfig, rpcParam.SubmissionIndex)
	if err == nil && info != nil {
		var status uint8
		if info.Pruned {
			status = uint8(Pruned)
		} else if info.Finalized {
			status = uint8(Uploaded)
		} else if info.UploadedSegNum > 0 {
			status = uint8(PartialUploaded)
		}

		if status > fileInfo.Status {
			fileInfo.Status = status
			fileInfo.UploadedSegNum = info.UploadedSegNum
			updated = true
		}
	}

	if !updated {
		return nil, errors.Errorf("Submit %v with status %v not updated", rpcParam.SubmissionIndex, rpcParam.Status)
	}

	return &fileInfo, nil
}

func BatchGetFileInfos(ctx context.Context, storageConfig StorageConfig, rpcParams []FileInfoParam) (
	map[uint64]*FileInfoResult, error) {
	executor := FileInfoExecutor{
		storageConfig: storageConfig,
		rpcParams:     rpcParams,
		rpcResults:    make(map[uint64]*FileInfoResult),
	}

	start := time.Now()
	opt := parallel.SerialOption{Routines: BatchGetSubmitsByGoroutines}
	if err := parallel.Serial(ctx, &executor, len(rpcParams), opt); err != nil {
		return nil, err
	}
	elapsed := time.Since(start)

	logrus.WithFields(logrus.Fields{
		"files":       len(rpcParams),
		"elapsed(ms)": elapsed,
		"average(ms)": elapsed.Milliseconds() / int64(len(rpcParams)),
	}).Debug("Batch get file info")

	return executor.rpcResults, nil
}

func GetFileInfoByTxSeq(storageConfig StorageConfig, seqNo uint64) (*node.FileInfo, error) {
	url := fmt.Sprintf("%s/file/info/%v", storageConfig.Indexer, seqNo)

	var result node.FileInfo
	client := resty.New()
	resp, err := client.R().SetResult(&result).Get(url)
	if err != nil || resp.IsError() {
		return nil, errors.WithMessagef(err, "Failed to get file info, seqNo %v %s", seqNo, resp.String())
	}

	return &result, nil
}

func GetNodeStatus(storageConfig StorageConfig) (*node.Status, error) {
	url := fmt.Sprintf("%s/node/status", storageConfig.Indexer)

	var result node.Status
	client := resty.New()
	resp, err := client.R().SetResult(&result).Get(url)
	if err != nil || resp.IsError() {
		return nil, errors.WithMessagef(err, "Failed to get node status %s", resp.String())
	}

	return &result, nil
}
