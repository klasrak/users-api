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

		t.Run("Error invalid id", func(t *testing.T) {
			mockUserRepository := new(mocks.MockUserRepository)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.GetByID(ctx, "invalid_id")

			mockUserRepository.AssertNotCalled(t, "GetByID")

			assert.Nil(t, us)
			assert.Error(t, err)
			assert.Equal(t, err, rerrors.NewBadRequest("invalid id"))

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Error not found", func(t *testing.T) {
			uid, _ := uuid.NewRandom()

			user := &model.User{}

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("GetByID", mock.AnythingOfType("*context.emptyCtx"), uid).Return(user, rerrors.NewNotFound("id", uid.String()))

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.GetByID(ctx, uid.String())

			mockUserRepository.AssertNumberOfCalls(t, "GetByID", 1)
			mockUserRepository.AssertCalled(t, "GetByID", mock.AnythingOfType("*context.emptyCtx"), uid)

			assert.Error(t, err)
			assert.Equal(t, err, rerrors.NewNotFound("id", uid.String()))
			assert.Equal(t, user, us)
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			user := &model.User{
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			userMockResponse := user
			uid, _ := uuid.NewRandom()
			userMockResponse.UID = uid

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), user).Return(userMockResponse, nil)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Create(ctx, user)

			mockUserRepository.AssertCalled(t, "Create", mock.AnythingOfType("*context.emptyCtx"), user)
			mockUserRepository.AssertNumberOfCalls(t, "Create", 1)

			assert.NoError(t, err)
			assert.Equal(t, userMockResponse, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Error unique violation email", func(t *testing.T) {
			user := &model.User{
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewConflict("user", "created", "unique_violation_email")

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), user).Return(nil, mockErrorResponse)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Create(ctx, user)

			mockUserRepository.AssertCalled(t, "Create", mock.AnythingOfType("*context.emptyCtx"), user)
			mockUserRepository.AssertNumberOfCalls(t, "Create", 1)

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Error unique violation cpf", func(t *testing.T) {
			user := &model.User{
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewConflict("user", "created", "unique_violation_cpf")

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), user).Return(nil, mockErrorResponse)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Create(ctx, user)

			mockUserRepository.AssertCalled(t, "Create", mock.AnythingOfType("*context.emptyCtx"), user)
			mockUserRepository.AssertNumberOfCalls(t, "Create", 1)

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Internal Server Error", func(t *testing.T) {
			user := &model.User{
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewInternal()

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), user).Return(nil, mockErrorResponse)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Create(ctx, user)

			mockUserRepository.AssertCalled(t, "Create", mock.AnythingOfType("*context.emptyCtx"), user)
			mockUserRepository.AssertNumberOfCalls(t, "Create", 1)

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Bad request underage", func(t *testing.T) {
			user := &model.User{
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Now(),
			}

			mockErrorResponse := rerrors.NewBadRequest("underage")

			mockUserRepository := new(mocks.MockUserRepository)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Create(ctx, user)

			mockUserRepository.AssertNotCalled(t, "Create")

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Bad request invalid cpf", func(t *testing.T) {
			user := &model.User{
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "invalid_cpf",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("cpf invalid")

			mockUserRepository := new(mocks.MockUserRepository)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Create(ctx, user)

			mockUserRepository.AssertNotCalled(t, "Create")

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			uid, _ := uuid.NewRandom()
			oldCpf := "313.716.772-80"
			oldBirthdate := time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC)

			userUpdateParams := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "",
				BirthDate: time.Time{},
			}

			userResponse := &model.User{
				UID:       uid,
				Name:      userUpdateParams.Name,
				Email:     userUpdateParams.Email,
				Cpf:       oldCpf,
				BirthDate: oldBirthdate,
			}

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), userUpdateParams).Return(userResponse, nil)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Update(ctx, uid.String(), userUpdateParams)

			mockUserRepository.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), userUpdateParams)
			mockUserRepository.AssertNumberOfCalls(t, "Update", 1)

			assert.NoError(t, err)
			assert.Equal(t, userResponse, us)
		})

		t.Run("Error unique violation email", func(t *testing.T) {
			uid, _ := uuid.NewRandom()
			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewConflict("user", "updated", "unique_violation_email")

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), user).Return(nil, mockErrorResponse)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Update(ctx, uid.String(), user)

			mockUserRepository.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), user)
			mockUserRepository.AssertNumberOfCalls(t, "Update", 1)

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Error unique violation cpf", func(t *testing.T) {
			uid, _ := uuid.NewRandom()
			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewConflict("user", "updated", "unique_violation_cpf")

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), user).Return(nil, mockErrorResponse)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Update(ctx, uid.String(), user)

			mockUserRepository.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), user)
			mockUserRepository.AssertNumberOfCalls(t, "Update", 1)

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Internal Server Error", func(t *testing.T) {
			uid, _ := uuid.NewRandom()
			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewInternal()

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), user).Return(nil, mockErrorResponse)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Update(ctx, uid.String(), user)

			mockUserRepository.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), user)
			mockUserRepository.AssertNumberOfCalls(t, "Update", 1)

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Bad request underage", func(t *testing.T) {
			uid, _ := uuid.NewRandom()
			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Now(),
			}

			mockErrorResponse := rerrors.NewBadRequest("underage")

			mockUserRepository := new(mocks.MockUserRepository)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Update(ctx, uid.String(), user)

			mockUserRepository.AssertNotCalled(t, "Update")

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Bad request invalid cpf", func(t *testing.T) {
			uid, _ := uuid.NewRandom()
			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "invalid_cpf",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("cpf invalid")

			mockUserRepository := new(mocks.MockUserRepository)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Update(ctx, uid.String(), user)

			mockUserRepository.AssertNotCalled(t, "Update")

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})

		t.Run("Bad request invalid email", func(t *testing.T) {
			uid, _ := uuid.NewRandom()
			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     "invalid_email",
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 0, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("invalid e-mail")

			mockUserRepository := new(mocks.MockUserRepository)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			us, err := userService.Update(ctx, uid.String(), user)

			mockUserRepository.AssertNotCalled(t, "Update")

			assert.Error(t, err)
			assert.Equal(t, mockErrorResponse, err)
			assert.Nil(t, us)

			mockUserRepository.AssertExpectations(t)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			uid, _ := uuid.NewRandom()

			mockUserRepository := new(mocks.MockUserRepository)
			mockUserRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"), uid.String()).Return(nil)

			userService := &UserService{
				UserRepository: mockUserRepository,
			}

			ctx := context.Background()

			err := userService.Delete(ctx, uid.String())

			mockUserRepository.AssertCalled(t, "Delete", mock.AnythingOfType("*context.emptyCtx"), uid.String())
			mockUserRepository.AssertNumberOfCalls(t, "Delete", 1)

			assert.NoError(t, err)
		})
	})
}
