package config

import (
	"os"
)

// Config model
type Config struct {
	DatabaseURL         string
	LocationsURL        string
	Port                string
	SatelliteTilesURL   string
	TopographicTilesURL string
}

// FromEnv creates a configuration based on environment variables
func FromEnv() (config Config) {
	config = Config{
		DatabaseURL:         os.Getenv("DATABASE_URL"),
		LocationsURL:        os.Getenv("LOCATIONS_URL"),
		Port:                os.Getenv("PORT"),
		SatelliteTilesURL:   os.Getenv("SATELLITE_TILES_URL"),
		TopographicTilesURL: os.Getenv("TOPOGRAPHIC_TILES_URL"),
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	return config
}
