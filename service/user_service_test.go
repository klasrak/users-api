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
			mockUserRepository.On("GetAll", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).Return(users, nil)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.GetAll(ctx, "")

			assert.NoError(t, err)
			assert.Equal(t, users, us)

			mockUserRepository.AssertExpectations(t)
		})
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
		mockUserRepository.On("GetAll", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(users, nil)

		userService := &UserService{
			UserRepository: mockUserRepository,
		}

		ctx := context.Background()

		us, err := userService.GetAll(ctx, "John")

		assert.NoError(t, err)
		assert.Equal(t, users, us)

		mockUserRepository.AssertExpectations(t)
	})
}
