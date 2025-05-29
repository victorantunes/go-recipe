package v1

import (
	"cql-backend/db"
	"github.com/gorilla/mux"
)

func RegisterV1Routes(r *mux.Router, store *db.Store) {
	v1Router := r.PathPrefix("/api/v1").Subrouter()

	RegisterHealthRoutes(v1Router)
	RegisterUserRoutes(v1Router, store)
}
