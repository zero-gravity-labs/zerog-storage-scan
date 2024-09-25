package api

import (
	"github.com/pkg/errors"
)

// Unexpected system error, e.g. database error, blockchain rpc error, io error. Http status is 600

func ErrBlockchainRPC(err error) error {
	return errors.WithMessage(err, "Blockchain RPC exception")
}

func ErrStorageNodeRPC(err error) error {
	return errors.WithMessage(err, "Storage node RPC exception")
}

func ErrDatabase(err error) error {
	return errors.WithMessage(err, "Database exception")
}

func ErrBatchGetAddress(err error) error {
	return ErrDatabase(errors.WithMessage(err, "Failed to batch get addresses' info"))
}
