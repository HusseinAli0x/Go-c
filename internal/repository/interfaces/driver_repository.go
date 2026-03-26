package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/repository"
)

// DriverRepository defines operations for Driver entity
type DriverRepository interface {

	// CRUD
	Create(ctx context.Context, driver *entities.Driver)
	GetByID(ctx context.Context, id string) (*entities.Driver, error)
	Update(ctx context.Context, driver *entities.Driver) error
	Delete(ctx context.Context, id string) error

	// Relations
	GetByUserID(ctx context.Context, userID string) (*entities.Driver, error)

	// Availability
	GetOnlineDrivers(ctx context.Context, opts repository.QueryOptions) ([]*entities.Driver, error)
	GetAvailableDrivers(ctx context.Context, opts repository.QueryOptions) ([]*entities.Driver, error)

	// Filtering & Listing
	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.Driver, error)

	// Status Management
	SetOnline(ctx context.Context, driverID string) error
	SetOffline(ctx context.Context, driverID string) error

	// Rating & Stats
	UpdateRating(ctx context.Context, driverID string, rating float64, totalReviews int) error

	// Validation
	ExistsByUserID(ctx context.Context, userID string) (bool, error)
}
