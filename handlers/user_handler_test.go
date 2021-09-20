package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/klasrak/users-api/mocks"
	model "github.com/klasrak/users-api/models"
	"github.com/klasrak/users-api/rerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedContainer struct {
	Handler *Handler
}

type MockedRouter struct {
	r *gin.Engine
}

func (router *MockedRouter) Initialize(c *MockedContainer) {
	// Handlers
	h := c.Handler

	// Default gin engine instance
	r := gin.Default()

	// ####### MIDDLEWARES #######
	// CORS
	r.Use(cors.Default())

	// ####### API V1 #######
	v1Group := r.Group("/api/v1")

	// ---- USERS RESOURCES /users ----
	usersGroup := v1Group.Group("/users")

	// ## GET ##
	usersGroup.GET("", h.GetAll)
	usersGroup.GET("/:id", h.GetByID)

	// ## POST ##
	usersGroup.POST("", h.Create)

	// ## PUT ##
	usersGroup.PUT("/:id", h.Update)

	// ## DELETE ##
	usersGroup.DELETE("/:id", h.Delete)

	// ####### inject implementation of gin engine #######
	router.r = r
}

func TestUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAll", func(t *testing.T) {
		t.Run("Success without name filter", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			var users []model.User

			err := faker.FakeData(&users)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users", nil)
			request.Header.Set("Content-Type", "application/json")

			mockUserService.On("GetAll", request.Context(), "").Return(users, nil)

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertNumberOfCalls(t, "GetAll", 1)
			mockUserService.AssertCalled(t, "GetAll", request.Context(), "")

			respBody, _ := json.Marshal(users)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, respBody, rr.Body.Bytes())
			mockUserService.AssertExpectations(t)
		})

		t.Run("Success with name filter", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			var users []model.User

			uid, _ := uuid.NewRandom()
			user := model.User{
				UID:       uid,
				Name:      "John Doe",
				Email:     faker.Email(),
				Cpf:       "123.456.789-10",
				BirthDate: time.Date(1993, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			users = append(users, user)

			// err := faker.FakeData(&users)
			// assert.NoError(t, err)

			rr := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users", nil)
			q := request.URL.Query()
			q.Add("name", "John Doe")
			request.URL.RawQuery = q.Encode()
			request.Header.Set("Content-Type", "application/json")

			mockUserService.On("GetAll", mock.AnythingOfType("*context.emptyCtx"), "John Doe").Return(users, nil)

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertNumberOfCalls(t, "GetAll", 1)
			mockUserService.AssertCalled(t, "GetAll", mock.AnythingOfType("*context.emptyCtx"), "John Doe")

			respBody, _ := json.Marshal(users)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, respBody, rr.Body.Bytes())
			mockUserService.AssertExpectations(t)
		})

		t.Run("Success no content", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			var users []model.User

			rr := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users", nil)

			request.Header.Set("Content-Type", "application/json")

			mockUserService.On("GetAll", mock.AnythingOfType("*context.emptyCtx"), "").Return(users, nil)

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertNumberOfCalls(t, "GetAll", 1)
			mockUserService.AssertCalled(t, "GetAll", mock.AnythingOfType("*context.emptyCtx"), "")

			assert.Equal(t, http.StatusNoContent, rr.Code)
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			var users []model.User

			rr := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/users", nil)

			request.Header.Set("Content-Type", "application/json")

			mockErrorResponse := rerrors.NewInternal()

			mockUserService.On("GetAll", mock.AnythingOfType("*context.emptyCtx"), "").Return(users, mockErrorResponse)

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertNumberOfCalls(t, "GetAll", 1)
			mockUserService.AssertCalled(t, "GetAll", mock.AnythingOfType("*context.emptyCtx"), "")

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
			mockUserService.AssertExpectations(t)
		})
	})

	t.Run("GetByID", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			user := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "123.456.789-10",
				BirthDate: time.Date(2003, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			mockUserService.On("GetByID", mock.AnythingOfType("*context.emptyCtx"), uid.String()).Return(user, nil)

			rr := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", uid.String()), nil)

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "GetByID", mock.AnythingOfType("*context.emptyCtx"), uid.String())
			mockUserService.AssertNumberOfCalls(t, "GetByID", 1)

			respBody, _ := json.Marshal(user)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, respBody, rr.Body.Bytes())
			mockUserService.AssertExpectations(t)
		})

		t.Run("Invalid ID", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			mockErrorResponse := rerrors.NewBadRequest("invalid id")

			mockUserService.On("GetByID", mock.AnythingOfType("*context.emptyCtx"), "invalid_id").Return(nil, mockErrorResponse)

			rr := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", "invalid_id"), nil)

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "GetByID", mock.AnythingOfType("*context.emptyCtx"), "invalid_id")
			mockUserService.AssertNumberOfCalls(t, "GetByID", 1)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error ID not found", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			mockErrorResponse := rerrors.NewNotFound("id", uid.String())

			mockUserService.On("GetByID", mock.AnythingOfType("*context.emptyCtx"), "invalid_id").Return(nil, mockErrorResponse)

			rr := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", "invalid_id"), nil)

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "GetByID", mock.AnythingOfType("*context.emptyCtx"), "invalid_id")
			mockUserService.AssertNumberOfCalls(t, "GetByID", 1)

			assert.Equal(t, http.StatusNotFound, rr.Code)
			mockUserService.AssertExpectations(t)
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			u := &model.User{
				Name:      "John Doe",
				Email:     "test@mail.com",
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(2003, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			createdUser := &model.User{
				UID:       uid,
				Email:     u.Email,
				Cpf:       u.Cpf,
				BirthDate: u.BirthDate,
			}

			mockUserService.On("Create", mock.AnythingOfType("*context.emptyCtx"), u).Return(createdUser, nil)

			rr := httptest.NewRecorder()

			body, err := json.Marshal(gin.H{
				"name":      u.Name,
				"email":     u.Email,
				"cpf":       u.Cpf,
				"birthdate": u.BirthDate,
			})

			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users", bytes.NewBuffer(body))

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Create", mock.AnythingOfType("*context.emptyCtx"), u)
			mockUserService.AssertNumberOfCalls(t, "Create", 1)

			u.UID = uid
			respBody, _ := json.Marshal(createdUser)

			assert.Equal(t, http.StatusCreated, rr.Code)
			assert.Equal(t, respBody, rr.Body.Bytes())
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error underage", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			u := &model.User{
				Name:      "John Doe",
				Email:     "test@mail.com",
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(2019, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("underage")

			mockUserService.On("Create", mock.AnythingOfType("*context.emptyCtx"), u).Return(nil, mockErrorResponse)

			rr := httptest.NewRecorder()

			body, err := json.Marshal(gin.H{
				"name":      u.Name,
				"email":     u.Email,
				"cpf":       u.Cpf,
				"birthdate": u.BirthDate,
			})

			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users", bytes.NewBuffer(body))

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Create", mock.AnythingOfType("*context.emptyCtx"), u)
			mockUserService.AssertNumberOfCalls(t, "Create", 1)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error invalid cpf", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			u := &model.User{
				Name:      "John Doe",
				Email:     "test@mail.com",
				Cpf:       "invalid_cpf",
				BirthDate: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("cpf invalid")

			mockUserService.On("Create", mock.AnythingOfType("*context.emptyCtx"), u).Return(nil, mockErrorResponse)

			rr := httptest.NewRecorder()

			body, err := json.Marshal(gin.H{
				"name":      u.Name,
				"email":     u.Email,
				"cpf":       u.Cpf,
				"birthdate": u.BirthDate,
			})

			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users", bytes.NewBuffer(body))

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Create", mock.AnythingOfType("*context.emptyCtx"), u)
			mockUserService.AssertNumberOfCalls(t, "Create", 1)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			mockUserService.AssertExpectations(t)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			oldCpf := "313.716.772-80"
			oldBirthdate := time.Date(2003, 1, 1, 1, 1, 1, 1, time.UTC)

			u := &model.User{
				Name:      "John Doe",
				Email:     "test@mail.com",
				Cpf:       "",
				BirthDate: time.Time{},
			}

			updatedUser := &model.User{
				UID:       uid,
				Name:      u.Name,
				Email:     u.Email,
				Cpf:       oldCpf,
				BirthDate: oldBirthdate,
			}

			mockUserService.On("Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u).Return(updatedUser, nil)

			rr := httptest.NewRecorder()

			body, err := json.Marshal(gin.H{
				"name":  u.Name,
				"email": u.Email,
			})

			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", uid.String()), bytes.NewBuffer(body))

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u)
			mockUserService.AssertNumberOfCalls(t, "Update", 1)

			u.UID = uid
			respBody, _ := json.Marshal(updatedUser)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, respBody, rr.Body.Bytes())
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error underage", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			u := &model.User{
				Name:      "John Doe",
				Email:     "test@mail.com",
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(2019, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("underage")

			mockUserService.On("Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u).Return(nil, mockErrorResponse)

			rr := httptest.NewRecorder()

			body, err := json.Marshal(gin.H{
				"name":      u.Name,
				"email":     u.Email,
				"cpf":       u.Cpf,
				"birthdate": u.BirthDate,
			})

			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", uid.String()), bytes.NewBuffer(body))

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u)
			mockUserService.AssertNumberOfCalls(t, "Update", 1)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error invalid email", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			u := &model.User{
				Name:      "John Doe",
				Email:     "invalid_email",
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("invalid email")

			mockUserService.On("Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u).Return(nil, mockErrorResponse)

			rr := httptest.NewRecorder()

			body, err := json.Marshal(gin.H{
				"name":      u.Name,
				"email":     u.Email,
				"cpf":       u.Cpf,
				"birthdate": u.BirthDate,
			})

			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", uid.String()), bytes.NewBuffer(body))

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u)
			mockUserService.AssertNumberOfCalls(t, "Update", 1)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error invalid cpf", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			u := &model.User{
				Name:      "John Doe",
				Email:     "test@test.com",
				Cpf:       "invalid_cpf",
				BirthDate: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
			}

			mockErrorResponse := rerrors.NewBadRequest("cpf invalid")

			mockUserService.On("Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u).Return(nil, mockErrorResponse)

			rr := httptest.NewRecorder()

			body, err := json.Marshal(gin.H{
				"name":      u.Name,
				"email":     u.Email,
				"cpf":       u.Cpf,
				"birthdate": u.BirthDate,
			})

			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", uid.String()), bytes.NewBuffer(body))

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Update", mock.AnythingOfType("*context.emptyCtx"), uid.String(), u)
			mockUserService.AssertNumberOfCalls(t, "Update", 1)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			mockUserService.AssertExpectations(t)
		})

	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			mockUserService.On("Delete", mock.AnythingOfType("*context.emptyCtx"), uid.String()).Return(nil)

			rr := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", uid.String()), nil)

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Delete", mock.AnythingOfType("*context.emptyCtx"), uid.String())
			mockUserService.AssertNumberOfCalls(t, "Delete", 1)

			assert.Equal(t, http.StatusNoContent, rr.Code)
			mockUserService.AssertExpectations(t)
		})

		t.Run("Error", func(t *testing.T) {
			mockUserService := new(mocks.MockUserService)

			h := &Handler{
				UserService: mockUserService,
			}

			c := &MockedContainer{
				Handler: h,
			}

			router := &MockedRouter{}

			router.Initialize(c)

			uid, err := uuid.NewRandom()
			assert.NoError(t, err)

			mockErrorResponse := rerrors.NewInternal()

			mockUserService.On("Delete", mock.AnythingOfType("*context.emptyCtx"), uid.String()).Return(mockErrorResponse)

			rr := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", uid.String()), nil)

			request.Header.Set("Content-Type", "application/json")

			router.r.ServeHTTP(rr, request)

			mockUserService.AssertCalled(t, "Delete", mock.AnythingOfType("*context.emptyCtx"), uid.String())
			mockUserService.AssertNumberOfCalls(t, "Delete", 1)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
			mockUserService.AssertExpectations(t)
		})
	})
}
