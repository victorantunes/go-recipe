package v1

import (
	"cql-backend/handlers/v1"
	"github.com/gorilla/mux"
)

// RegisterHealthRoutes registers the health/ping endpoints.
func RegisterHealthRoutes(r *mux.Router) {
	r.HandleFunc("/ping", v1.Ping).Methods("GET")
}
