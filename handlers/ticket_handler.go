package handler

import (
	"net/http"
	"strconv"

	"github.com/WahyuPratama222/Ticket-Api-Golang/services"
	"github.com/WahyuPratama222/Ticket-Api-Golang/utils"
	"github.com/gorilla/mux"
)

// TicketHandler handles HTTP requests for ticket operations
type TicketHandler struct {
	service *service.TicketService
}

// NewTicketHandler creates a new ticket handler
func NewTicketHandler() *TicketHandler {
	return &TicketHandler{
		service: service.NewTicketService(),
	}
}

// GetAllTickets retrieves all tickets
func (h *TicketHandler) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.service.GetAllTickets()
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteSuccessJSON(w, http.StatusOK, "tickets retrieved successfully", tickets)
}

// GetTicket retrieves a ticket by ID
func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "invalid ticket id")
		return
	}

	ticket, err := h.service.GetTicketByID(id)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteSuccessJSON(w, http.StatusOK, "ticket retrieved successfully", ticket)
}