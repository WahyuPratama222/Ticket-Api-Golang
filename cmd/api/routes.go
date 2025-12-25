package main

import (
	"github.com/WahyuPratama222/Ticket-Api-Golang/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Initialize all handlers
	userHandler := handler.NewUserHandler()
	eventHandler := handler.NewEventHandler()
	bookingHandler := handler.NewBookingHandler()
	ticketHandler := handler.NewTicketHandler()

	// User routes
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Event routes
	r.HandleFunc("/events", eventHandler.CreateEvent).Methods("POST")
	r.HandleFunc("/events", eventHandler.GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", eventHandler.GetEvent).Methods("GET")
	r.HandleFunc("/events/{id}", eventHandler.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", eventHandler.DeleteEvent).Methods("DELETE")

	// Booking routes
	r.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings", bookingHandler.GetAllBookings).Methods("GET")
	r.HandleFunc("/bookings/{id}", bookingHandler.GetBooking).Methods("GET")

	// Ticket routes
	r.HandleFunc("/tickets", ticketHandler.GetAllTickets).Methods("GET")
	r.HandleFunc("/tickets/{id}", ticketHandler.GetTicket).Methods("GET")
	r.HandleFunc("/tickets/{id}/use", ticketHandler.UseTicket).Methods("PUT")
	return r
}