package service

import (
	"errors"
	"time"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
	"github.com/WahyuPratama222/Ticket-Api-Golang/repositories"
	"github.com/WahyuPratama222/Ticket-Api-Golang/validations"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles user business logic
type UserService struct {
	repo      *repositories.UserRepository
	validator *validations.UserValidator
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		repo:      repositories.NewUserRepository(),
		validator: validations.NewUserValidator(),
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) error {
	// Validate input
	if err := s.validator.ValidateCreate(user); err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	// Save to database
	return s.repo.Create(user)
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

// GetUserByID retrieves user by ID
func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.repo.FindByID(id)
}

// UpdateUser updates user information (without role)
func (s *UserService) UpdateUser(id int, updated models.User) error {
	// Get existing user
	existingUser, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Preserve existing role (role cannot be updated)
	updated.Role = existingUser.Role

	// Validate update data
	if err := s.validator.ValidateUpdate(&updated, &existingUser); err != nil {
		return err
	}

	// Handle password update
	if updated.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(updated.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updated.Password = string(hash)
	} else {
		// Keep existing password
		existingPassword, err := s.repo.GetPasswordByID(id)
		if err != nil {
			return err
		}
		updated.Password = existingPassword
	}

	// Update in database
	return s.repo.Update(id, &updated)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int) error {
	// Check if user exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Check if user has any bookings
	bookingCount, err := s.repo.CountBookingsByUserID(id)
	if err != nil {
		return err
	}
	if bookingCount > 0 {
		return errors.New("cannot delete user with existing bookings")
	}

	// Check if user (organizer) has any events
	eventCount, err := s.repo.CountEventsByUserID(id)
	if err != nil {
		return err
	}
	if eventCount > 0 {
		return errors.New("cannot delete organizer with existing events")
	}

	// Delete user
	return s.repo.Delete(id)
}