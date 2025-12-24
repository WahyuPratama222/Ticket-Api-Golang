package seeders

import (
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() error {
	type dummyUser struct {
		Name  string
		Email string
		Pass  string
		Role  string
	}

	users := []dummyUser{
		{"organizer", "organizer@mail.com", "organizer123", "organizer"},
		{"customer1", "customer1@mail.com", "customer123", "customer"},
		{"customer2", "customer2@mail.com", "customer123", "customer"},
	}

	for _, u := range users {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Pass), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		_, err = db.DB.Exec(`
			INSERT INTO user (name, email, password, role, created_at)
			VALUES (?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE email=email
		`,
			u.Name,
			u.Email,
			string(hashed),
			u.Role,
			time.Now(),
		)

		if err != nil {
			return err
		}
	}

	return nil
}
