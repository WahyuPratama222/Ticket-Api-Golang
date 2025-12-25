package service

import (
	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/repositories"
	"github.com/WahyuPratama222/Ticket-Api-Golang/validations"
)

// TicketService handles ticket business logic
type TicketService struct {
	repo      *repositories.TicketRepository
	validator *validations.TicketValidator
}

// NewTicketService creates a new ticket service
func NewTicketService() *TicketService {
	return &TicketService{
		repo:      repositories.NewTicketRepository(),
		validator: validations.NewTicketValidator(),
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

// UseTicket marks a ticket as used
func (s *TicketService) UseTicket(id int) error {
	// Get existing ticket
	ticket, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Validate ticket can be used
	if err := s.validator.ValidateTicketUsage(&ticket); err != nil {
		return err
	}

	// Update status to 'used'
	return s.repo.UpdateStatus(id, "used")
}
