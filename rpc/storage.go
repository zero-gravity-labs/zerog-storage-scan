package rpc

import (
	"context"
	"time"

	"github.com/0glabs/0g-storage-client/node"
	"github.com/Conflux-Chain/go-conflux-util/parallel"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Status uint8

const (
	NotUploaded Status = iota
	Uploading
	Uploaded
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
	l2sdks     []*node.Client
	rpcParams  []FileInfoParam
	rpcResults map[uint64]*FileInfoResult
}

// ParallelDo implements the parallel.Interface
func (executor *FileInfoExecutor) ParallelDo(ctx context.Context, routine, task int) (*FileInfoResult, error) {
	rpcParam := executor.rpcParams[task]
	var result FileInfoResult
	result.Data, result.Err = executor.getFileInfo(executor.l2sdks, rpcParam, task)

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

// getFileInfo implements the rpcFunc interface
func (executor *FileInfoExecutor) getFileInfo(l2Sdks []*node.Client, rpcParam FileInfoParam, task int) (*FileInfo, error) {
	fileInfo := FileInfo{rpcParam, 0}
	updated := false

	for _, l2Sdk := range l2Sdks {
		info, err := l2Sdk.ZeroGStorage().GetFileInfoByTxSeq(rpcParam.SubmissionIndex)
		if err == nil && info != nil {
			var status uint8
			if info.Finalized {
				status = uint8(Uploaded)
			} else if info.UploadedSegNum > 0 {
				status = uint8(Uploading)
			}

			if status > fileInfo.Status {
				fileInfo.Status = status
				fileInfo.UploadedSegNum = info.UploadedSegNum
				updated = true
			}

			if status == uint8(Uploaded) {
				break
			}
		}
	}

	if !updated {
		return nil, errors.Errorf("Submit %v with status %v not updated", rpcParam.SubmissionIndex, rpcParam.Status)
	}

	return &fileInfo, nil
}

func BatchGetFileInfos(ctx context.Context, l2sdks []*node.Client, rpcParams []FileInfoParam) (
	map[uint64]*FileInfoResult, error) {
	executor := FileInfoExecutor{
		l2sdks:     l2sdks,
		rpcParams:  rpcParams,
		rpcResults: make(map[uint64]*FileInfoResult),
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
