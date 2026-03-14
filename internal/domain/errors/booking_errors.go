package errors

import "errors"

var (
    ErrBookingNotFound      = errors.New("booking not found")
    ErrBookingConflict      = errors.New("booking conflict")
    ErrBookingAlreadyPaid   = errors.New("booking already paid")
    ErrBookingCannotCancel  = errors.New("booking cannot be cancelled")
)