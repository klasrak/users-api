package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Repository combines all repositories
type Repository struct {
	UserRepository *UserRepository
}

// CreateRepository create a implementation of repository with all injected dependencies
func CreateRepository(options *Options) (*Repository, error) {
	return &Repository{
		UserRepository: &UserRepository{
			DB: options.DB,
		},
	}, nil
}

// Options is a utility to define all dependencies and parameters to inject
type Options struct {
	DB *sqlx.DB
}
