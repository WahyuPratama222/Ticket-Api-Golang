package repositories

import (
	"database/sql"
	"errors"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
)

// EventRepository handles database operations for events
type EventRepository struct{}

// NewEventRepository creates a new event repository
func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

// Create inserts a new event into database
func (r *EventRepository) Create(event *models.Event) error {
	query := `INSERT INTO event (organizer_id, title, location, capacity, available_seat, price, status, date) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query, event.OrganizerID, event.Title, event.Location,
		event.Capacity, event.AvailableSeat, event.Price, event.Status, event.Date)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	event.ID = int(id)

	return nil
}

// FindAll retrieves all events
func (r *EventRepository) FindAll() ([]models.Event, error) {
	query := `SELECT id_event, organizer_id, title, location, capacity, available_seat, price, status, date 
	          FROM event ORDER BY id_event ASC`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Event{}
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.OrganizerID, &event.Title, &event.Location,
			&event.Capacity, &event.AvailableSeat, &event.Price, &event.Status, &event.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// FindByID retrieves an event by ID
func (r *EventRepository) FindByID(id int) (models.Event, error) {
	var event models.Event
	query := `SELECT id_event, organizer_id, title, location, capacity, available_seat, price, status, date 
	          FROM event WHERE id_event = ?`

	row := db.DB.QueryRow(query, id)
	err := row.Scan(&event.ID, &event.OrganizerID, &event.Title, &event.Location,
		&event.Capacity, &event.AvailableSeat, &event.Price, &event.Status, &event.Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return event, errors.New("event not found")
		}
		return event, err
	}
	return event, nil
}

// Update updates event information
func (r *EventRepository) Update(id int, event *models.Event) error {
	query := `UPDATE event SET title=?, location=?, capacity=?, available_seat=?, price=?, status=?, date=? 
	          WHERE id_event=?`

	_, err := db.DB.Exec(query, event.Title, event.Location, event.Capacity,
		event.AvailableSeat, event.Price, event.Status, event.Date, id)

	return err
}

// Delete removes an event from database
func (r *EventRepository) Delete(id int) error {
	query := `DELETE FROM event WHERE id_event=?`
	_, err := db.DB.Exec(query, id)
	return err
}

// GetOrganizerRole gets the role of a user by ID
func (r *EventRepository) GetOrganizerRole(organizerID int) (string, error) {
	var role string
	err := db.DB.QueryRow(`SELECT role FROM user WHERE id_user = ?`, organizerID).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("organizer not found")
		}
		return "", err
	}
	return role, nil
}

// CountSuccessfulBookings counts successful bookings for an event
func (r *EventRepository) CountSuccessfulBookings(eventID int) (int, error) {
	var count int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM booking WHERE event_id = ? AND status = 'success'`, eventID).Scan(&count)
	return count, err
}
