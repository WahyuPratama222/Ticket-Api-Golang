package service

import (
	"errors"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
)

// CreateEvent hanya bisa dipanggil oleh user role organizer
func CreateEvent(event models.Event, userRole string) error {
	if userRole != "organizer" {
		return errors.New("only organizers can create events")
	}

	if event.OrganizerID == 0 || event.Title == "" || event.Location == "" || event.Capacity <= 0 || event.Price < 0 || event.Date.IsZero() {
		return errors.New("all fields are required and must be valid")
	}

	event.AvailableSeat = event.Capacity
	event.Status = "available"
	query := `INSERT INTO event (organizer_id, title, location, capacity, available_seat, price, status, date) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.DB.Exec(query, event.OrganizerID, event.Title, event.Location, event.Capacity, event.AvailableSeat, event.Price, event.Status, event.Date)
	if err != nil {
		return err
	}

	return nil
}

// GetEventByID
func GetEventByID(id int) (models.Event, error) {
	var event models.Event
	query := `SELECT id_event, organizer_id, title, location, capacity, available_seat, price, status, date FROM event WHERE id_event = ?`
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&event.ID, &event.OrganizerID, &event.Title, &event.Location, &event.Capacity, &event.AvailableSeat, &event.Price, &event.Status, &event.Date)
	if err != nil {
		return event, errors.New("event not found")
	}
	return event, nil
}

// UpdateEvent
func UpdateEvent(id int, updated models.Event, userRole string) error {
	if userRole != "organizer" {
		return errors.New("only organizers can update events")
	}

	if updated.Title == "" || updated.Location == "" || updated.Capacity <= 0 || updated.Price < 0 || updated.Date.IsZero() {
		return errors.New("all fields are required and must be valid")
	}

	event, err := GetEventByID(id)
	if err != nil {
		return err
	}

	event.Title = updated.Title
	event.Location = updated.Location
	event.Capacity = updated.Capacity
	event.AvailableSeat = updated.AvailableSeat
	event.Price = updated.Price
	event.Status = updated.Status
	event.Date = updated.Date

	query := `UPDATE event SET title=?, location=?, capacity=?, available_seat=?, price=?, status=?, date=? WHERE id_event=?`
	_, err = db.DB.Exec(query, event.Title, event.Location, event.Capacity, event.AvailableSeat, event.Price, event.Status, event.Date, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEvent
func DeleteEvent(id int, userRole string) error {
	if userRole != "organizer" {
		return errors.New("only organizers can delete events")
	}

	_, err := GetEventByID(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM event WHERE id_event=?`
	_, err = db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
