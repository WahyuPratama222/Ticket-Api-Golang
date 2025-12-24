package migrations

import (
	"fmt"
)

func MigrateAll() error {
	if err := CreateUserTable(); err != nil {
		return fmt.Errorf("failed to create user table: %v", err)
	}
	if err := CreateEventTable(); err != nil {
		return fmt.Errorf("failed to create event table: %v", err)
	}
	if err := CreateBookingTable(); err != nil {
		return fmt.Errorf("failed to create booking table: %v", err)
	}
	if err := CreateTicketTable(); err != nil {
		return fmt.Errorf("failed to create ticket table: %v", err)
	}
	fmt.Println("All tables migrated successfully")
	return nil
}
