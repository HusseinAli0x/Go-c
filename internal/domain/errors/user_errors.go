package errors

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrEmailExists     = errors.New("email already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInactiveUser    = errors.New("user is inactive or deleted")
)
