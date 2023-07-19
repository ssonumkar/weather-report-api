package encrypt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var tokenPool map[string]bool

func InitTokenPool() {
	tokenPool = make(map[string]bool)
}
func AddTokenToPool(token string) error {
	_, found := tokenPool[token]
	if found {
		return fmt.Errorf("user already logged in")
	}
	tokenPool[token] = true
	return nil
}

func RemoveTokenFromPool(token string) error {
	_, found := tokenPool[token]
	if !found {
		return fmt.Errorf("user not logged in")
	}
	delete(tokenPool, token)
	return nil
}

func IsLoggedIn(token string) bool {
	_, found := tokenPool[token]
	return found
}

// GenerateToken generates a JWT token for the specified user ID and username
func GenerateToken(userID int, username string, jwtSecret string) (string, error) {
	// Define the claims for the JWT token
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours from now)
	}

	// Create a new token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
