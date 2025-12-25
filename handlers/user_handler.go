package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/services"
	"github.com/WahyuPratama222/Ticket-Api-Golang/utils"
	"github.com/gorilla/mux"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		service: service.NewUserService(),
	}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Don't return password in response
	response := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	// Don't return password
	response := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// UpdateUser updates user information
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.UpdateUser(id, updated); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get updated user data
	user, _ := h.service.GetUserByID(id)
	response := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// DeleteUser removes a user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "user deleted successfully"})
}