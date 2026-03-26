package interfaces

import (
	"context"
)

// TransactionManager manages DB transactions across multiple repositories
type TransactionManager interface {
	Begin(ctx context.Context) (TxContext, error)
}

// TxContext represents a transactional scope
type TxContext interface {
	// Repositories inside the transaction
	User() UserRepository
	Driver() DriverRepository
	Car() CarRepository
	CarImage() CarImageRepository
	DriverLocation() DriverLocationRepository
	DriverReview() DriverReviewRepository
	Booking() BookingRepository
	Payment() PaymentRepository
	Notification() NotificationRepository

	// Commit or rollback
	Commit() error
	Rollback() error
}
