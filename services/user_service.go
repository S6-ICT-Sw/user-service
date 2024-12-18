package services

import (
	"context"
	"errors"

	"regexp"

	"user-service/models"
	"user-service/repository"
)

func Register(ctx context.Context, user *models.User) error {

	// Validate the email to check if it's an email
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	existEmail, err := repository.CheckIfEmailExist(ctx, user.Email)
	if err != nil {
		return err
	}

	if existEmail {
		return errors.New("email aleady exist")
	}

	return repository.CreateUser(ctx, user)
}

func GetUser(ctx context.Context, id int) (*models.GetUser, error) {
	user, err := repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if no user was found
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func DeleteUser(ctx context.Context, id int) error {
	// Validate that the user ID is positive
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	// Add if ID exist

	err := repository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func EditUser(ctx context.Context, user *models.EditUser) (*models.EditUser, error) {
	if user.ID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	// Add if ID exist

	return repository.EditUser(ctx, user)
}

func ValidateEmail(email string) error {
	// A simple regex to check if the email is in a valid format
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !regex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}
