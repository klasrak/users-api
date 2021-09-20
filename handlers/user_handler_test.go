package handlers

import (
	"encoding/json"
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
		})
	})
}
