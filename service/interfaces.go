package service

import (
	"context"

	"github.com/google/uuid"
	model "github.com/klasrak/users-api/models"
)

// UserRepository representes the user repository implementation
type UserRepository interface {
	GetAll(ctx context.Context, name string) ([]model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	Create(ctx context.Context, u *model.User) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id string) error
}
