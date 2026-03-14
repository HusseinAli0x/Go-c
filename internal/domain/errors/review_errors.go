package errors

import "errors"

var (
	ErrRatingInvalid       = errors.New("rating must be between 1 and 5")
	ErrReviewAlreadyExists = errors.New("review already exists for this booking")
	ErrReviewNotAllowed    = errors.New("review is not allowed for this booking")
)
