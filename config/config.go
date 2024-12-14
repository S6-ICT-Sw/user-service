package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig ensures the .env file is loaded at application startup.
func LoadConfig() {
	// Check if .env file exists before loading
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		log.Println("Configuration loaded from .env file")
	} else {
		log.Println(".env file not found, using environment variables")
	}

}

// GetEnv fetches the value of a specific environment variable.
func GetEnv(key string) string {
	return os.Getenv(key)
}
