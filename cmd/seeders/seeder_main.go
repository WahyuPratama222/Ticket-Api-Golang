package main

import (
	"log"

	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"github.com/WahyuPratama222/Ticket-Api-Golang/seeders"
)

func main() {
	// Connect to database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Run All Seeders
	if err := seeders.RunAllSeeder(); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding completed successfully")
}
