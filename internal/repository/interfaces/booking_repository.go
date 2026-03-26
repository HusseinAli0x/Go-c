package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/domain/enums"
	"Go_c/internal/repository"
)

// BookingRepository defines operations for Booking entity
type BookingRepository interface {

	// CRUD
	Create(ctx context.Context, booking *entities.Booking) error
	GetByID(ctx context.Context, id string) (*entities.Booking, error)
	Update(ctx context.Context, booking *entities.Booking) error
	Delete(ctx context.Context, id string) error

	// Relations
	GetByUser(ctx context.Context, userID string, opts repository.QueryOptions) ([]*entities.Booking, error)
	GetByDriver(ctx context.Context, driverID string, opts repository.QueryOptions) ([]*entities.Booking, error)
	GetByCar(ctx context.Context, carID string, opts repository.QueryOptions) ([]*entities.Booking, error)

	// Active Bookings
	GetActiveByUser(ctx context.Context, userID string) ([]*entities.Booking, error)
	GetActiveByDriver(ctx context.Context, driverID string) ([]*entities.Booking, error)

	// Filtering & Listing
	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.Booking, error)
	GetByStatus(ctx context.Context, status enums.BookingStatus, opts repository.QueryOptions) ([]*entities.Booking, error)

	// Search (Full-Text)
	Search(ctx context.Context, query string, opts repository.QueryOptions) ([]*entities.Booking, error)

	// Status Updates (Optimized)
	UpdateStatus(ctx context.Context, bookingID string, status enums.BookingStatus) error

	// Existence & Validation
	ExistsActiveByUser(ctx context.Context, userID string) (bool, error)
	ExistsActiveByDriver(ctx context.Context, driverID string) (bool, error)
}
