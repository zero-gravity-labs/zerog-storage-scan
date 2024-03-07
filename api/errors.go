package api

import (
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
)

var (
	ErrStatTypeNotSupported  = commonApi.NewBusinessError(1003, "Stat type not supported", nil)
	ErrStorageBaseFeeNotStat = commonApi.NewBusinessError(1004, "Storage base fee not stat", nil)
)
