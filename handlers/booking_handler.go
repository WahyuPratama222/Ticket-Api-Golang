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

// BookingHandler handles HTTP requests for booking operations
type BookingHandler struct {
	service *service.BookingService
}

// NewBookingHandler creates a new booking handler
func NewBookingHandler() *BookingHandler {
	return &BookingHandler{
		service: service.NewBookingService(),
	}
}

// CreateBooking handles booking creation
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.CreateBooking(&booking); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	response := map[string]interface{}{
		"message":     "booking created successfully",
		"booking_id":  booking.ID,
		"status":      booking.Status,
		"total_price": booking.TotalPrice,
		"created_at":  booking.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

// GetBooking retrieves booking details with tickets
func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid booking id")
		return
	}

	booking, tickets, err := h.service.GetBookingByID(id)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	response := map[string]interface{}{
		"booking": booking,
		"tickets": tickets,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}