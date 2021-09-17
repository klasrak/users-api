package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	model "github.com/klasrak/users-api/models"
	_ "github.com/lib/pq"
)

// UserRepository is a repository implementation of service layer UserRepository interface
type UserRepository struct {
	DB *sqlx.DB
}

// GetAll returns all users or error
func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

// GetByID fetches user by ID or return error
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return nil, nil
}

// Create a user
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	return nil
}

// Update a user
func (r *UserRepository) Update(ctx context.Context, u *model.User) error {
	return nil
}

// Delete a user
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
