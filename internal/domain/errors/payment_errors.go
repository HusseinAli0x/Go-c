package errors

import "errors"

var (
    ErrPaymentFailed        = errors.New("payment failed")
    ErrPaymentNotFound      = errors.New("payment not found")
    ErrPaymentAlreadyExists = errors.New("payment already exists for this booking")
)