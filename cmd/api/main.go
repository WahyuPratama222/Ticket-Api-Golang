package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/WahyuPratama222/Ticket-Api-Golang/handlers"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"github.com/WahyuPratama222/Ticket-Api-Golang/migrations"
	"github.com/gorilla/mux"
)

func main() {

	// Connect ke DB
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Jalankan semua migration
	if err := migrations.MigrateAll(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/users/register", handler.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/events", handler.CreateEventHandler).Methods("POST")
	r.HandleFunc("/events/{id}", handler.GetEventHandler).Methods("GET")
	r.HandleFunc("/events/{id}", handler.UpdateEventHandler).Methods("PUT")
	r.HandleFunc("/events/{id}", handler.DeleteEventHandler).Methods("DELETE")

	r.HandleFunc("/bookings", handler.CreateBookingHandler).Methods("POST")
	r.HandleFunc("/bookings/{id}", handler.GetBookingHandler).Methods("GET")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
