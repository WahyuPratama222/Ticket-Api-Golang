package service

import (
	"database/sql"
	"errors"
	"regexp"
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	mysql "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword checks if password meets minimum requirements
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

// CreateUser membuat user baru
func CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("all fields are required")
	}

	if !ValidateEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if err := ValidatePassword(user.Password); err != nil {
		return err
	}

	if user.Role != "customer" && user.Role != "organizer" {
		return errors.New("role must be either 'customer' or 'organizer'")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO user (name, password, email, role, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(query, user.Name, string(hashedPassword), user.Email, user.Role, time.Now())
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok && driverErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	user.CreatedAt = time.Now()

	return nil
}

// GetUserByID
func GetUserByID(id int) (models.User, error) {
	var user models.User
	query := `SELECT id_user, name, email, role, created_at FROM user WHERE id_user = ?`
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

// UpdateUser update user (name, email, password, role)
func UpdateUser(id int, updated models.User) error {
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}

	// Preserve existing values if not provided
	if updated.Name == "" {
		updated.Name = user.Name
	}
	if updated.Email == "" {
		updated.Email = user.Email
	} else {
		if !ValidateEmail(updated.Email) {
			return errors.New("invalid email format")
		}
	}
	if updated.Role == "" {
		updated.Role = user.Role
	} else {
		if updated.Role != "customer" && updated.Role != "organizer" {
			return errors.New("role must be either 'customer' or 'organizer'")
		}
	}

	var hashedPassword string
	if updated.Password != "" {
		if err := ValidatePassword(updated.Password); err != nil {
			return err
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(updated.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		hashedPassword = string(hash)
	} else {
		var pw string
		err := db.DB.QueryRow("SELECT password FROM user WHERE id_user = ?", id).Scan(&pw)
		if err != nil {
			return err
		}
		hashedPassword = pw
	}

	query := `UPDATE user SET name = ?, email = ?, password = ?, role = ? WHERE id_user = ?`
	_, err = db.DB.Exec(query, updated.Name, updated.Email, hashedPassword, updated.Role, id)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok && driverErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}

	return nil
}

// DeleteUser hapus user berdasarkan id
func DeleteUser(id int) error {
	_, err := GetUserByID(id)
	if err != nil {
		return err
	}

	// Check if user has any bookings
	var bookingCount int
	err = db.DB.QueryRow(`SELECT COUNT(*) FROM booking WHERE customer_id = ?`, id).Scan(&bookingCount)
	if err != nil {
		return err
	}

	if bookingCount > 0 {
		return errors.New("cannot delete user with existing bookings")
	}

	// Check if user (organizer) has any events
	var eventCount int
	err = db.DB.QueryRow(`SELECT COUNT(*) FROM event WHERE organizer_id = ?`, id).Scan(&eventCount)
	if err != nil {
		return err
	}

	if eventCount > 0 {
		return errors.New("cannot delete organizer with existing events")
	}

	query := `DELETE FROM user WHERE id_user = ?`
	_, err = db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}