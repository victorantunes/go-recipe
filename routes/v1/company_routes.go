package v1

import (
	"go-recipe/db"
	"go-recipe/handlers/v1"
	"go-recipe/services"

	"github.com/gorilla/mux"
)

func RegisterCompanyRoutes(r *mux.Router, store *db.DataStore) {
	companyService := services.NewCompanyService(store.CompanyRepository)
	companyHandler := v1.NewCompanyHandler(companyService)

	r.HandleFunc("/companies/{id}", companyHandler.GetCompany).Methods("GET")
	r.HandleFunc("/companies", companyHandler.CreateCompany).Methods("POST")
	r.HandleFunc("/companies", companyHandler.ListCompanies).Methods("GET")
	r.HandleFunc("/companies/{id}", companyHandler.DeleteCompany).Methods("DELETE")
}
