package v1

import (
	"go-recipe/db"

	"github.com/gorilla/mux"
)

func RegisterV1Routes(r *mux.Router, dataStore *db.DataStore) {
	RegisterHealthRoutes(r)

	v1Router := r.PathPrefix("/api/v1").Subrouter()
	RegisterUserRoutes(v1Router, dataStore)
	RegisterCompanyRoutes(v1Router, dataStore)
}
