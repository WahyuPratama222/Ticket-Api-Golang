package validations

import (
	"errors"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
)

// TicketValidator handles ticket input validation
type TicketValidator struct{}

// NewTicketValidator creates a new ticket validator
func NewTicketValidator() *TicketValidator {
	return &TicketValidator{}
}

// ValidateTicketUsage validates if a ticket can be used
func (v *TicketValidator) ValidateTicketUsage(ticket *models.Ticket) error {
	if ticket.Status == "used" {
		return errors.New("ticket has already been used")
	}

	if ticket.Status != "unused" {
		return errors.New("invalid ticket status")
	}

	return nil
}