package handlers

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
	service *services.BookingService
}

// NewBookingHandler creates a new booking handler
func NewBookingHandler() *BookingHandler {
	return &BookingHandler{
		service: services.NewBookingService(),
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
        // Check specific error messages for appropriate status codes
        if err.Error() == "event not found" {
            utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
            return
        }
        if err.Error() == "event is not available" {
            utils.WriteErrorJSON(w, http.StatusConflict, err.Error())
            return
        }
        if err.Error() == "not enough seats available" {
            utils.WriteErrorJSON(w, http.StatusConflict, err.Error())
            return
        }
        utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
        return
    }

    response := map[string]any{
        "id":          booking.ID,
        "customer_id": booking.CustomerID,
        "event_id":    booking.EventID,
        "quantity":    booking.Quantity,
        "total_price": booking.TotalPrice,
        "status":      booking.Status,
        "created_at":  booking.CreatedAt,
        "updated_at":  booking.UpdatedAt,
    }

    utils.WriteSuccessJSON(w, http.StatusCreated, "booking created successfully", response)
}

// GetAllBookings retrieves all bookings
func (h *BookingHandler) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.service.GetAllBookings()
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteSuccessJSON(w, http.StatusOK, "bookings retrieved successfully", bookings)
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

	response := map[string]any{
		"booking": booking,
		"tickets": tickets,
	}

	utils.WriteSuccessJSON(w, http.StatusOK, "booking retrieved successfully", response)
}