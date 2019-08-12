package database

import (
	"database/sql"
	"github.com/Anrop/Arma-Worlds-API/config"
	_ "github.com/lib/pq"
)

type Database struct {
	client *sql.DB
}

func New(config config.Config) (database *Database, err error) {
	client, err := sql.Open("postgres", config.DatabaseURL)

	if err != nil {
		return nil, err
	}

	err = client.Ping()
	if err != nil {
		return nil, err
	}

	database = &Database{
		client: client,
	}

	return database, err
}
