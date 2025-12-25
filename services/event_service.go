package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
)

// CreateEvent hanya bisa dipanggil oleh user role organizer
func CreateEvent(event *models.Event) error {
	if event.OrganizerID == 0 || event.Title == "" || event.Location == "" || event.Capacity <= 0 || event.Price < 0 {
		return errors.New("all fields are required and must be valid")
	}

	if event.Date.IsZero() || event.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}

	// Check if organizer exists and has organizer role
	var role string
	err := db.DB.QueryRow(`SELECT role FROM user WHERE id_user = ?`, event.OrganizerID).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("organizer not found")
		}
		return err
	}

	if role != "organizer" {
		return errors.New("only users with organizer role can create events")
	}

	event.AvailableSeat = event.Capacity
	event.Status = "available"
	query := `INSERT INTO event (organizer_id, title, location, capacity, available_seat, price, status, date) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query, event.OrganizerID, event.Title, event.Location, event.Capacity, event.AvailableSeat, event.Price, event.Status, event.Date)
	if err != nil {
		return err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	event.ID = int(id)

	return nil
}

// GetEventByID
func GetEventByID(id int) (models.Event, error) {
	var event models.Event
	query := `SELECT id_event, organizer_id, title, location, capacity, available_seat, price, status, date FROM event WHERE id_event = ?`
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&event.ID, &event.OrganizerID, &event.Title, &event.Location, &event.Capacity, &event.AvailableSeat, &event.Price, &event.Status, &event.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return event, errors.New("event not found")
		}
		return event, err
	}
	return event, nil
}

// UpdateEvent
func UpdateEvent(id int, updated models.Event) error {
	if updated.Title == "" || updated.Location == "" || updated.Price < 0 {
		return errors.New("title, location, and price are required")
	}

	if !updated.Date.IsZero() && updated.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}

	event, err := GetEventByID(id)
	if err != nil {
		return err
	}

	// Check if the organizer exists and has organizer role
	var role string
	err = db.DB.QueryRow(`SELECT role FROM user WHERE id_user = ?`, event.OrganizerID).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("organizer not found")
		}
		return err
	}

	if role != "organizer" {
		return errors.New("only users with organizer role can update events")
	}

	// Update fields
	event.Title = updated.Title
	event.Location = updated.Location
	event.Price = updated.Price
	event.Date = updated.Date

	// Only update capacity if provided and valid
	if updated.Capacity > 0 {
		// Calculate the difference to adjust available seats proportionally
		bookedSeats := event.Capacity - event.AvailableSeat
		if updated.Capacity < bookedSeats {
			return errors.New("cannot reduce capacity below already booked seats")
		}
		event.Capacity = updated.Capacity
		event.AvailableSeat = updated.Capacity - bookedSeats
	}

	// Update status if provided
	if updated.Status != "" {
		if updated.Status != "available" && updated.Status != "unavailable" {
			return errors.New("invalid status: must be 'available' or 'unavailable'")
		}
		event.Status = updated.Status
	}

	// Auto-set status based on available seats
	if event.AvailableSeat == 0 {
		event.Status = "unavailable"
	}

	query := `UPDATE event SET title=?, location=?, capacity=?, available_seat=?, price=?, status=?, date=? WHERE id_event=?`
	_, err = db.DB.Exec(query, event.Title, event.Location, event.Capacity, event.AvailableSeat, event.Price, event.Status, event.Date, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEvent
func DeleteEvent(id int) error {
	event, err := GetEventByID(id)
	if err != nil {
		return err
	}

	// Check if the organizer exists and has organizer role
	var role string
	err = db.DB.QueryRow(`SELECT role FROM user WHERE id_user = ?`, event.OrganizerID).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("organizer not found")
		}
		return err
	}

	if role != "organizer" {
		return errors.New("only users with organizer role can delete events")
	}

	// Check if there are any bookings for this event
	var bookingCount int
	err = db.DB.QueryRow(`SELECT COUNT(*) FROM booking WHERE event_id = ? AND status = 'success'`, id).Scan(&bookingCount)
	if err != nil {
		return err
	}

	if bookingCount > 0 {
		return errors.New("cannot delete event with existing bookings")
	}

	query := `DELETE FROM event WHERE id_event=?`
	_, err = db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}