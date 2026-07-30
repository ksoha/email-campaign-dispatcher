package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ksoha/email-dispatcher/internal/database"
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
