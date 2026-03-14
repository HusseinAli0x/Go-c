package errors

import "errors"

var (
	ErrDriverNotFound = errors.New("driver not found")
	ErrDriverOffline  = errors.New("driver is offline")
	ErrDriverBusy     = errors.New("driver is currently busy")
	ErrLicenseInvalid = errors.New("driver license is invalid")
)
