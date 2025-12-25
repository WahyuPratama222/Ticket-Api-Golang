package validations

import (
	"errors"
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
)

// EventValidator handles event input validation
type EventValidator struct{}

// NewEventValidator creates a new event validator
func NewEventValidator() *EventValidator {
	return &EventValidator{}
}

// ValidateCreate validates event data for creation
func (v *EventValidator) ValidateCreate(event *models.Event) error {
	if event.OrganizerID == 0 {
		return errors.New("organizer_id is required")
	}

	if event.Title == "" {
		return errors.New("title is required")
	}

	if event.Location == "" {
		return errors.New("location is required")
	}

	if event.Capacity <= 0 {
		return errors.New("capacity must be greater than 0")
	}

	if event.Price < 0 {
		return errors.New("price cannot be negative")
	}

	if event.Date.IsZero() {
		return errors.New("date is required")
	}

	if event.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}

	return nil
}

// ValidateUpdate validates event data for update
func (v *EventValidator) ValidateUpdate(updated *models.Event, existing *models.Event) error {
	if updated.Title == "" {
		return errors.New("title is required")
	}

	if updated.Location == "" {
		return errors.New("location is required")
	}

	if updated.Price < 0 {
		return errors.New("price cannot be negative")
	}

	if !updated.Date.IsZero() && updated.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}

	return nil
}

// ValidateCapacityUpdate validates capacity changes
func (v *EventValidator) ValidateCapacityUpdate(newCapacity int, bookedSeats int) error {
	if newCapacity <= 0 {
		return errors.New("capacity must be greater than 0")
	}

	if newCapacity < bookedSeats {
		return errors.New("cannot reduce capacity below already booked seats")
	}

	return nil
}

// ValidateStatus validates event status
func (v *EventValidator) ValidateStatus(status string) error {
	if status != "" && status != "available" && status != "unavailable" {
		return errors.New("invalid status: must be 'available' or 'unavailable'")
	}
	return nil
}

// ValidateOrganizerRole validates if user has organizer role
func (v *EventValidator) ValidateOrganizerRole(role string) error {
	if role != "organizer" {
		return errors.New("only users with organizer role can manage events")
	}
	return nil
}