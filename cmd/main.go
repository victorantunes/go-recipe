package main

import (
	"cql-backend/config"
	"cql-backend/db"
	"log"
	"net/http"

	v1routes "cql-backend/routes/v1"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	db.InitPostgres()

	r := mux.NewRouter()
	v1routes.RegisterV1Routes(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
