package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	services "github.com/WahyuPratama222/Ticket-Api-Golang/services"
	"github.com/WahyuPratama222/Ticket-Api-Golang/utils"
	"github.com/gorilla/mux"
)

// RegisterUserHandler buat user baru
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := services.CreateUser(&user); err != nil {
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

// GetUserHandler ambil user berdasarkan ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := services.GetUserByID(id)
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

// UpdateUserHandler update user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := services.UpdateUser(id, updated); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get updated user data
	user, _ := services.GetUserByID(id)
	response := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// DeleteUserHandler hapus user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := services.DeleteUser(id); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "user deleted successfully"})
}