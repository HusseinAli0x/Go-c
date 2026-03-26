package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/domain/enums"
	"Go_c/internal/repository"
)

// PaymentRepository defines operations for Payment entity
type PaymentRepository interface {

	// CRUD
	Create(ctx context.Context, payment *entities.Payment) error
	GetByID(ctx context.Context, id string) (*entities.Payment, error)
	Update(ctx context.Context, payment *entities.Payment) error
	Delete(ctx context.Context, id string) error

	// Relations
	GetByBooking(ctx context.Context, bookingID string) ([]*entities.Payment, error)

	// Filtering & Listing
	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.Payment, error)
	GetByStatus(ctx context.Context, status enums.PaymentStatus, opts repository.QueryOptions) ([]*entities.Payment, error)

	// Status Management
	MarkPaid(ctx context.Context, paymentID string) error
	MarkFailed(ctx context.Context, paymentID string) error

	// Validation
	ExistsByBooking(ctx context.Context, bookingID string) (bool, error)
}
