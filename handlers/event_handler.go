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

// EventHandler handles HTTP requests for event operations
type EventHandler struct {
	service *service.EventService
}

// NewEventHandler creates a new event handler
func NewEventHandler() *EventHandler {
	return &EventHandler{
		service: service.NewEventService(),
	}
}

// CreateEvent handles event creation
func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.CreateEvent(&event); err != nil {
		// Check specific error messages for appropriate status codes
		if err.Error() == "organizer not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "only users with organizer role can manage events" {
			utils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, event)
}

// GetEvent retrieves an event by ID
func (h *EventHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid event id")
		return
	}

	event, err := h.service.GetEventByID(id)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, event)
}

// UpdateEvent updates event information
func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.UpdateEvent(id, updated); err != nil {
		// Check specific error messages for appropriate status codes
		if err.Error() == "event not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "organizer not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "only users with organizer role can manage events" {
			utils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get updated event
	event, _ := h.service.GetEventByID(id)
	utils.WriteJSON(w, http.StatusOK, event)
}

// DeleteEvent removes an event
func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid event id")
		return
	}

	if err := h.service.DeleteEvent(id); err != nil {
		// Check specific error messages for appropriate status codes
		if err.Error() == "event not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "organizer not found" {
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "only users with organizer role can manage events" {
			utils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "event deleted successfully"})
}