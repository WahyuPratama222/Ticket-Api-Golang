package main

import (
	"log"

	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"github.com/WahyuPratama222/Ticket-Api-Golang/migrations"
)

func main() {
	// Connect ke database (DB sudah ada)
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Jalankan semua migration (buat tabel)
	if err := migrations.MigrateAll(); err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed successfully")
}
