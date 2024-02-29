package api

import (
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
)

var (
	ErrConfigNotFound        = commonApi.NewBusinessError(1001, "Config not found", nil)
	ErrAddressNotFound       = commonApi.NewBusinessError(1002, "Account not found", nil)
	ErrStatTypeNotSupported  = commonApi.NewBusinessError(1003, "Stat type not supported", nil)
	ErrStorageBaseFeeNotStat = commonApi.NewBusinessError(1004, "Storage base fee not stat", nil)
)
