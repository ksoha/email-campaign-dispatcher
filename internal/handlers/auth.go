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

		// Check if password matches the hash stored in database
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

		// Email + password are valid.
		// Generate JWT for this user.
		token, err := auth.GenerateToken(user.ID)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to generate token",
			)
			return
		}

		// Return token to client
		err = response.WriteJSON(
			w,
			http.StatusOK,
			map[string]string{
				"message": "Login successful",
				"token":   token,
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
