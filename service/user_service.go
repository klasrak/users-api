package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	model "github.com/klasrak/users-api/models"
	"github.com/klasrak/users-api/rerrors"
	"github.com/klasrak/users-api/utils"
)

// UserService is a struct to inject a implementation of UserRepository
type UserService struct {
	UserRepository UserRepository
}

// GetAll calls repository GetAll and returns
func (s *UserService) GetAll(ctx context.Context, name string) ([]model.User, error) {
	return s.UserRepository.GetAll(ctx, name)
}

// GetByID call repository GetById and returns
func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	uid, err := uuid.Parse(id)

	if err != nil {
		return nil, rerrors.NewBadRequest("invalid id")
	}

	return s.UserRepository.GetByID(ctx, uid)
}

// Create call repository Create and returns
func (s *UserService) Create(ctx context.Context, u *model.User) (*model.User, error) {

	if utils.IsUnderage(u.BirthDate, time.Now()) {
		return nil, rerrors.NewBadRequest("underage")
	}

	if !utils.IsBrazilianCPFValid(u.Cpf) {
		return nil, rerrors.NewBadRequest("cpf invalid")
	}

	return s.UserRepository.Create(ctx, u)
}
