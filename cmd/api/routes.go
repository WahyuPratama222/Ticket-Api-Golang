package main

import (
	handler "github.com/WahyuPratama222/Ticket-Api-Golang/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users/register", handler.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUserHandler).Methods("DELETE")

	// Event routes
	r.HandleFunc("/events", handler.CreateEventHandler).Methods("POST")
	r.HandleFunc("/events/{id}", handler.GetEventHandler).Methods("GET")
	r.HandleFunc("/events/{id}", handler.UpdateEventHandler).Methods("PUT")
	r.HandleFunc("/events/{id}", handler.DeleteEventHandler).Methods("DELETE")

	// Booking routes
	r.HandleFunc("/bookings", handler.CreateBookingHandler).Methods("POST")
	r.HandleFunc("/bookings/{id}", handler.GetBookingHandler).Methods("GET")

	return r
}