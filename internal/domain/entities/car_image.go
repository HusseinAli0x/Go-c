package entities

import (
	domainErrors "Go_c/internal/domain/errors"

	"github.com/google/uuid"
)

// CarImage represents an image associated with a car
type CarImage struct {
	Base

	CarID     uuid.UUID `json:"car_id" db:"car_id"`         // Related car
	ImagePath string    `json:"image_path" db:"image_path"` // Path or URL of image
	IsPrimary bool      `json:"is_primary" db:"is_primary"` // Indicates main image
}

//
// ==========================
// Behaviors
// ==========================
//

// Validate ensures the image data is valid
func (ci *CarImage) Validate() error {
	if ci.CarID == uuid.Nil {
		return domainErrors.ErrInvalidInput
	}
	if ci.ImagePath == "" {
		return domainErrors.ErrInvalidInput
	}
	return nil
}

// SetPrimary marks this image as the main image
func (ci *CarImage) SetPrimary() {
	ci.IsPrimary = true
	ci.UpdateTimestamp()
}

// RemovePrimary removes primary flag
func (ci *CarImage) RemovePrimary() {
	ci.IsPrimary = false
	ci.UpdateTimestamp()
}

// UpdatePath updates image path
func (ci *CarImage) UpdatePath(path string) error {
	if path == "" {
		return domainErrors.ErrInvalidInput
	}

	ci.ImagePath = path
	ci.UpdateTimestamp()
	return nil
}
