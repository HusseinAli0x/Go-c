package entities

import (
	"Go_c/internal/domain/enums"
	domainErrors "Go_c/internal/domain/errors"

	"github.com/google/uuid"
)

// Car represents a vehicle assigned to a driver
type Car struct {
	Base

	DriverID    uuid.UUID       `json:"driver_id" db:"driver_id"`                 // Owner driver
	Brand       string          `json:"brand" db:"brand"`                         // Car brand
	Model       string          `json:"model" db:"model"`                         // Car model
	Year        *int            `json:"year,omitempty" db:"year"`                 // Manufacturing year
	PlateNumber *string         `json:"plate_number,omitempty" db:"plate_number"` // Unique plate number
	Color       *string         `json:"color,omitempty" db:"color"`               // Car color
	Status      enums.CarStatus `json:"status" db:"status"`                       // Car status
}

//
// ==========================
// Behaviors
// ==========================
//

// Validate checks required fields
func (c *Car) Validate() error {
	if c.Brand == "" || c.Model == "" {
		return domainErrors.ErrInvalidInput
	}
	return nil
}

// IsAvailable checks if car is ready for booking
func (c *Car) IsAvailable() bool {
	return c.Status == enums.CarAvailable
}

// MarkBusy sets car status to busy
func (c *Car) MarkBusy() error {
	if c.Status != enums.CarAvailable {
		return domainErrors.ErrCarUnavailable
	}

	c.Status = enums.CarBusy
	c.UpdateTimestamp()
	return nil
}

// MarkAvailable sets car back to available
func (c *Car) MarkAvailable() {
	c.Status = enums.CarAvailable
	c.UpdateTimestamp()
}

// SetMaintenance sets car to maintenance mode
func (c *Car) SetMaintenance() {
	c.Status = enums.CarMaintenance
	c.UpdateTimestamp()
}

// SetOffline sets car to offline
func (c *Car) SetOffline() {
	c.Status = enums.CarOffline
	c.UpdateTimestamp()
}

// AssignDriver assigns a driver to the car
func (c *Car) AssignDriver(driverID uuid.UUID) {
	c.DriverID = driverID
	c.UpdateTimestamp()
}
