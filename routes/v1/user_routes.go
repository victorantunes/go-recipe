package v1

import (
	"cql-backend/db"
	"cql-backend/handlers/v1"
	"cql-backend/services"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, store *db.Store) {
	userService := services.NewUserService(store.UserRepository)
	userHandler := v1.NewUserHandler(userService)

	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
}
