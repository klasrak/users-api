package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	model "github.com/klasrak/users-api/models"
	"github.com/klasrak/users-api/rerrors"
	"github.com/klasrak/users-api/utils"
	"github.com/lib/pq"
)

// UserRepository is a repository implementation of service layer UserRepository interface
type UserRepository struct {
	DB *sqlx.DB
}

// GetAll returns all users or error
func (r *UserRepository) GetAll(ctx context.Context, name string) ([]model.User, error) {
	users := []model.User{}

	query := "SELECT * FROM users u %s;"

	if name != "" {
		query = fmt.Sprintf(query, fmt.Sprintf(`WHERE u.name LIKE '%%%s%%' `, name))
	} else {
		query = strings.Trim(fmt.Sprintf(query, ""), "")
	}

	if err := r.DB.SelectContext(ctx, &users, query); err != nil {

		return users, rerrors.NewInternal()
	}

	return users, nil
}

// GetByID fetches user by ID or return error
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE id=$1;"

	if err := r.DB.Get(user, query, id); err != nil {
		return user, rerrors.NewNotFound("id", id.String())
	}

	return user, nil
}

// Create a user
func (r *UserRepository) Create(ctx context.Context, u *model.User) (*model.User, error) {
	query := "INSERT INTO users (name, email, cpf, birthdate) VALUES ($1, $2, $3, $4) RETURNING *;"

	if err := r.DB.GetContext(ctx, u, query, u.Name, u.Email, u.Cpf, u.BirthDate); err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("could not create user. Reason: %v\n", err.Error())
			return nil, rerrors.NewConflict("user", err.Detail)
		}

		log.Printf("failed to create user. Reason: %v\n", err)
		return nil, rerrors.NewInternal()
	}

	return u, nil
}

// Update a user
func (r *UserRepository) Update(ctx context.Context, u *model.User) (*model.User, error) {

	query := `
	UPDATE users u SET
		name = COALESCE(:name, u."name"),
		email = COALESCE(:email, u.email),
		cpf = COALESCE(:cpf, u.cpf),
		birthdate = COALESCE(:birthdate, u.birthdate)
	WHERE u.id = :id
	RETURNING *;
	`

	user, err := utils.SanitizeUpdateParams(u)

	if err != nil {
		return nil, err
	}

	nstmt, err := r.DB.PrepareNamedContext(ctx, query)

	if err != nil {
		log.Printf("unable to prepare user update query: %v\n", err)
		return nil, rerrors.NewInternal()
	}

	if err := nstmt.GetContext(ctx, u, user); err != nil {
		log.Printf("failed to update user: %v. Reason: %v\n", err, u)
		return nil, rerrors.NewInternal()
	}

	return u, err
}

// Delete a user
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
