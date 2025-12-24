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

// CreateEventHandler buat event baru (hanya organizer)
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	userRole := r.Header.Get("Role") 

	if err := services.CreateEvent(event, userRole); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, event)
}

// GetEventHandler ambil event berdasarkan ID
func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid event id"})
		return
	}

	event, err := services.GetEventByID(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, event)
}

// UpdateEventHandler update event (hanya organizer)
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid event id"})
		return
	}

	var updated models.Event
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	userRole := r.Header.Get("Role") // ambil role dari header

	if err := services.UpdateEvent(id, updated, userRole); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, updated)
}

// DeleteEventHandler hapus event (hanya organizer)
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid event id"})
		return
	}

	userRole := r.Header.Get("Role") // ambil role dari header

	if err := services.DeleteEvent(id, userRole); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
