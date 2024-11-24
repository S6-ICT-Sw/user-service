package repository

import (
	"context"
	"user-service/models"
	"user-service/pkg/database"
)

func CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3)`
	_, err := database.DB.Exec(ctx, query, user.Email, user.Username, user.Password)
	return err
}
