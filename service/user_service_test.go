package service

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
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
}
