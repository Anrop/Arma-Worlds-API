package config

import (
	"os"
)

type Config struct {
	DatabaseURL             string
	Port                    string
	SatelliteTilesBaseURL   string
	TopographicTilesBaseURL string
}

func FromEnv() (config Config) {
	config = Config{
		DatabaseURL:             os.Getenv("DATABASE_URL"),
		Port:                    os.Getenv("PORT"),
		SatelliteTilesBaseURL:   os.Getenv("SATELLITE_TILES_BASE_URL"),
		TopographicTilesBaseURL: os.Getenv("TOPOGRAPHIC_TILES_BASE_URL"),
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	return config
}
