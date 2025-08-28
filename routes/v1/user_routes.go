package v1

import (
	"go-recipe/db"
	"go-recipe/handlers/v1"
	"go-recipe/services"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, store *db.DataStore) {
	userService := services.NewUserService(store.UserRepository)
	userHandler := v1.NewUserHandler(userService)

	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
}
