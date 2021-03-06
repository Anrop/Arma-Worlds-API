package database

import (
	"database/sql"
	"github.com/Anrop/Arma-Worlds-API/config"
	// Load postgres integration for sql
	_ "github.com/lib/pq"
	"time"
)

// Database object wrapping the DB connection
type Database struct {
	client *sql.DB
}

// New creates a new database connection based on configuration
func New(config config.Config) (database *Database, err error) {
	client, err := sql.Open("postgres", config.DatabaseURL)

	if err != nil {
		return nil, err
	}

	err = client.Ping()
	if err != nil {
		return nil, err
	}

	client.SetConnMaxLifetime(15 * time.Minute)
	client.SetMaxOpenConns(4)
	client.SetMaxIdleConns(2)

	database = &Database{
		client: client,
	}

	return database, err
}
