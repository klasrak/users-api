package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klasrak/users-api/rerrors"
)

// GetAll godoc
// @Summary Get all users
// @Description Get all users
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
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetByID godoc
// @Summary Get a single user by ID
// @Description Get a single user by ID
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
