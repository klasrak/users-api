package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	model "github.com/klasrak/users-api/models"
)

// pgUserRepository is the repository implementation of UserRepository
type userRepository struct {
	DB *sqlx.DB
}

// NewUserRepository is a factory to initialize UserRepository
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// GetAll returns all users or error
func (r *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

// GetByID fetches user by ID or return error
func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return nil, nil
}

// Create a user
func (r *userRepository) Create(ctx context.Context, u *model.User) error {
	return nil
}

// Update a user
func (r *userRepository) Update(ctx context.Context, u *model.User) error {
	return nil
}

// Delete a user
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
