package entities

import (
	"github.com/google/uuid"
)

// Driver represents a driver in the system.
// Each driver is linked to a user account and contains
// driver-specific information such as license, rating,
// and availability status.
type Driver struct {
	Base                    // ID, CreatedAt, UpdatedAt
	UserID        uuid.UUID `json:"user_id" db:"user_id"`               // Reference to the associated user account
	LicenseNumber string    `json:"license_number" db:"license_number"` // Driver license number
	Rating        float64   `json:"rating" db:"rating"`                 // Average rating (e.g., 4.8)
	TotalReviews  int       `json:"total_reviews" db:"total_reviews"`   // Total number of reviews received
	IsOnline      bool      `json:"is_online" db:"is_online"`           // Indicates whether the driver is currently online
	Version       int       `json:"version" db:"version"`               // Optimistic locking version
}

// ==========================
// Driver Behaviors
// ==========================

// GoOnline marks the driver as available for accepting rides.
func (d *Driver) GoOnline() {
	d.IsOnline = true
	d.UpdateTimestamp()
}

// GoOffline marks the driver as unavailable for rides.
func (d *Driver) GoOffline() {
	d.IsOnline = false
	d.UpdateTimestamp()
}

// UpdateRating recalculates the driver's average rating
// when a new review is submitted.
func (d *Driver) UpdateRating(newRating float64) {
	total := float64(d.Rating*float64(d.TotalReviews)) + newRating
	d.TotalReviews++
	d.Rating = total / float64(d.TotalReviews)
	d.UpdateTimestamp()
}
