package models

import "time"

// struct to create an user
type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// struct to handle login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// struct to handle sign up request
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
