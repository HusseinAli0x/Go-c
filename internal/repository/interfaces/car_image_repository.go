package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/repository"
)

// CarImageRepository defines operations for CarImage entity
type CarImageRepository interface {

	// CRUD
	Create(ctx context.Context, img *entities.CarImage) error
	GetByID(ctx context.Context, id string) (*entities.CarImage, error)
	Update(ctx context.Context, img *entities.CarImage) error
	Delete(ctx context.Context, id string) error

	// Get all images for a specific car
	GetByCarID(ctx context.Context, carID string, opts repository.QueryOptions) ([]*entities.CarImage, error)

	// Special operation: set an image as primary for the car
	SetPrimary(ctx context.Context, carID string, imageID string) error
}
