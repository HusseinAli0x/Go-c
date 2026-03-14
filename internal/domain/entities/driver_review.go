package entities

import (
	"Go_c/internal/domain/errors"

	"github.com/google/uuid"
)

// DriverReview represents a review submitted by a user for a specific driver
// after completing a booking. It contains the rating, optional comment,
// and metadata about when the review was created.
type DriverReview struct {
	Base                // ID, CreatedAt, UpdatedAt
	BookingID uuid.UUID `json:"booking_id" db:"booking_id"`     // The booking associated with this review
	DriverID  uuid.UUID `json:"driver_id" db:"driver_id"`       // The driver being reviewed
	UserID    uuid.UUID `json:"user_id" db:"user_id"`           // The user who submitted the review
	Rating    int       `json:"rating" db:"rating"`             // Rating value (1 to 5)
	Comment   *string   `json:"comment,omitempty" db:"comment"` // Optional user comment
}

// ==========================
// DriverReview Behaviors
// ==========================

// ValidateRating checks whether the current rating value
// is within the allowed range (1..5).
func (r *DriverReview) ValidateRating() error {
	if r.Rating < 1 || r.Rating > 5 {
		return errors.ErrRatingInvalid
	}
	return nil
}

// SetRating assigns a rating to the review after validating
// that the provided value is between 1 and 5.
func (r *DriverReview) SetRating(rating int) error {
	if rating < 1 || rating > 5 {
		return errors.ErrRatingInvalid
	}

	r.Rating = rating
	r.UpdateTimestamp()
	return nil
}

// SetComment attaches a comment to the review.
func (r *DriverReview) SetComment(comment string) {
	r.Comment = &comment
	r.UpdateTimestamp()
}

// ClearComment removes the existing comment from the review.
func (r *DriverReview) ClearComment() {
	r.Comment = nil
	r.UpdateTimestamp()
}

// HasComment checks whether the review contains a comment.
func (r *DriverReview) HasComment() bool {
	return r.Comment != nil
}

// IsPositiveReview determines whether the review is considered positive.
// Typically ratings of 4 or 5 are treated as positive feedback.
func (r *DriverReview) IsPositiveReview() bool {
	return r.Rating >= 4
}

// IsNegativeReview determines whether the review is considered negative.
// Ratings of 1 or 2 usually indicate a negative experience.
func (r *DriverReview) IsNegativeReview() bool {
	return r.Rating <= 2
}
