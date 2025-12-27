package validations

import (
	"errors"
	"regexp"

	"github.com/WahyuPratama222/Ticket-Api-Golang/models"
)

// UserValidator handles user input validation
type UserValidator struct{}

// NewUserValidator creates a new user validator
func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

// ValidateEmail checks if email format is valid
func (v *UserValidator) ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword checks if password meets minimum requirements
func (v *UserValidator) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

// ValidateRole checks if role is valid
func (v *UserValidator) ValidateRole(role string) error {
	if role != "customer" && role != "organizer" {
		return errors.New("role must be either 'customer' or 'organizer'")
	}
	return nil
}

// ValidateCreate validates user data for creation
func (v *UserValidator) ValidateCreate(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("all fields are required")
	}

	if err := v.ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := v.ValidatePassword(user.Password); err != nil {
		return err
	}

	if err := v.ValidateRole(user.Role); err != nil {
		return err
	}

	return nil
}

// ValidateUpdate validates user data for update
func (v *UserValidator) ValidateUpdate(updated *models.User, existing *models.User) error {
	// Preserve existing values if not provided
	if updated.Name == "" {
		updated.Name = existing.Name
	}
	if updated.Email == "" {
		updated.Email = existing.Email
	} else {
		if err := v.ValidateEmail(updated.Email); err != nil {
			return err
		}
	}
	if updated.Role == "" {
		updated.Role = existing.Role
	} else {
		if err := v.ValidateRole(updated.Role); err != nil {
			return err
		}
	}

	// Validate password if provided
	if updated.Password != "" {
		if err := v.ValidatePassword(updated.Password); err != nil {
			return err
		}
	}

	return nil
}