package service

import (
	"context"

	"github.com/google/uuid"
	model "github.com/klasrak/users-api/models"
	"github.com/klasrak/users-api/rerrors"
)

// UserService is a struct to inject a implementation of UserRepository
type UserService struct {
	UserRepository UserRepository
}

// GetAll calls repository GetAll and returns
func (s *UserService) GetAll(ctx context.Context, name string) ([]model.User, error) {
	return s.UserRepository.GetAll(ctx, name)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	uid, err := uuid.Parse(id)

	if err != nil {
		return nil, rerrors.NewBadRequest("invalid id")
	}

	return s.UserRepository.GetByID(ctx, uid)
}

func (s *UserService) Create(ctx context.Context, u *model.User) (*model.User, error) {
	return s.UserRepository.Create(ctx, u)
}
