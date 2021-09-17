package service

import (
	"context"

	model "github.com/klasrak/users-api/models"
)

// UserService is a struct to inject a implementation of UserRepository
type UserService struct {
	UserRepository UserRepository
}

// GetAll calls repository GetAll and returns
func (s *UserService) GetAll(ctx context.Context) ([]*model.User, error) {
	return s.UserRepository.GetAll(ctx)
}
