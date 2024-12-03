package api

import (
	"user-service/api/user"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", user.RegisterHandler).Methods("POST")
	router.HandleFunc("/users/{id}", user.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", user.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users/{id}", user.EditUserHandler).Methods("PUT")

	return router
}
