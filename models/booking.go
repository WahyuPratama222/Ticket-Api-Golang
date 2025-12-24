package models

import "time"

type Booking struct {
	ID          int       `json:"id" db:"id_booking"`           // Primary Key
	CustomerID  int       `json:"customer_id" db:"customer_id"` // Foreign Key To ID User Role Customer
	EventID     int       `json:"event_id" db:"event_id"`       // Foreign Key To ID Event
	TotalPrice  int       `json:"total_price" db:"total_price"`
	Quantity    int       `json:"quantity" db:"quantity"`
	Status      string    `json:"status" db:"status"` // success / pending / failed
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	HolderNames []string  `json:"holder_names,omitempty"` // array input untuk Holder Names di Ticket (row tidak masuk ke tabel booking)
}
