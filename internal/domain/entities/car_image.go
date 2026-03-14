package entities

import (
	"github.com/google/uuid"
)

// CarImage represents an image associated with a car.
// Inherits ID, CreatedAt, UpdatedAt from Base.
// A car can have multiple images, and one of them can be marked as the primary image.
type CarImage struct {
	Base                // ID, CreatedAt, UpdatedAt
	CarID     uuid.UUID `json:"car_id" db:"car_id"`         // Reference to the related car
	ImagePath string    `json:"image_path" db:"image_path"` // Path or URL of the stored image
	IsPrimary bool      `json:"is_primary" db:"is_primary"` // Indicates if this image is the main car image
}

//
// ==========================
// CarImage Behaviors
// ==========================
//

// SetPrimary marks this image as the primary image for the car.
func (c *CarImage) SetPrimary() {
	c.IsPrimary = true
	c.UpdateTimestamp()
}

// RemovePrimary removes the primary flag from the image.
func (c *CarImage) RemovePrimary() {
	c.IsPrimary = false
	c.UpdateTimestamp()
}

// IsMainImage checks whether this image is the primary image of the car.
func (c *CarImage) IsMainImage() bool {
	return c.IsPrimary
}
