package mocks

import (
	"context"

	model "github.com/klasrak/users-api/models"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

// GetAll is a mock for UserService GetAll
func (m *MockUserService) GetAll(ctx context.Context, name string) ([]model.User, error) {
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

// GetByID is a mock for UserService GetByID
func (m *MockUserService) GetByID(ctx context.Context, id string) (*model.User, error) {
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

// Create is a mock for UserService Create
func (m *MockUserService) Create(ctx context.Context, u *model.User) (*model.User, error) {
	ret := m.Called(ctx, u)

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

// Update is a mock for UserService Update
func (m *MockUserService) Update(ctx context.Context, id string, u *model.User) (*model.User, error) {
	ret := m.Called(ctx, id, u)

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

// Delete is a mock for UserService Delete
func (m *MockUserService) Delete(ctx context.Context, id string) error {
	ret := m.Called(ctx, id)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
