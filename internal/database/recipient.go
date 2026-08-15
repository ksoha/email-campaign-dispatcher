package database

import (
	"database/sql"
	"fmt"

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

// New function for batch insertion into the database
func CreateRecipientsBatch(db *sql.DB, recipients []models.CreateRecipientRequest) error {

	//check if the recipients slice is empty
	if len(recipients) == 0 {
		return nil
	}

	//create a transaction
	//begin the transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	//if anything fails before commit , rollback the transaction
	defer tx.Rollback()

	// Build the multi-row INSERT query
	query := `
		INSERT INTO recipients (name, email)
		VALUES
	`

	//batching through multiple inserts in PSQL
	args := make([]interface{}, 0, len(recipients)*2)

	//build the query string with placeholders for each recipient
	for i, recipient := range recipients {
		// PostgreSQL placeholders:
		// batch 0 -> $1, $2
		// batch 1 -> $3, $4
		// batch 2 -> $5, $6
		position := i * 2

		if i > 0 {
			//add a coma to separate the values for each recipient
			query += ","
		}

		//append the placeholders for each recipient to the query string
		query += fmt.Sprintf(
			"($%d, $%d)",
			position+1,
			position+2,
		)

		args = append(
			args,
			recipient.Name,
			recipient.Email,
		)
	}

	//execute one SQL query for entire batch of recipients
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	//commit the transaction
	return tx.Commit()
}
