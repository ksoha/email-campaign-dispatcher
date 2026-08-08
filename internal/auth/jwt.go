package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key")

// Claims struct represents the JWT claims used for authentication and authorization.
type Claims struct {
	UserID               string `json:"user_id"`
	jwt.RegisteredClaims        //struct embedding
}

func GenerateToken(userID string) (string, error) {
	claims := Claims{ //creating an instance of the Claim struct
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims( //creating a new JWT token with the specified signing method and claims
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// function to validate the JWT token
func ValidateToken(tokenString string) (string, error) {
	//empty claims struct that will be populated with the claims from the token
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			//make sure he token uses the expected
			//HMAC signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			//use the same secret that was used
			//when generating the token .
			return jwtSecret, nil
		},
	)

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	return claims.UserID, nil
}
