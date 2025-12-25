package validations

import (
	"errors"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
)

// BookingValidator handles booking input validation
type BookingValidator struct{}

// NewBookingValidator creates a new booking validator
func NewBookingValidator() *BookingValidator {
	return &BookingValidator{}
}

// ValidateBookingInput validates basic booking input fields
func (v *BookingValidator) ValidateBookingInput(booking *models.Booking) error {
	if booking.CustomerID == 0 {
		return errors.New("customer_id is required")
	}

	if booking.EventID == 0 {
		return errors.New("event_id is required")
	}

	if booking.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return nil
}

// ValidateEventAvailability checks if event is available for booking
func (v *BookingValidator) ValidateEventAvailability(event models.Event) error {
	if event.Status != "available" {
		return errors.New("event is not available")
	}
	return nil
}

// ValidateSeatAvailability checks if enough seats are available
func (v *BookingValidator) ValidateSeatAvailability(availableSeats int, requestedQuantity int) error {
	if availableSeats < requestedQuantity {
		return errors.New("not enough seats available")
	}
	return nil
}

// ValidateHolderNames validates holder names array
func (v *BookingValidator) ValidateHolderNames(holderNames []string, quantity int) error {
	// Holder names are optional, but if provided, should match quantity
	if len(holderNames) > 0 && len(holderNames) != quantity {
		return errors.New("number of holder names must match quantity")
	}
	return nil
}