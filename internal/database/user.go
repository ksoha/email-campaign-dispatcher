package database

import (
	"database/sql"

	"github.com/ksoha/email-dispatcher/internal/models"
)

func GetUserByEmail(db *sql.DB, email string) (models.User, error) {

	var user models.User

	query := `
	SELECT id, email, password_hash, created_at, updated_at
	FROM users
	WHERE email = $1`

	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// inserting new user into the database
func CreateUser(db *sql.DB, email string, passwordHash string) (models.User, error) {

	var user models.User

	query := `
	INSERT INTO users (email, password_hash)
	VALUES ($1, $2)
	RETURNING 
	id, 
	email, 
	password_hash, 
	created_at, 
	updated_at`

	err := db.QueryRow(query, email, passwordHash).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
