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
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := services.CreateUser(user); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	user.Password = ""
	utils.WriteJSON(w, http.StatusCreated, user)
}

// GetUserHandler ambil user berdasarkan ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		return
	}

	user, err := services.GetUserByID(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	user.Password = ""
	utils.WriteJSON(w, http.StatusOK, user)
}

// UpdateUserHandler update user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		return
	}

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := services.UpdateUser(id, updated); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	updated.Password = ""
	utils.WriteJSON(w, http.StatusOK, updated)
}

// DeleteUserHandler hapus user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		return
	}

	if err := services.DeleteUser(id); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
