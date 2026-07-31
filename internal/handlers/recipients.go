package handlers

import (
	"database/sql"
	"net/http"

	"encoding/json"

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

// CreateRecipientHandler handles the POST /recipients endpoint (import)
func CreateRecipientHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var recipient models.CreateRecipientRequest

		//the decode fucntion will not return a recipient struct it will just fill the struct recipient
		err := json.NewDecoder(r.Body).Decode(&recipient)
		if err != nil {
			response.WriteError(
				w,
				http.StatusBadRequest,
				"Invalid request",
			)
			return
		}

		err = database.CreateRecipient(db, recipient)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to create recipient",
			)
			return
		}

		//everythin succeded create recipient
		err = response.WriteJSON(
			w,
			http.StatusCreated,
			map[string]string{"message": "Recipient Created Successfully"},
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
