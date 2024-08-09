package api

import (
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
)

const (
	ErrCodeNoMatchingRecords = 1001
	ErrCodeBlockchain        = 1002
)

func ErrNoMatchingRecords(err error) *commonApi.BusinessError {
	return commonApi.NewBusinessError(ErrCodeNoMatchingRecords, "No matching records found", err.Error())
}

func ErrBlockchainRPC(err error) *commonApi.BusinessError {
	return commonApi.NewBusinessError(ErrCodeBlockchain, "Blockchain RPC failure", err.Error())
}
