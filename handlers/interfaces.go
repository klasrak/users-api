package handlers

import (
	"context"

	"github.com/google/uuid"
	model "github.com/klasrak/users-api/models"
)

// UserService represents the user service implementation
type UserService interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
}
