package v1

import (
	v1 "cql-backend/handlers/v1"
	"github.com/gorilla/mux"
)

func RegisterV1Routes(r *mux.Router) {
	v1Router := r.PathPrefix("/api/v1").Subrouter()
	v1Router.HandleFunc("/ping", v1.Ping).Methods("GET")
}
