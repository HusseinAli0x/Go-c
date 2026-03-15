package entities

import (
	"time"

	"github.com/google/uuid"
)

// Base struct with common fields for all entities (without soft delete)
type Base struct {
	ID        uuid.UUID `json:"id" db:"id"`                 // Primary key
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Record creation timestamp
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // Last update timestamp
}

// ==========================
// Base Behaviors
// ==========================

// SetTimestamps sets CreatedAt and UpdatedAt to now (useful on insert)
func (b *Base) SetTimestamps() {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
}

// UpdateTimestamp updates UpdatedAt to now
func (b *Base) UpdateTimestamp() {
	b.UpdatedAt = time.Now()
}

// Point represents a geographical coordinate.
// Can be used by Booking, DriverLocation, etc.
type Point struct {
	Latitude  float64 `json:"latitude"`  // Latitude in decimal degrees
	Longitude float64 `json:"longitude"` // Longitude in decimal degrees
}
