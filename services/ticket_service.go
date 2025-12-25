package service

import (
	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/repositories"
)

// TicketService handles ticket business logic
type TicketService struct {
	repo *repositories.TicketRepository
}

// NewTicketService creates a new ticket service
func NewTicketService() *TicketService {
	return &TicketService{
		repo: repositories.NewTicketRepository(),
	}
}

// GetAllTickets retrieves all tickets
func (s *TicketService) GetAllTickets() ([]models.Ticket, error) {
	return s.repo.FindAll()
}

// GetTicketByID retrieves ticket by ID
func (s *TicketService) GetTicketByID(id int) (models.Ticket, error) {
	return s.repo.FindByID(id)
}