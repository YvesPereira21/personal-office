package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the runtime configuration loaded from environment variables.
type Config struct {
	Port        string
	DatabaseURL string
}

// Load reads settings from a local .env file when present and validates required environment variables.
func Load() (*Config, error) {
	// Ignore missing .env so production environments can provide variables directly.
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("variável DATABASE_URL não definida")
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil
}
