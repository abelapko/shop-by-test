package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold the configuration data
type Config struct {
	DatabaseURL     string
	TestDatabaseURL string
	Port            string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists
	dotenvPath := os.Getenv("DOTENV_PATH")
	if dotenvPath == "" {
		log.Fatal("DOTENV_PATH is not set")
	}

	err := godotenv.Load(dotenvPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Assemble the database URLs from the environment variables
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		getEnv("MYSQL_USER", "user"),
		getEnv("MYSQL_PASSWORD", "userpassword"),
		getEnv("MYSQL_HOST", "mysql"),
		getEnv("MYSQL_PORT", "3306"),
		getEnv("MYSQL_DATABASE", "currency_service"))

	testDatabaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root",
		getEnv("MYSQL_ROOT_PASSWORD", "root"),
		getEnv("MYSQL_HOST", "mysql"),
		getEnv("MYSQL_PORT", "3306"),
		getEnv("MYSQL_DATABASE_TEST", "currency_service_test"))

	// Return the loaded config
	return &Config{
		DatabaseURL:     databaseURL,
		TestDatabaseURL: testDatabaseURL,
		Port:            getEnv("PORT", "8080"),
	}
}

// Helper function to get the environment variable or fallback to default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Check if required configuration variables are set
func (c *Config) Check() {
	if c.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	if c.TestDatabaseURL == "" {
		log.Fatal("TEST_DATABASE_URL is required")
	}
	if c.Port == "" {
		log.Fatal("PORT is required")
	}
}
