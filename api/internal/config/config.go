package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	ServerConfig   ServerConfig
}

type DatabaseConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

type ServerConfig struct {
	ServerPort string
}

func NewConfig() *Config {
	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatal(fmt.Errorf(".env file not found or could not be loaded: %w", err))
	}

	databaseConfig, err := initializeDatabaseConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize database config: %w", err))
	}

	serverConfig, err := initializeServerConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize server config: %w", err))
	}

	return &Config{
		DatabaseConfig: *databaseConfig,
		ServerConfig:   *serverConfig,
	}
}

func initializeDatabaseConfig() (*DatabaseConfig, error) {
	dbUser := getEnv("DB_USER", "")
	dbPass := getEnv("DB_PASS", "")
	dbHost := getEnv("DB_HOST", "")
	dbPort := getEnv("DB_PORT", "")
	dbName := getEnv("DB_NAME", "")

	return &DatabaseConfig{
		DBUser: dbUser,
		DBPass: dbPass,
		DBHost: dbHost,
		DBPort: dbPort,
		DBName: dbName,
	}, nil
}

func initializeServerConfig() (*ServerConfig, error) {
	serverPort := getEnv("SERVER_PORT", "8080")

	return &ServerConfig{
		ServerPort: serverPort,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if defaultValue == "" {
		log.Fatal(fmt.Errorf("environment variable %s is required but not set", key))
	}

	return defaultValue
}
