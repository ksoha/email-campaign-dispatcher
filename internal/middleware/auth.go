package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ksoha/email-dispatcher/internal/auth"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(
				w,
				"Authorization header required",
				http.StatusUnauthorized,
			)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(
				w,
				"Invalid Authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		tokenString := parts[1]

		userID, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(
				w,
				"Invalid or expired token",
				http.StatusUnauthorized,
			)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			userIDKey,
			userID,
		)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
