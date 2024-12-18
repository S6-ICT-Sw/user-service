package repository

import (
	"context"
	"user-service/models"
	"user-service/pkg/database"

	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, username, auth0_user_id) VALUES ($1, $2, $3)`
	_, err := database.DB.Exec(ctx, query, user.Email, user.Username, user.Auth0_user_id)
	return err
}

func CheckIfEmailExist(ctx context.Context, email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE email = $1 LIMIT 1`

	var exists int

	// Execute the query and scan it
	row := database.DB.QueryRow(ctx, query, email)
	err := row.Scan(&exists)

	if err != nil {
		// If no rows are found, return false (email does not exist)
		if err == pgx.ErrNoRows {
			return false, nil
		}
		// Return any unexpected errors
		return false, err
	}
	// If the query returns a row, the email exists
	return true, nil
}

func GetById(ctx context.Context, id int) (*models.GetUser, error) {
	query := `SELECT id, email, username FROM users WHERE id = $1`

	var user models.GetUser

	// Execute the query and scan the result into the user struct
	err := database.DB.QueryRow(ctx, query, id).Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		// Handle "no rows" error
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := database.DB.Exec(ctx, query, id)
	return err
}

func EditUser(ctx context.Context, user *models.EditUser) (*models.EditUser, error) {
	query := `UPDATE users SET email = COALESCE($1, email), username = COALESCE($2, username) WHERE id = $3 
	RETURNING email, username, id`

	updatedUser := &models.EditUser{}

	// Execute the query and scan the updated row
	err := database.DB.QueryRow(ctx, query, user.Email, user.Username, user.ID).Scan(
		&updatedUser.Email, &updatedUser.Username, &updatedUser.ID)

	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
