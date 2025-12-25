package seeders

import (
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
)

func SeedEvents() error {
	events := []struct {
		OrganizerID int
		Title       string
		Location    string
		Capacity    int
		Price       int
		Date        time.Time
	}{
		{1, "Konser Rock", "Jakarta", 100, 50000, time.Now().AddDate(0, 1, 0)},
		{1, "Seminar Tech", "Bandung", 50, 75000, time.Now().AddDate(0, 2, 0)},
	}

	for _, e := range events {
		now := time.Now()
		_, err := db.DB.Exec(`
			INSERT INTO event 
			(organizer_id, title, location, capacity, available_seat, price, status, date, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, 'available', ?, ?, ?)
			ON DUPLICATE KEY UPDATE title=title
		`,
			e.OrganizerID,
			e.Title,
			e.Location,
			e.Capacity,
			e.Capacity,
			e.Price,
			e.Date,
			now,
			now,
		)

		if err != nil {
			return err
		}
	}

	return nil
}