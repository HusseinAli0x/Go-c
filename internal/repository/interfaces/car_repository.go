package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/domain/enums"
	"Go_c/internal/repository"
)

// CarRepository defines operations for Car entity
type CarRepository interface {

	// CRUD
	Create(ctx context.Context, car *entities.Car) error
	GetByID(ctx context.Context, id string) (*entities.Car, error)
	Update(ctx context.Context, car *entities.Car) error
	Delete(ctx context.Context, id string) error

	// Relations
	GetByDriver(ctx context.Context, driverID string) ([]*entities.Car, error)

	// Filtering & Listing
	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.Car, error)
	GetAvailableCars(ctx context.Context, opts repository.QueryOptions) ([]*entities.Car, error)
	GetByStatus(ctx context.Context, status enums.CarStatus, opts repository.QueryOptions) ([]*entities.Car, error)

	// Search (Full-Text)
	Search(ctx context.Context, query string, opts repository.QueryOptions) ([]*entities.Car, error)

	// Status Management
	UpdateStatus(ctx context.Context, carID string, status enums.CarStatus) error
}
