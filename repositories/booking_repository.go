package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
)

// BookingRepository handles database operations for bookings
type BookingRepository struct{}

// NewBookingRepository creates a new booking repository
func NewBookingRepository() *BookingRepository {
	return &BookingRepository{}
}

// BeginTransaction starts a database transaction
func (r *BookingRepository) BeginTransaction() (*sql.Tx, error) {
	return db.DB.Begin()
}

// GetEventWithLock retrieves event with row-level lock (FOR UPDATE)
func (r *BookingRepository) GetEventWithLock(tx *sql.Tx, eventID int) (models.Event, error) {
	var event models.Event
	query := `SELECT id_event, available_seat, price, status FROM event WHERE id_event=? FOR UPDATE`

	row := tx.QueryRow(query, eventID)
	err := row.Scan(&event.ID, &event.AvailableSeat, &event.Price, &event.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return event, errors.New("event not found")
		}
		return event, err
	}

	return event, nil
}

// UpdateEventSeats updates available seats and status for an event
func (r *BookingRepository) UpdateEventSeats(tx *sql.Tx, eventID int, newAvailableSeat int, newStatus string) error {
	query := `UPDATE event SET available_seat=?, status=? WHERE id_event=?`
	_, err := tx.Exec(query, newAvailableSeat, newStatus, eventID)
	return err
}

// CreateBooking inserts a new booking into database within transaction
func (r *BookingRepository) CreateBooking(tx *sql.Tx, booking *models.Booking) error {
	query := `INSERT INTO booking (customer_id, event_id, quantity, total_price, status, created_at, updated_at) 
	          VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := tx.Exec(query,
		booking.CustomerID,
		booking.EventID,
		booking.Quantity,
		booking.TotalPrice,
		booking.Status,
		booking.CreatedAt,
		booking.UpdatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	booking.ID = int(id)

	return nil
}

// UpdateBookingStatus updates booking status within transaction
func (r *BookingRepository) UpdateBookingStatus(tx *sql.Tx, bookingID int, status string) error {
	query := `UPDATE booking SET status=?, updated_at=? WHERE id_booking=?`
	_, err := tx.Exec(query, status, time.Now(), bookingID)
	return err
}

// CreateTicket inserts a new ticket into database within transaction
func (r *BookingRepository) CreateTicket(tx *sql.Tx, ticket *models.Ticket) error {
	query := `INSERT INTO ticket (booking_id, holder_name, ticket_code, status, created_at, updated_at) 
	          VALUES (?, ?, ?, ?, ?, ?)`

	_, err := tx.Exec(query,
		ticket.BookingID,
		ticket.HolderName,
		ticket.TicketCode,
		ticket.Status,
		ticket.CreatedAt,
		ticket.UpdatedAt,
	)

	return err
}

// FindAll retrieves all bookings
func (r *BookingRepository) FindAll() ([]models.Booking, error) {
	query := `SELECT id_booking, customer_id, event_id, quantity, total_price, status, created_at, updated_at 
	          FROM booking ORDER BY id_booking ASC`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := []models.Booking{}
	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(
			&booking.ID,
			&booking.CustomerID,
			&booking.EventID,
			&booking.Quantity,
			&booking.TotalPrice,
			&booking.Status,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// FindByID retrieves a booking by ID
func (r *BookingRepository) FindByID(id int) (models.Booking, error) {
	var booking models.Booking
	query := `SELECT id_booking, customer_id, event_id, quantity, total_price, status, created_at, updated_at 
	          FROM booking WHERE id_booking=?`

	row := db.DB.QueryRow(query, id)
	err := row.Scan(
		&booking.ID,
		&booking.CustomerID,
		&booking.EventID,
		&booking.Quantity,
		&booking.TotalPrice,
		&booking.Status,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return booking, errors.New("booking not found")
		}
		return booking, err
	}

	return booking, nil
}

// FindTicketsByBookingID retrieves all tickets for a booking
func (r *BookingRepository) FindTicketsByBookingID(bookingID int) ([]models.Ticket, error) {
	query := `SELECT id_ticket, booking_id, holder_name, ticket_code, status, created_at, updated_at 
	          FROM ticket WHERE booking_id=?`

	rows, err := db.DB.Query(query, bookingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := []models.Ticket{}
	for rows.Next() {
		var t models.Ticket
		err := rows.Scan(
			&t.ID,
			&t.BookingID,
			&t.HolderName,
			&t.TicketCode,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}
