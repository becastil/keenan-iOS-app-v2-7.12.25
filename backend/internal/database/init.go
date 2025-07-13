package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sydney-health/backend/pkg/database"
)

// InitDB initializes the database connection from environment variables
func InitDB() (*database.DB, error) {
	// Read database configuration from environment
	host := getEnv("DB_HOST", "localhost")
	portStr := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "sydney_health")
	password := getEnv("DB_PASSWORD", "dev_password")
	dbName := getEnv("DB_NAME", "sydney_health")
	sslMode := getEnv("DB_SSL_MODE", "disable")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	config := database.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}

	db, err := database.NewConnection(config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}