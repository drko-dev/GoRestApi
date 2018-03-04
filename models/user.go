package models

// Users estructura que representa la bd user
type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
