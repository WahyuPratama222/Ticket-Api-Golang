package migrations

import "github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"

func CreateEventTable() error {
	query := `CREATE TABLE IF NOT EXISTS event (
		id_event INT AUTO_INCREMENT PRIMARY KEY,
		organizer_id INT NOT NULL,
		title VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		capacity INT NOT NULL,
		available_seat INT NOT NULL,
		price INT NOT NULL,
		status ENUM('available','unavailable') NOT NULL DEFAULT 'available',
		date DATETIME NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (organizer_id) REFERENCES user(id_user)
	);`
	_, err := db.DB.Exec(query)
	return err
}