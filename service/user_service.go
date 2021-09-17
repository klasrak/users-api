package service

import (
	"context"

	"github.com/google/uuid"
	model "github.com/klasrak/users-api/models"
)

// UserService is a struct to inject a implementation of UserRepository
type UserService struct {
	UserRepository UserRepository
}

// GetAll calls repository GetAll and returns
func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.UserRepository.GetAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return s.UserRepository.GetByID(ctx, id)
}
