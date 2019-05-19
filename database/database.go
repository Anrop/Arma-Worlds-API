package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	client *sql.DB
}

func New(url string) (database *Database, err error) {
	client, err := sql.Open("postgres", url)

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
