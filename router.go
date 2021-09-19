package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/klasrak/users-api/docs"
)

// Router represents a reference to a router
type Router struct {
	r *gin.Engine
}

func (router *Router) Initialize(c *Container) {
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

func generateSwaggerDocs(rg *gin.RouterGroup) {

	docsURL := fmt.Sprintf("http://%s:%s/docs/doc.json", os.Getenv("DOMAIN"), os.Getenv("PORT"))

	rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(docsURL)))
}
