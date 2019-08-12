package main

import (
	"fmt"
	"github.com/Anrop/Arma-Worlds-API/config"
	"github.com/Anrop/Arma-Worlds-API/database"
	"github.com/Anrop/Arma-Worlds-API/server"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config.FromEnv()

	var err error

	database, err := database.New(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %q\n", err)
		os.Exit(1)
	}

	server, err := server.New(&config, database)
	server.Serve(config.Port)
}
