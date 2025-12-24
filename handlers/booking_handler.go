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
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := services.CreateBooking(booking); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, booking)
}

// GetBookingHandler ambil detail booking + tiket
func GetBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid booking id"})
		return
	}

	booking, tickets, err := services.GetBookingByID(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	resp := struct {
		Booking models.Booking  `json:"booking"`
		Tickets []models.Ticket `json:"tickets"`
	}{
		Booking: booking,
		Tickets: tickets,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
