package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseSources struct {
	DB *sqlx.DB
}

func (ds *DatabaseSources) Initialize() error {
	log.Println("Connecting to database")

	// ####### PostgreSQL #######
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DATABASE")
	pgSSL := os.Getenv("POSTGRES_SSL")

	pgConnectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgHost,
		pgPort,
		pgUser,
		pgPassword,
		pgDB,
		pgSSL,
	)

	log.Println("Starting postgres connection")
	db, err := sqlx.Open("postgres", pgConnectionString)

	if err != nil {
		return fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error connecting to db: %w", err)
	}

	ds.DB = db

	return nil
}

func (ds *DatabaseSources) Close() error {
	if err := ds.DB.Close(); err != nil {
		return fmt.Errorf("error closing postgres: %w", err)
	}

	return nil
}
