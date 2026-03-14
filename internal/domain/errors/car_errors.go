package errors

import "errors"

var (
    ErrCarNotFound      = errors.New("car not found")
    ErrCarUnavailable   = errors.New("car is unavailable")
    ErrCarMaintenance   = errors.New("car is under maintenance")
)