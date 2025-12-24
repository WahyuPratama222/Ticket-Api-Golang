package models

import "time"

type Ticket struct {
	ID         int       `json:"id" db:"id_ticket"`          // Primary key
	BookingID  int       `json:"booking_id" db:"booking_id"` // Foreign Key To ID Booking
	HolderName string    `json:"holder_name" db:"holder_name"`
	TicketCode string    `json:"ticket_code" db:"ticket_code"`
	Status     string    `json:"status" db:"status"` // unused / used
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
