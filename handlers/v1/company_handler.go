package v1

import (
	"encoding/json"
	"go-recipe/models"
	"go-recipe/models/dto"
	"go-recipe/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(service *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{service: service}
}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		ReturnError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.service.CreateCompany(&company); err != nil {
		ReturnError(w, http.StatusBadRequest, err.Error())
		return
	}

	ReturnJSON(w, http.StatusCreated, company)
}

func (h *CompanyHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	company, err := h.service.GetCompanyByID(uint(id))
	if err != nil {
		ReturnError(w, http.StatusNotFound, "Company not found")
		return
	}

	log.Println("Fetched company:", company.Users)
	companiesDto := dto.MapCompanyToResponse(*company)
	ReturnJSON(w, http.StatusOK, companiesDto)
}

func (h *CompanyHandler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.service.ListCompanies()
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	companiesDto := dto.MapCompaniesToResponse(companies)
	ReturnJSON(w, http.StatusOK, companiesDto)
}

func (h *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	if err := h.service.DeleteCompany(uint(id)); err != nil {
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
