package v1

import (
	"cql-backend/models"
	"cql-backend/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		returnError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		returnError(w, http.StatusBadRequest, err.Error())
		return
	}

	returnJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		returnError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		returnError(w, http.StatusNotFound, "User not found")
		return
	}

	returnJSON(w, http.StatusOK, user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		returnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	returnJSON(w, http.StatusOK, users)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		returnError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		returnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func returnJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func returnError(w http.ResponseWriter, code int, message string) {
	returnJSON(w, code, map[string]string{"returnError": message})
}
