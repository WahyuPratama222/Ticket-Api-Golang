package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handler "github.com/WahyuPratama222/Ticket-Api-Golang/handlers"
	"github.com/WahyuPratama222/Ticket-Api-Golang/pkg/db"
	"github.com/gorilla/mux"
)

func main() {
	// Connect ke DB
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create server with timeouts
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		fmt.Printf("Server running on :%s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("Server exited properly")
}