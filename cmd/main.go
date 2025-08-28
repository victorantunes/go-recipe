package main

import (
	"go-recipe/config"
	"go-recipe/db"
	"go-recipe/models"
	"log"
	"net/http"

	v1routes "go-recipe/routes/v1"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	dbConn := db.InitPostgres()

	if err := dbConn.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	store := db.NewDataStore(dbConn)

	r := mux.NewRouter()
	v1routes.RegisterV1Routes(r, store)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
