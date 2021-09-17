package main

import "github.com/gin-gonic/gin"

// Router represents a reference to a router
type Router struct {
	r *gin.Engine
}

func (router *Router) Initialize(c *Container) {
	// Handlers
	h := c.Handler

	// Default gin engine instance
	r := gin.Default()

	// ####### API V1 #######
	v1Group := r.Group("/api/v1")

	// ---- USERS RESOURCES /users ----
	usersGroup := v1Group.Group("/users")

	usersGroup.GET("", h.GetAll)

	// ####### inject implementation of gin engine
	router.r = r
}
