package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"github.com/google/uuid"
)

func CreateBooking(booking *models.Booking) error {
	if booking.CustomerID == 0 || booking.EventID == 0 || booking.Quantity <= 0 {
		return errors.New("all fields are required")
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	
	// Simplified defer - just rollback if needed
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Lock event
	var event models.Event
	row := tx.QueryRow(`SELECT id_event, available_seat, price, status FROM event WHERE id_event=? FOR UPDATE`, booking.EventID)
	err = row.Scan(&event.ID, &event.AvailableSeat, &event.Price, &event.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("event not found")
		}
		return err
	}

	if event.Status != "available" {
		return errors.New("event is not available")
	}

	if event.AvailableSeat < booking.Quantity {
		return errors.New("not enough seats available")
	}

	// Update available seat & event status
	newAvailable := event.AvailableSeat - booking.Quantity
	newStatus := event.Status
	if newAvailable == 0 {
		newStatus = "unavailable"
	}
	_, err = tx.Exec(`UPDATE event SET available_seat=?, status=? WHERE id_event=?`, newAvailable, newStatus, event.ID)
	if err != nil {
		return err
	}

	// Hitung total price
	booking.TotalPrice = event.Price * booking.Quantity

	// Insert booking (pending dulu)
	res, err := tx.Exec(`INSERT INTO booking (customer_id, event_id, quantity, total_price, status, created_at, updated_at) 
	                     VALUES (?, ?, ?, ?, ?, ?, ?)`,
		booking.CustomerID, booking.EventID, booking.Quantity, booking.TotalPrice, "pending", time.Now(), time.Now())
	if err != nil {
		return err
	}

	bookingID, _ := res.LastInsertId()
	booking.ID = int(bookingID)
	booking.Status = "pending"
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	// Generate tickets
	for i := 0; i < booking.Quantity; i++ {
		ticketCode := uuid.New().String()[:8]

		var holder string
		if i < len(booking.HolderNames) && booking.HolderNames[i] != "" {
			holder = booking.HolderNames[i]
		} else {
			holder = fmt.Sprintf("Ticket %d", i+1)
		}

		_, err := tx.Exec(`INSERT INTO ticket (booking_id, holder_name, ticket_code, status, created_at, updated_at) 
                       VALUES (?, ?, ?, ?, ?, ?)`,
			bookingID, holder, ticketCode, "unused", time.Now(), time.Now())
		if err != nil {
			// Update booking status to failed before rollback
			tx.Exec("UPDATE booking SET status='failed', updated_at=? WHERE id_booking=?", time.Now(), booking.ID)
			return fmt.Errorf("failed to create ticket: %v", err)
		}
	}

	// Update status booking jadi success sebelum commit
	_, err = tx.Exec("UPDATE booking SET status='success', updated_at=? WHERE id_booking=?", time.Now(), booking.ID)
	if err != nil {
		return err
	}

	booking.Status = "success"

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetBookingByID ambil booking + tiket
func GetBookingByID(id int) (models.Booking, []models.Ticket, error) {
	var booking models.Booking
	row := db.DB.QueryRow(`SELECT id_booking, customer_id, event_id, quantity, total_price, status, created_at, updated_at 
	                       FROM booking WHERE id_booking=?`, id)
	err := row.Scan(&booking.ID, &booking.CustomerID, &booking.EventID, &booking.Quantity, &booking.TotalPrice,
		&booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return booking, nil, errors.New("booking not found")
		}
		return booking, nil, err
	}

	// Ambil tickets
	rows, err := db.DB.Query(`SELECT id_ticket, booking_id, holder_name, ticket_code, status, created_at, updated_at 
	                          FROM ticket WHERE booking_id=?`, id)
	if err != nil {
		return booking, nil, err
	}
	defer rows.Close()

	tickets := []models.Ticket{}
	for rows.Next() {
		var t models.Ticket
		err := rows.Scan(&t.ID, &t.BookingID, &t.HolderName, &t.TicketCode, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return booking, nil, err
		}
		tickets = append(tickets, t)
	}

	return booking, tickets, nil
}