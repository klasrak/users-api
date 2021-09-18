package mocks

import (
	"context"

	"github.com/google/uuid"
	model "github.com/klasrak/users-api/models"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock type for service.UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

// GetAll is a mock for UserRepository GetAll
func (m *MockUserRepository) GetAll(ctx context.Context, name string) ([]model.User, error) {
	ret := m.Called(ctx, name)

	var r0 []model.User

	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// GetByID is a mock for UserRepository GetByID
func (m *MockUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	ret := m.Called(ctx, id)

	var r0 *model.User

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// Create is a mock for UserRepository Create
func (m *MockUserRepository) Create(ctx context.Context, u *model.User) error {
	ret := m.Called(ctx, u)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

// Update is a mock for UserRepository Update
func (m *MockUserRepository) Update(ctx context.Context, u *model.User) error {
	ret := m.Called(ctx, u)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

// Delete is a mock for UserRepository Delete
func (m *MockUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := m.Called(ctx, id)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
