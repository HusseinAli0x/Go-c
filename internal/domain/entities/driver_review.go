package entities

import (
	domainErrors "Go_c/internal/domain/errors"

	"github.com/google/uuid"
)

// DriverReview represents a user's review for a driver
type DriverReview struct {
	Base

	BookingID uuid.UUID `json:"booking_id" db:"booking_id"` // Related booking
	DriverID  uuid.UUID `json:"driver_id" db:"driver_id"`   // Driver being reviewed
	UserID    uuid.UUID `json:"user_id" db:"user_id"`       // User who wrote the review
	Rating    int       `json:"rating" db:"rating"`         // Rating 1..5
	Comment   string    `json:"comment" db:"comment"`       // Optional comment
}

//
// ==========================
// Behaviors
// ==========================
//

// Validate ensures review data is correct
func (r *DriverReview) Validate() error {
	if r.BookingID == uuid.Nil || r.DriverID == uuid.Nil || r.UserID == uuid.Nil {
		return domainErrors.ErrInvalidInput
	}
	if r.Rating < 1 || r.Rating > 5 {
		return domainErrors.ErrRatingInvalid
	}
	return nil
}

// UpdateComment updates the review comment
func (r *DriverReview) UpdateComment(comment string) {
	r.Comment = comment
	r.UpdateTimestamp()
}

// UpdateRating updates the rating value (1..5)
func (r *DriverReview) UpdateRating(rating int) error {
	if rating < 1 || rating > 5 {
		return domainErrors.ErrRatingInvalid
	}
	r.Rating = rating
	r.UpdateTimestamp()
	return nil
}

// IsValid checks if review is valid (rating + IDs)
func (r *DriverReview) IsValid() bool {
	return r.Rating >= 1 && r.Rating <= 5 &&
		r.BookingID != uuid.Nil &&
		r.DriverID != uuid.Nil &&
		r.UserID != uuid.Nil
}
