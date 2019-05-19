package main

import (
	"fmt"
	"github.com/Anrop/Arma-Worlds-API/database"
	"github.com/Anrop/Arma-Worlds-API/server"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")

	if port == "" {
		port = "8080"
	}

	var err error

	database, err := database.New(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %q\n", err)
		os.Exit(1)
	}

	server, err := server.New(database)
	server.Serve(port)
}
