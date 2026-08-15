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

		// Read and process one record at a time
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

			// Insert recipient into database
			err = database.CreateRecipient(
				db,
				recipient,
			)
			if err != nil {
				response.WriteError(
					w,
					http.StatusInternalServerError,
					"Failed to create recipient",
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
