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

// CreateEventHandler buat event baru
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := services.CreateEvent(&event); err != nil {
		// Check specific error messages for appropriate status codes
		if err.Error() == "organizer not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "only users with organizer role can create events" {
			utils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
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
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid event id")
		return
	}

	event, err := services.GetEventByID(id)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, event)
}

// UpdateEventHandler update event
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid event id")
		return
	}

	var updated models.Event
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := services.UpdateEvent(id, updated); err != nil {
		// Check specific error messages for appropriate status codes
		if err.Error() == "event not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "organizer not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "only users with organizer role can update events" {
			utils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get updated event
	event, _ := services.GetEventByID(id)
	utils.WriteJSON(w, http.StatusOK, event)
}

// DeleteEventHandler hapus event
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid event id")
		return
	}

	if err := services.DeleteEvent(id); err != nil {
		// Check specific error messages for appropriate status codes
		if err.Error() == "event not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "organizer not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "only users with organizer role can delete events" {
			utils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "event deleted successfully"})
}