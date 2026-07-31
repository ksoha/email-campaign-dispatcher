package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ksoha/email-dispatcher/internal/auth"
	"github.com/ksoha/email-dispatcher/internal/database"
	"github.com/ksoha/email-dispatcher/internal/models"
	"github.com/ksoha/email-dispatcher/internal/response"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var loginRequest models.LoginRequest

		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			response.WriteError(
				w,
				http.StatusBadRequest,
				"Invalid request Body",
			)
			return
		}

		user, err := database.GetUserByEmail(db, loginRequest.Email)
		if err != nil {
			response.WriteError(
				w,
				http.StatusUnauthorized,
				"Invalid email or password",
			)
			return
		}

		//check if the password matches to hash stored in the database
		err = auth.CheckPasswordHash(
			loginRequest.Password,
			user.PasswordHash,
		)
		if err != nil {

			response.WriteError(
				w,
				http.StatusUnauthorized,
				"Invalid email or password",
			)
			return
		}

		//password matches
		err = response.WriteJSON(
			w,
			http.StatusOK,
			map[string]string{
				"message": "Login successful",
			},
		)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to encode response",
			)
			return
		}

	}
}
