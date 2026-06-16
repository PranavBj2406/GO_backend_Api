package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment variables.
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	AppPort    int
}

// LoadConfig loads environment variables (from .env if present) and returns a Config.
func LoadConfig() (*Config, error) {
	// Load .env file if present; ignore error if file not found
	_ = godotenv.Load()

	dbHost := getEnv("DB_HOST", "localhost")
	dbPortStr := getEnv("DB_PORT", "5432")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "userdb")
	dbSSL := getEnv("DB_SSLMODE", "disable")

	appPortStr := getEnv("APP_PORT", "8080")
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid APP_PORT: %w", err)
	}

	return &Config{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBSSLMode:  dbSSL,
		AppPort:    appPort,
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
