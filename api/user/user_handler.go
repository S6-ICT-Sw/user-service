package user

import (
	"context"
	"encoding/json"
	"net/http"

	"user-service/models"
	"user-service/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse JSON body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Call the service to register the user
	err = services.Register(context.Background(), &user)
	if err != nil {
		//http.Error(w, "Failed to create user", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))
}
