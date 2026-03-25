package entities

import (
	domainErrors "Go_c/internal/domain/errors"
	"time"

	"github.com/google/uuid"
)

// Driver represents a driver in the system
type Driver struct {
	Base

	UserID        uuid.UUID  `json:"user_id" db:"user_id"`               // Linked user account
	LicenseNumber string     `json:"license_number" db:"license_number"` // Driver license
	Rating        float64    `json:"rating" db:"rating"`                 // Average rating
	TotalReviews  int        `json:"total_reviews" db:"total_reviews"`   // Number of reviews
	IsOnline      bool       `json:"is_online" db:"is_online"`           // Driver online status
	Version       int        `json:"version" db:"version"`               // Optimistic locking
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

//
// ==========================
// Behaviors
// ==========================
//

// Validate ensures driver fields are valid
func (d *Driver) Validate() error {
	if d.UserID == uuid.Nil {
		return domainErrors.ErrInvalidInput
	}
	if d.LicenseNumber == "" {
		return domainErrors.ErrLicenseInvalid
	}
	return nil
}

// SetOnline marks driver as available
func (d *Driver) SetOnline() {
	d.IsOnline = true
	d.UpdateTimestamp()
}

// SetOffline marks driver as offline
func (d *Driver) SetOffline() {
	d.IsOnline = false
	d.UpdateTimestamp()
}

// UpdateRating recalculates average rating given a new review
func (d *Driver) UpdateRating(newRating int) error {
	if newRating < 1 || newRating > 5 {
		return domainErrors.ErrRatingInvalid
	}

	total := float64(d.TotalReviews)*d.Rating + float64(newRating)
	d.TotalReviews++
	d.Rating = total / float64(d.TotalReviews)
	d.UpdateTimestamp()
	return nil
}

// SoftDelete marks the driver as deleted (does not remove from DB)
func (d *Driver) SoftDelete() {
	now := time.Now()
	d.DeletedAt = &now
	d.UpdateTimestamp()
}

// IsActive checks if driver is not deleted
func (d *Driver) IsActive() bool {
	return d.DeletedAt == nil
}
