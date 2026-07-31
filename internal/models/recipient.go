package models

import "time"

// struct representing a recipient stored in the database
type Recipient struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// struct representing a request body for creating a new recipient
type CreateRecipientRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
