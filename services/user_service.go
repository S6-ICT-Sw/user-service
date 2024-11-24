package services

import (
	"context"
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"user-service/models"
	"user-service/repository"
)

func Register(ctx context.Context, user *models.User) error {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	if err := validateEmail(user.Email); err != nil {
		return err
	}

	// Set the hashed password on the user model
	user.Password = string(hashedPassword)

	return repository.CreateUser(ctx, user)
}

func validateEmail(email string) error {
	// A simple regex to check if the email is in a valid format
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !regex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}
