package service

import (
	"context"
	"net/mail"

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

	if utils.IsUnderage(u.BirthDate) {
		return nil, rerrors.NewBadRequest("underage")
	}

	if !utils.IsBrazilianCPFValid(u.Cpf) {
		return nil, rerrors.NewBadRequest("cpf invalid")
	}

	return s.UserRepository.Create(ctx, u)
}

// Update call repository Update and returns
func (s *UserService) Update(ctx context.Context, id string, u *model.User) (*model.User, error) {

	if !u.BirthDate.IsZero() {
		if utils.IsUnderage(u.BirthDate) {
			return nil, rerrors.NewBadRequest("underage")
		}
	}

	if u.Cpf != "" {
		if !utils.IsBrazilianCPFValid(u.Cpf) {
			return nil, rerrors.NewBadRequest("cpf invalid")
		}
	}

	if u.Email != "" {
		_, err := mail.ParseAddress(u.Email)

		if err != nil {
			return nil, rerrors.NewBadRequest("invalid e-mail")
		}

	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return nil, rerrors.NewBadRequest("invalid id")
	}

	u.UID = uid

	return s.UserRepository.Update(ctx, u)
}
