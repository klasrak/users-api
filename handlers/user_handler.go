package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/klasrak/users-api/rerrors"
)

func (h *Handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.UserService.GetAll(ctx)

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

func (h *Handler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	uid, _ := uuid.Parse(id)

	user, err := h.UserService.GetByID(ctx, uid)

	if err != nil {
		log.Printf("Failed to user: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
