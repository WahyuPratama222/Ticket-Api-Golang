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

	response := map[string]any{
		"id":             event.ID,
		"organizer_id":   event.OrganizerID,
		"title":          event.Title,
		"location":       event.Location,
		"capacity":       event.Capacity,
		"available_seat": event.AvailableSeat,
		"price":          event.Price,
		"status":         event.Status,
		"date":           event.Date,
	}

	utils.WriteSuccessJSON(w, http.StatusCreated, "event created successfully", response)
}

// GetAllEvents retrieves all events
func (h *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetAllEvents()
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteSuccessJSON(w, http.StatusOK, "events retrieved successfully", events)
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

	utils.WriteSuccessJSON(w, http.StatusOK, "event retrieved successfully", event)
}

// UpdateEvent updates event information (without organizer_id)
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

	// Remove organizer_id from update (should not be changed)
	updated.OrganizerID = 0

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
	utils.WriteSuccessJSON(w, http.StatusOK, "event updated successfully", event)
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

	utils.WriteSuccessJSON(w, http.StatusOK, "event deleted successfully", nil)
}