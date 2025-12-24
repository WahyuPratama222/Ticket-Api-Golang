package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"golang.org/x/crypto/bcrypt"

	mysql "github.com/go-sql-driver/mysql"
)

// CreateUser membuat user baru
func CreateUser(user models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("all fields are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO user (name, password, email, role, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err = db.DB.Exec(query, user.Name, string(hashedPassword), user.Email, user.Role, time.Now())
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok && driverErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}

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

	if updated.Name == "" {
		updated.Name = user.Name
	}
	if updated.Email == "" {
		updated.Email = user.Email
	}
	if updated.Role == "" {
		updated.Role = user.Role
	}

	var hashedPassword string
	if updated.Password != "" {
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

	query := `DELETE FROM user WHERE id_user = ?`
	_, err = db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
