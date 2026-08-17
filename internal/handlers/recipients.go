package handlers

import (
	"database/sql"
	"encoding/csv"
	"io"
	"net/http"

	"github.com/ksoha/email-dispatcher/internal/database"
	"github.com/ksoha/email-dispatcher/internal/models"
	"github.com/ksoha/email-dispatcher/internal/response"
)

// GetRecipientsHandler handles the GET /recipients endpoint
// will pass the database a dependancy injection
func GetRecipientsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		recipients, err := database.GetRecipient(db)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				err.Error(),
			)
			return
		}
		err = response.WriteJSON(
			w,
			http.StatusOK,
			recipients,
		)
		if err != nil {
			http.Error(
				w,
				"Failed to encode response",
				http.StatusInternalServerError,
			)
			return
		}
	}
}

// CreateRecipientHandler handles POST /recipients/import
func CreateRecipientHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the uploaded CSV file
		file, _, err := r.FormFile("file")
		if err != nil {
			response.WriteError(
				w,
				http.StatusBadRequest,
				"CSV file is required",
			)
			return
		}
		defer file.Close()

		// Create a CSV reader
		reader := csv.NewReader(file)

		// Batch size of 1000 recipients to be inserted into the database at once
		const batchSize = 1000

		// Create a batch to hold recipients
		batch := make(
			[]models.CreateRecipientRequest,
			0,
			batchSize,
		)

		// Read the header row
		_, err = reader.Read()
		if err != nil {
			response.WriteError(
				w,
				http.StatusBadRequest,
				"CSV file is empty",
			)
			return
		}

		for {

			record, err := reader.Read()

			// End of CSV
			if err == io.EOF {
				break
			}

			// CSV parsing error
			if err != nil {
				response.WriteError(
					w,
					http.StatusBadRequest,
					"Failed to read CSV file",
				)
				return
			}

			// Every row should contain name and email
			if len(record) < 2 {
				response.WriteError(
					w,
					http.StatusBadRequest,
					"Invalid CSV row",
				)
				return
			}

			recipient := models.CreateRecipientRequest{
				Name:  record[0],
				Email: record[1],
			}

			// Add recipient to the current batch
			batch = append(batch, recipient)

			// When batch reaches 1000 recipients,
			// insert the entire batch into the database
			if len(batch) == batchSize {

				err = database.CreateRecipientsBatch(
					db,
					batch,
				)

				if err != nil {
					response.WriteError(
						w,
						http.StatusInternalServerError,
						"Failed to create recipients",
					)
					return
				}

				// Clear the batch while keeping
				// the allocated memory
				batch = batch[:0]
			}
		}

		// Insert any remaining recipients
		// if the total isn't exactly divisible by 1000
		if len(batch) > 0 {

			err = database.CreateRecipientsBatch(
				db,
				batch,
			)

			if err != nil {
				response.WriteError(
					w,
					http.StatusInternalServerError,
					"Failed to create recipients",
				)
				return
			}
		}

		// Everything succeeded
		err = response.WriteJSON(
			w,
			http.StatusCreated,
			map[string]string{
				"message": "Recipients imported successfully",
			},
		)

		if err != nil {
			http.Error(
				w,
				"Failed to encode response",
				http.StatusInternalServerError,
			)
			return
		}
	}
}
