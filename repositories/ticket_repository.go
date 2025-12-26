package repositories

import (
	"database/sql"
	"errors"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
)

type TicketRepository struct{}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

// FindAll retrieves all tickets
func (r *TicketRepository) FindAll() ([]models.Ticket, error) {
	query := `SELECT id_ticket, booking_id, holder_name, ticket_code, status, created_at, updated_at 
	          FROM ticket ORDER BY id_ticket ASC`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := []models.Ticket{}
	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(
			&ticket.ID,
			&ticket.BookingID,
			&ticket.HolderName,
			&ticket.TicketCode,
			&ticket.Status,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

// FindByID retrieves a ticket by ID
func (r *TicketRepository) FindByID(id int) (models.Ticket, error) {
	var ticket models.Ticket
	query := `SELECT id_ticket, booking_id, holder_name, ticket_code, status, created_at, updated_at 
	          FROM ticket WHERE id_ticket = ?`

	row := db.DB.QueryRow(query, id)
	err := row.Scan(
		&ticket.ID,
		&ticket.BookingID,
		&ticket.HolderName,
		&ticket.TicketCode,
		&ticket.Status,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return ticket, errors.New("ticket not found")
		}
		return ticket, err
	}

	return ticket, nil
}

// UpdateStatus updates ticket status to 'used'
func (r *TicketRepository) UpdateStatus(id int, status string) error {
	query := `UPDATE ticket SET status = ? WHERE id_ticket = ?`
	result, err := db.DB.Exec(query, status, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("ticket not found")
	}

	return nil
}