package migrations

import "github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"

func CreateBookingTable() error {
	query := `CREATE TABLE IF NOT EXISTS booking (
		id_booking INT AUTO_INCREMENT PRIMARY KEY,
		customer_id INT NOT NULL,
		event_id INT NOT NULL,
		total_price INT NOT NULL,
		quantity INT NOT NULL,
		status ENUM('pending','success','failed') NOT NULL DEFAULT 'pending',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (customer_id) REFERENCES user(id_user),
		FOREIGN KEY (event_id) REFERENCES event(id_event)
	);`
	_, err := db.DB.Exec(query)
	return err
}
