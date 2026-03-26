package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/repository"
)

// DriverLocationRepository defines operations for DriverLocation entity
type DriverLocationRepository interface {

	// CRUD
	Create(ctx context.Context, loc *entities.DriverLocation) error
	GetByDriverID(ctx context.Context, driverID string) (*entities.DriverLocation, error)
	Update(ctx context.Context, loc *entities.DriverLocation) error
	Delete(ctx context.Context, driverID string) error

	// Filtering & Pagination
	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.DriverLocation, error)

	// Geo Queries
	// GetDriversNear returns drivers within radius (meters) of given point
	GetDriversNear(ctx context.Context, point entities.Point, radiusMeters float64, opts repository.QueryOptions) ([]*entities.DriverLocation, error)
}
