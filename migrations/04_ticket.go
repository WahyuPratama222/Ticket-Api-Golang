package migrations

import "github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"

func CreateTicketTable() error {
	query := `CREATE TABLE IF NOT EXISTS ticket (
		id_ticket INT AUTO_INCREMENT PRIMARY KEY,
		booking_id INT NOT NULL,
		holder_name VARCHAR(100) NOT NULL,
		ticket_code VARCHAR(50) UNIQUE NOT NULL,
		status ENUM('unused','used') NOT NULL DEFAULT 'unused',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (booking_id) REFERENCES booking(id_booking)
	);`
	_, err := db.DB.Exec(query)
	return err
}
