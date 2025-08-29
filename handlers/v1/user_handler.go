package v1

import (
	"encoding/json"
	"go-recipe/models"
	"go-recipe/models/dto"
	"go-recipe/services"
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
		ReturnError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		ReturnError(w, http.StatusBadRequest, err.Error())
		return
	}

	ReturnJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		ReturnError(w, http.StatusNotFound, "User not found")
		return
	}

	userDto := dto.MapUserToResponse(*user)
	ReturnJSON(w, http.StatusOK, userDto)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	companyId := query.Get("company_id")

	users, err := h.service.ListUsers(companyId)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	usersDto := dto.MapUsersToResponse(users)
	ReturnJSON(w, http.StatusOK, usersDto)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
