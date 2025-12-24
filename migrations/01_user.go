package migrations

import "github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"

func CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS user (
		id_user INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		role ENUM('customer','organizer') NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.DB.Exec(query)
	return err
}