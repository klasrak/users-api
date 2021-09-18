package service

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/klasrak/users-api/mocks"
	model "github.com/klasrak/users-api/models"
	"github.com/klasrak/users-api/rerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAll", func(t *testing.T) {
		t.Run("Success without name filter", func(t *testing.T) {
			var users []model.User

			err := faker.FakeData(&users)

			assert.NoError(t, err)

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("GetAll", mock.AnythingOfType("*context.emptyCtx"), "").Return(users, nil)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.GetAll(ctx, "")

			mockUserRepository.AssertNumberOfCalls(t, "GetAll", 1)
			mockUserRepository.AssertCalled(t, "GetAll", mock.AnythingOfType("*context.emptyCtx"), "")

			assert.NoError(t, err)
			assert.Equal(t, users, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Success with name filter", func(t *testing.T) {
			var users []model.User

			user := model.User{
				UID:       uuid.New(),
				Name:      "John Doe",
				Email:     faker.Email(),
				Cpf:       "123.456.789-10",
				BirthDate: time.Now(),
			}

			users = append(users, user)

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("GetAll", mock.AnythingOfType("*context.emptyCtx"), "John").Return(users, nil)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.GetAll(ctx, "John")

			mockUserRepository.AssertNumberOfCalls(t, "GetAll", 1)
			mockUserRepository.AssertCalled(t, "GetAll", mock.AnythingOfType("*context.emptyCtx"), "John")

			assert.NoError(t, err)
			assert.Equal(t, users, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Internal Server Error", func(t *testing.T) {

			var users []model.User

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("GetAll", mock.Anything, mock.Anything).Return(users, rerrors.NewInternal())

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.GetAll(ctx, "")

			mockUserRepository.AssertNumberOfCalls(t, "GetAll", 1)
			mockUserRepository.AssertCalled(t, "GetAll", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"))

			assert.Error(t, err)
			assert.Equal(t, rerrors.NewInternal(), err)
			assert.Equal(t, users, us)

			mockUserRepository.AssertExpectations(t)
		})
	})

	t.Run("GetByID", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			uid, _ := uuid.NewRandom()

			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "123.456.789-10",
				BirthDate: time.Now(),
			}

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("GetByID", mock.AnythingOfType("*context.emptyCtx"), uid).Return(user, nil)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.GetByID(ctx, uid.String())

			mockUserRepository.AssertNumberOfCalls(t, "GetByID", 1)
			mockUserRepository.AssertCalled(t, "GetByID", mock.AnythingOfType("*context.emptyCtx"), uid)

			assert.NoError(t, err)
			assert.Equal(t, user, us)
		})
	})
}
