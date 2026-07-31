package database

import (
	"database/sql"

	"github.com/ksoha/email-dispatcher/internal/models"
)

// recipient query
func GetRecipient(db *sql.DB) ([]models.Recipient, error) {
	//slice to store the recipients
	var recipients []models.Recipient

	//query to get the recipients from the database
	rows, err := db.Query(`
	       SELECT 
		   id, 
		   name, 
		   email, 
		   created_at, 
		   updated_at
		   FROM recipients 
    `)
	if err != nil {
		return nil, err
	}

	//defef the close rows
	defer rows.Close()

	//loop throuh every row and scan the data into the recipient struct
	for rows.Next() {
		var recipient models.Recipient
		err := rows.Scan(
			&recipient.ID,
			&recipient.Name,
			&recipient.Email,
			&recipient.CreatedAt,
			&recipient.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		recipients = append(recipients, recipient)
	}
	return recipients, nil
}

//to create new recipient in the database

func CreateRecipient(db *sql.DB, recipient models.CreateRecipientRequest) error {

	query := `
	 INSERT INTO recipients (name, email)
	 VALUES ($1, $2)`

	//insert the recipient into the database
	_, err := db.Exec(
		query,
		recipient.Name,
		recipient.Email,
	)
	if err != nil {
		return err
	}
	return nil
}
