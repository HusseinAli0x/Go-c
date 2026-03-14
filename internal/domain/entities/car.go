package entities

import (
	"time"

	"Go_c/internal/domain/enums"

	"github.com/google/uuid"
)

// Car represents a vehicle owned by a driver in the system.
// Inherits ID, CreatedAt, UpdatedAt from Base.
type Car struct {
	Base                        // ID, CreatedAt, UpdatedAt
	DriverID    uuid.UUID       `json:"driver_id" db:"driver_id"`             // Driver who owns this car
	Brand       string          `json:"brand" db:"brand"`                     // Car manufacturer
	Model       string          `json:"model" db:"model"`                     // Car model
	Year        int             `json:"year" db:"year"`                       // Manufacturing year
	PlateNumber string          `json:"plate_number" db:"plate_number"`       // Unique plate number
	Color       string          `json:"color" db:"color"`                     // Car color
	Status      enums.CarStatus `json:"status" db:"status"`                   // Current operational status
	Version     int             `json:"version" db:"version"`                 // Optimistic locking version
	DeletedAt   *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"` // Soft delete timestamp
}

//
// ==========================
// Car Behaviors
// ==========================
//

// SetAvailable marks the car as available for accepting rides.
func (c *Car) SetAvailable() {
	c.Status = enums.CarAvailable
	c.UpdateTimestamp()
}

// SetBusy marks the car as currently being used for a ride.
func (c *Car) SetBusy() {
	c.Status = enums.CarBusy
	c.UpdateTimestamp()
}

// SetOffline marks the car as offline (driver not active).
func (c *Car) SetOffline() {
	c.Status = enums.CarOffline
	c.UpdateTimestamp()
}

// SetMaintenance marks the car as unavailable due to maintenance.
func (c *Car) SetMaintenance() {
	c.Status = enums.CarMaintenance
	c.UpdateTimestamp()
}

// UpdateInfo updates the car basic information.
func (c *Car) UpdateInfo(brand, model, color string, year int) {
	c.Brand = brand
	c.Model = model
	c.Color = color
	c.Year = year
	c.UpdateTimestamp()
}

// IsActive checks whether the car has not been soft deleted.
func (c *Car) IsActive() bool {
	return c.DeletedAt == nil
}

// MarkDeleted performs a soft delete by setting the DeletedAt timestamp.
func (c *Car) MarkDeleted() {
	now := time.Now()
	c.DeletedAt = &now
}
