package handlers

import (
	"context"

	model "github.com/klasrak/users-api/models"
)

// UserService represents the user service implementation
type UserService interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
}
