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
