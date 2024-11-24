package main

import (
	//"fmt"
	"log"
	"net/http"
	"user-service/api"
	"user-service/config"
	"user-service/pkg/database"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize the database
	database.InitSupabase()
	defer database.CloseSupabase() // Ensure proper cleanup on shutdown

	r := api.SetupRouter()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
