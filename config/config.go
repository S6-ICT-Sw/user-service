package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig ensures the .env file is loaded at application startup.
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Configuration loaded successfully")
}

// GetEnv fetches the value of a specific environment variable.
func GetEnv(key string) string {
	return os.Getenv(key)
}
