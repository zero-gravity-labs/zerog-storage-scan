package api

import (
	commonApi "github.com/Conflux-Chain/go-conflux-util/api"
)

var (
	ErrConfigNotFound  = commonApi.NewBusinessError(1001, "Config not found", nil)
	ErrAddressNotFound = commonApi.NewBusinessError(2001, "Account not found", nil)
)
