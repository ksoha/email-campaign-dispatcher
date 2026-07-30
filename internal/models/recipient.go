package models

import "time"

// struct of a receipient in the database

type Recipient struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
