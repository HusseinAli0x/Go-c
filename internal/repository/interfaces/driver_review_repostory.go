package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/repository"
)

// DriverReviewRepository defines operations for DriverReview entity
type DriverReviewRepository interface {

	// CRUD
	Create(ctx context.Context, review *entities.DriverReview) error
	GetByID(ctx context.Context, id string) (*entities.DriverReview, error)
	Update(ctx context.Context, review *entities.DriverReview) error
	Delete(ctx context.Context, id string) error

	// Relations
	GetByDriver(ctx context.Context, driverID string, opts repository.QueryOptions) ([]*entities.DriverReview, error)
	GetByUser(ctx context.Context, userID string, opts repository.QueryOptions) ([]*entities.DriverReview, error)
	GetByBooking(ctx context.Context, bookingID string) (*entities.DriverReview, error)

	// Aggregates
	AverageRating(ctx context.Context, driverID string) (float64, error)
	CountReviews(ctx context.Context, driverID string) (int, error)

	// Filtering & Pagination
	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.DriverReview, error)
}
