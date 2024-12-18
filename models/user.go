package models

type User struct {
	ID            int    `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Auth0_user_id string `json:"auth0_user_id"`
}
