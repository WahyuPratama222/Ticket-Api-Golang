package seeders

import "fmt"

func RunAllSeeder() error {
	if err := SeedUsers(); err != nil {
		return fmt.Errorf("failed to seed users: %v", err)
	}

	if err := SeedEvents(); err != nil {
		return fmt.Errorf("failed to seed events: %v", err)
	}

	fmt.Println("All seeders executed successfully")
	return nil
}
