package main

import (
	"log"

	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"github.com/WahyuPratama222/Ticket-Api-Golang/seeders"
)

func main() {
	// 1. Connect ke DB
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// 2. Jalankan seeder
	if err := seeders.RunAllSeeder(); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding completed successfully")
}
