package api

import (
	//"net/http"

	"user-service/api/user"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", user.RegisterHandler).Methods("POST")

	return router
}
