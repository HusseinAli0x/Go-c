package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/repository"
)

// NotificationRepository defines operations for Notification entity
type NotificationRepository interface {

	// CRUD
	Create(ctx context.Context, notification *entities.Notification) error
	GetByID(ctx context.Context, id string) (*entities.Notification, error)
	Update(ctx context.Context, notification *entities.Notification) error
	Delete(ctx context.Context, id string) error

	// Relations
	GetByUser(ctx context.Context, userID string, opts repository.QueryOptions) ([]*entities.Notification, error)

	// Filtering & Listing
	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.Notification, error)

	// Status Management
	MarkRead(ctx context.Context, notificationID string) error
	MarkUnread(ctx context.Context, notificationID string) error
	MarkAllRead(ctx context.Context, userID string) error

	// Counters
	CountUnread(ctx context.Context, userID string) (int, error)
}
