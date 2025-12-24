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

// CreateBookingHandler membuat booking baru
func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := services.CreateBooking(&booking); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	response := map[string]interface{}{
		"message":    "booking created successfully",
		"booking_id": booking.ID,
		"status":     booking.Status,
		"total_price": booking.TotalPrice,
		"created_at": booking.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

// GetBookingHandler ambil detail booking + tiket
func GetBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid booking id")
		return
	}

	booking, tickets, err := services.GetBookingByID(id)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	resp := map[string]interface{}{
		"booking": booking,
		"tickets": tickets,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}