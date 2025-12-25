package repositories

import (
	"database/sql"
	"errors"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	mysql "github.com/go-sql-driver/mysql"
)

// UserRepository handles database operations for users
type UserRepository struct{}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create inserts a new user into database
func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO user (name, password, email, role, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(query, user.Name, user.Password, user.Email, user.Role, user.CreatedAt)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok && driverErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)

	return nil
}

// FindAll retrieves all users
func (r *UserRepository) FindAll() ([]models.User, error) {
	query := `SELECT id_user, name, email, role, created_at FROM user ORDER BY id_user ASC`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// FindByID retrieves a user by ID
func (r *UserRepository) FindByID(id int) (models.User, error) {
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

// GetPasswordByID retrieves user password by ID (for update operations)
func (r *UserRepository) GetPasswordByID(id int) (string, error) {
	var password string
	err := db.DB.QueryRow("SELECT password FROM user WHERE id_user = ?", id).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// Update updates user information
func (r *UserRepository) Update(id int, user *models.User) error {
	query := `UPDATE user SET name = ?, email = ?, password = ? WHERE id_user = ?`
	_, err := db.DB.Exec(query, user.Name, user.Email, user.Password, id)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok && driverErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

// Delete removes a user from database
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM user WHERE id_user = ?`
	_, err := db.DB.Exec(query, id)
	return err
}

// CountBookingsByUserID counts bookings for a user
func (r *UserRepository) CountBookingsByUserID(id int) (int, error) {
	var count int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM booking WHERE customer_id = ?`, id).Scan(&count)
	return count, err
}

// CountEventsByUserID counts events created by organizer
func (r *UserRepository) CountEventsByUserID(id int) (int, error) {
	var count int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM event WHERE organizer_id = ?`, id).Scan(&count)
	return count, err
}