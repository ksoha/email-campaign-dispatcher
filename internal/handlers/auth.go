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

// sign in handler
func SignUpHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var signupRequest models.SignUpRequest

		//decoding the json
		err := json.NewDecoder(r.Body).Decode(&signupRequest)
		if err != nil {
			response.WriteError(
				w,
				http.StatusBadRequest,
				"Invalid request Body",
			)
			return
		}

		//Hash the password before storing into the database
		passwordHash, err := auth.HashPassword(signupRequest.Password)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to hash the password",
			)
			return
		}

		//Create the user in the database
		user, err := database.CreateUser(
			db,
			signupRequest.Email,
			passwordHash,
		)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to create user",
			)
			return
		}

		//returning the user details
		err = response.WriteJSON(
			w,
			http.StatusCreated,
			map[string]string{
				"message": "User created successfully",
				"user":    user.Email,
			},
		)
		if err != nil {
			response.WriteError(
				w,
				http.StatusBadGateway,
				"Failed to encode the response",
			)
			return
		}
	}
}
