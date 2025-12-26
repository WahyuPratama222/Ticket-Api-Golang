package service

import (
	"database/sql"
	"fmt"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/repositories"
	"github.com/WahyuPratama222/Ticket-Api-Golang/validations"
	"github.com/google/uuid"
)

type BookingService struct {
	repo      *repositories.BookingRepository
	validator *validations.BookingValidator
}

func NewBookingService() *BookingService {
	return &BookingService{
		repo:      repositories.NewBookingRepository(),
		validator: validations.NewBookingValidator(),
	}
}

// CreateBooking creates a new booking with tickets
func (s *BookingService) CreateBooking(booking *models.Booking) error {
	// Validate basic input
	if err := s.validator.ValidateBookingInput(booking); err != nil {
		return err
	}

	// Validate holder names if provided
	if err := s.validator.ValidateHolderNames(booking.HolderNames, booking.Quantity); err != nil {
		return err
	}

	// Begin transaction
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return err
	}

	// Ensure rollback on error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Lock event row and get event data
	event, err := s.repo.GetEventWithLock(tx, booking.EventID)
	if err != nil {
		return err
	}

	// Validate event availability
	if err := s.validator.ValidateEventAvailability(event); err != nil {
		return err
	}

	// Validate seat availability
	if err := s.validator.ValidateSeatAvailability(event.AvailableSeat, booking.Quantity); err != nil {
		return err
	}

	// Update event seats
	newAvailableSeat := event.AvailableSeat - booking.Quantity
	newStatus := "available"
	if newAvailableSeat == 0 {
		newStatus = "unavailable"
	}

	if err := s.repo.UpdateEventSeats(tx, event.ID, newAvailableSeat, newStatus); err != nil {
		return err
	}

	// Calculate total price
	booking.TotalPrice = event.Price * booking.Quantity

	// Create booking with pending status
	booking.Status = "pending"
	// created_at & updated_at will be set by MySQL

	if err := s.repo.CreateBooking(tx, booking); err != nil {
		return err
	}

	// Generate and create tickets
	if err := s.generateTickets(tx, booking); err != nil {
		// Mark booking as failed before rollback
		s.repo.UpdateBookingStatus(tx, booking.ID, "failed")
		return fmt.Errorf("failed to create tickets: %v", err)
	}

	// Update booking status to success
	if err := s.repo.UpdateBookingStatus(tx, booking.ID, "success"); err != nil {
		return err
	}
	booking.Status = "success"

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// generateTickets generates tickets for a booking within transaction
func (s *BookingService) generateTickets(tx *sql.Tx, booking *models.Booking) error {
	for i := 0; i < booking.Quantity; i++ {
		// Generate unique ticket code
		ticketCode := uuid.New().String()[:8]

		// Determine holder name
		var holderName string
		if i < len(booking.HolderNames) && booking.HolderNames[i] != "" {
			holderName = booking.HolderNames[i]
		} else {
			holderName = fmt.Sprintf("Ticket %d", i+1)
		}

		// Create ticket
		ticket := models.Ticket{
			BookingID:  booking.ID,
			HolderName: holderName,
			TicketCode: ticketCode,
			Status:     "unused",
			// created_at & updated_at will be set by MySQL
		}
		if err := s.repo.CreateTicket(tx, &ticket); err != nil {
			return err
		}
	}

	return nil
}

// GetAllBookings retrieves all bookings
func (s *BookingService) GetAllBookings() ([]models.Booking, error) {
	return s.repo.FindAll()
}

// GetBookingByID retrieves booking with tickets
func (s *BookingService) GetBookingByID(id int) (models.Booking, []models.Ticket, error) {
	// Get booking
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return booking, nil, err
	}
	// Get tickets
	tickets, err := s.repo.FindTicketsByBookingID(id)
	if err != nil {
		return booking, nil, err
	}

	return booking, tickets, nil
}