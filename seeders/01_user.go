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
		{"Dedi Mulyados", "mulmulmul@mail.com", "mullllya123", "organizer"},
		{"Pak Jkw", "jwkwkw@mail.com", "jwkwkww123", "customer"},
		{"Bahlilul", "bahlilu@mail.com", "bahabahha123", "customer"},
	}

	for _, u := range users {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Pass), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		now := time.Now()
		_, err = db.DB.Exec(`
			INSERT INTO user (name, email, password, role, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE email=email
		`,
			u.Name,
			u.Email,
			string(hashed),
			u.Role,
			now,
			now,
		)

		if err != nil {
			return err
		}
	}

	return nil
}