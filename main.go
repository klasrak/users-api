package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

// @title Users API
// @version 1.0
// @description This is a Users crud server.

// @license.name MIT

// @BasePath /v1
func main() {
	log.Println("Starting server...")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file\n")
	}

	// initialize database sources
	ds := &DatabaseSources{}

	if err := ds.Initialize(); err != nil {
		log.Fatalf("unable to initialize database sources: %v\n", err)
	}

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	c := &Container{}

	if err := c.Initialize(ds); err != nil {
		log.Fatalf("unable to initialize services via dependency injection: %v\n", err)
	}

	router := Router{}

	router.Initialize(c)

	// generate swagger docs
	generateSwaggerDocs(&router.r.RouterGroup)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.r,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal, 2)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown database sources
	if err := ds.Close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources: %v\n", err)
	}

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
