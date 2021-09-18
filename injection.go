package main

import (
	"fmt"
	"log"

	"github.com/klasrak/users-api/handlers"
	"github.com/klasrak/users-api/repository"
	"github.com/klasrak/users-api/service"
)

// Container used for injecting dependencies
type Container struct {
	Handler *handlers.Handler
}

// Initialize implementation of service and repository layers
func (c *Container) Initialize(ds *DatabaseSources) error {
	log.Println("Injecting dependencies")

	// container for initialize repositories
	r, err := repository.CreateRepository(&repository.Options{
		DB: ds.DB,
	})

	if err != nil {
		return fmt.Errorf("could not initialize database sources (PostgreSQL): %w", err)
	}

	// create UserService with a implementation of UserRepository
	userService := &service.UserService{
		UserRepository: r.UserRepository,
	}

	// create handler container with a implementation of UserService
	c.Handler = &handlers.Handler{
		UserService: userService,
	}

	return nil
}
