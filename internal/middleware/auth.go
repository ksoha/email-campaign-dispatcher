package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/ksoha/email-dispatcher/internal/auth"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("🔥 AUTH MIDDLEWARE HIT")

		// Get Authorization header
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			log.Println("❌ ERROR: Authorization header is missing")

			http.Error(
				w,
				"Authorization header required",
				http.StatusUnauthorized,
			)
			return
		}

		log.Println("✅ Authorization header received")

		// Split:
		// Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 {
			log.Println("❌ ERROR: Authorization header does not contain two parts")

			http.Error(
				w,
				"Invalid Authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		if parts[0] != "Bearer" {
			log.Printf(
				"❌ ERROR: Invalid authorization scheme: %s\n",
				parts[0],
			)

			http.Error(
				w,
				"Invalid Authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		log.Println("✅ Bearer authentication scheme detected")

		tokenString := parts[1]

		if tokenString == "" {
			log.Println("❌ ERROR: Token is empty")

			http.Error(
				w,
				"Token is missing",
				http.StatusUnauthorized,
			)
			return
		}

		log.Println("✅ JWT token extracted")

		// Validate JWT
		userID, err := auth.ValidateToken(tokenString)
		if err != nil {

			log.Printf(
				"❌ JWT VALIDATION ERROR: %v\n",
				err,
			)

			http.Error(
				w,
				"Invalid or expired token",
				http.StatusUnauthorized,
			)
			return
		}

		log.Printf("✅ JWT VALIDATED. User ID: %s\n", userID)

		// Store authenticated user ID in request context
		ctx := context.WithValue(
			r.Context(),
			userIDKey,
			userID,
		)

		r = r.WithContext(ctx)

		log.Println("✅ User ID added to request context")

		// Continue to actual handler
		log.Println("➡️ Passing request to next handler")

		next.ServeHTTP(w, r)
	})
}
