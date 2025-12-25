package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id_user"` // Primary Key
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"` // customer / organizer
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}