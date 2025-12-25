package service

import (
	"errors"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/repositories"
	"github.com/WahyuPratama222/Ticket-Api-Golang/validations"
)

// EventService handles event business logic
type EventService struct {
	repo      *repositories.EventRepository
	validator *validations.EventValidator
}

// NewEventService creates a new event service
func NewEventService() *EventService {
	return &EventService{
		repo:      repositories.NewEventRepository(),
		validator: validations.NewEventValidator(),
	}
}

// CreateEvent creates a new event
func (s *EventService) CreateEvent(event *models.Event) error {
	// Validate input
	if err := s.validator.ValidateCreate(event); err != nil {
		return err
	}

	// Check if organizer exists and has organizer role
	role, err := s.repo.GetOrganizerRole(event.OrganizerID)
	if err != nil {
		return err
	}

	if err := s.validator.ValidateOrganizerRole(role); err != nil {
		return err
	}

	// Set initial values
	event.AvailableSeat = event.Capacity
	event.Status = "available"

	// Save to database
	return s.repo.Create(event)
}

// GetAllEvents retrieves all events
func (s *EventService) GetAllEvents() ([]models.Event, error) {
	return s.repo.FindAll()
}

// GetEventByID retrieves event by ID
func (s *EventService) GetEventByID(id int) (models.Event, error) {
	return s.repo.FindByID(id)
}

// UpdateEvent updates event information (without organizer_id)
func (s *EventService) UpdateEvent(id int, updated models.Event) error {
	// Get existing event
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Check if organizer exists and has organizer role
	role, err := s.repo.GetOrganizerRole(existing.OrganizerID)
	if err != nil {
		return err
	}

	if err := s.validator.ValidateOrganizerRole(role); err != nil {
		return err
	}

	// Validate update data
	if err := s.validator.ValidateUpdate(&updated, &existing); err != nil {
		return err
	}

	// Preserve existing organizer_id (cannot be changed)
	updated.OrganizerID = existing.OrganizerID

	// Handle capacity update
	if updated.Capacity > 0 {
		bookedSeats := existing.Capacity - existing.AvailableSeat
		
		if err := s.validator.ValidateCapacityUpdate(updated.Capacity, bookedSeats); err != nil {
			return err
		}

		updated.AvailableSeat = updated.Capacity - bookedSeats
	} else {
		// Keep existing capacity and available seats
		updated.Capacity = existing.Capacity
		updated.AvailableSeat = existing.AvailableSeat
	}

	// Handle status update
	if updated.Status != "" {
		if err := s.validator.ValidateStatus(updated.Status); err != nil {
			return err
		}
	} else {
		updated.Status = existing.Status
	}

	// Auto-set status based on available seats
	if updated.AvailableSeat == 0 {
		updated.Status = "unavailable"
	}

	// Keep existing date if not provided
	if updated.Date.IsZero() {
		updated.Date = existing.Date
	}

	// Update in database
	return s.repo.Update(id, &updated)
}

// DeleteEvent deletes an event
func (s *EventService) DeleteEvent(id int) error {
	// Get existing event
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Check if organizer exists and has organizer role
	role, err := s.repo.GetOrganizerRole(existing.OrganizerID)
	if err != nil {
		return err
	}

	if err := s.validator.ValidateOrganizerRole(role); err != nil {
		return err
	}

	// Check if there are any successful bookings
	bookingCount, err := s.repo.CountSuccessfulBookings(id)
	if err != nil {
		return err
	}

	if bookingCount > 0 {
		return errors.New("cannot delete event with existing bookings")
	}

	// Delete event
	return s.repo.Delete(id)
}

