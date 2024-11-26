package database

import (
	"context"
	"log"

	"user-service/config"

	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/jackc/pgx/v5"
)

var DB *pgxpool.Pool

// InitSupabase initializes the database connection pool.
func InitSupabase() {
	// Fetch the DSN from the centralized config
	dsn := config.GetEnv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatalf("SUPABASE_DSN is not set in the environment")
	}

	// Parse the DSN configuration
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse database configuration: %v", err)
	}

	// Enable statement caching
	//config.ConnConfig.

	// Create a connection pool
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Println("Connected to Supabase!")
}

// CloseSupabase closes the database connection pool.
func CloseSupabase() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
