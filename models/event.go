package models

import "time"

type Event struct {
	ID            int       `json:"id" db:"id_event"`               // Primary Key
	OrganizerID   int       `json:"organizer_id" db:"organizer_id"` // Foreign Key To ID User Role Organizer
	Title         string    `json:"title" db:"title"`
	Location      string    `json:"location" db:"location"`
	Capacity      int       `json:"capacity" db:"capacity"`
	AvailableSeat int       `json:"available_seat" db:"available_seat"`
	Price         int       `json:"price" db:"price"`
	Status        string    `json:"status" db:"status"` // available / unavailable
	Date          time.Time `json:"date" db:"date"`
}
