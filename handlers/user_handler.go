package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/klasrak/users-api/models"
	"github.com/klasrak/users-api/rerrors"
)

type createPayload struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Cpf       string    `json:"cpf" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
}

type updatePayload struct {
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Cpf       string    `json:"cpf,omitempty"`
	Birthdate time.Time `json:"birthdate,omitempty"`
}

// GetAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept  json
// @Produce  json
// @Param name query string false "search by name"
// @Success 200 {object} []model.User
// @Router /users [get]
func (h *Handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	name := c.Query("name")

	users, err := h.UserService.GetAll(ctx, name)

	if err != nil {
		log.Printf("Failed to get all users: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetByID godoc
// @Summary Get a single user by ID
// @Description Get a single user by ID
// @Tags user
// @ID string
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	user, err := h.UserService.GetByID(ctx, id)

	if err != nil {
		log.Printf("Failed to get user: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

// Create godoc
// @Summary Create user
// @Description Add user to database
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body createPayload true "Add user"
// @Success 201 {object} model.User
// @Router /users [post]
func (h *Handler) Create(c *gin.Context) {
	var req createPayload

	// Bind incoming json to struct and check for validation errors
	ok := bindData(c, &req)

	if !ok {
		log.Println("failed to bind data")
		return
	}

	u := &model.User{
		Name:      req.Name,
		Email:     req.Email,
		Cpf:       req.Cpf,
		BirthDate: req.Birthdate,
	}

	ctx := c.Request.Context()

	user, err := h.UserService.Create(ctx, u)

	if err != nil {
		log.Printf("failed to create user: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Update godoc
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body updatePayload false "Update user"
// @Success 200 {object} model.User
// @Router /users/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	var req updatePayload

	id := c.Param("id")

	if id == "" {
		err := rerrors.NewBadRequest("invalid id")
		log.Printf("invalid ID: %v\n", err)

		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	// Bind incoming json to struct and check for validation errors
	ok := bindData(c, &req)

	if !ok {
		log.Println("failed to bind data")
		return
	}

	u := &model.User{
		Name:      req.Name,
		Email:     req.Email,
		Cpf:       req.Cpf,
		BirthDate: req.Birthdate,
	}

	ctx := c.Request.Context()

	user, err := h.UserService.Update(ctx, id, u)

	if err != nil {
		log.Printf("failed to create user: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Delete godoc
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		err := rerrors.NewBadRequest("missing user ID id")
		log.Printf("missing user ID ID: %v\n", err)

		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx := c.Request.Context()

	err := h.UserService.Delete(ctx, id)

	if err != nil {
		log.Printf("failed to create user: %v\n", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusNoContent, nil)
}
