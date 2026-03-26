package interfaces

import (
	"context"

	"Go_c/internal/domain/entities"
	"Go_c/internal/repository"
)

// UserRepository defines operations for User entity
type UserRepository interface {

	// CRUD

	Create(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id string) error

	// Unique Queries

	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetByPhone(ctx context.Context, phone string) (*entities.User, error)

	// Search & Filtering

	List(ctx context.Context, opts repository.QueryOptions) ([]*entities.User, error)
	Search(ctx context.Context, query string, opts repository.QueryOptions) ([]*entities.User, error)

	// Existence Checks

	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
}

//TODO: Hussein ALi Mahood
