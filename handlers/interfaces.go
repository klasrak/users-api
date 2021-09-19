package handlers

import (
	"context"

	model "github.com/klasrak/users-api/models"
)

// UserService represents the user service implementation
type UserService interface {
	GetAll(ctx context.Context, name string) ([]model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, u *model.User) (*model.User, error)
	Update(ctx context.Context, id string, u *model.User) (*model.User, error)
}
