package encrypt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

// type User string

type IAuthTokenPool interface {
	GenerateToken(logger log.CustomLogger, userID int, username string, jwtSecret string) (string, error)
	AddTokenToPool(logger log.CustomLogger, user string, token string) error
	IsValid(logger log.CustomLogger, token string) bool
	RemoveTokenFromPool(logger log.CustomLogger, token string) error
}

type AuthTokenPool struct {
	tokenPool map[string]string
}

func NewAuthTokenPool() *AuthTokenPool {
	return &AuthTokenPool{tokenPool: make(map[string]string)}
}

func isUserPresent(user string, tokenPool map[string]string) bool {
	for _, v := range tokenPool {
		if v == user {
			return true
		}
	}
	return false
}
func isTokenPresent(token string, tokenPool map[string]string) bool {
	_, found := tokenPool[token]
	return found
}

func (a *AuthTokenPool) AddTokenToPool(logger log.CustomLogger, user string, token string) error {
	if isTokenPresent(token, a.tokenPool) || isUserPresent(user, a.tokenPool) {
		logger.Error("Token for user is already present.")
		return fmt.Errorf("user is already logged in")
	}
	a.tokenPool[token] = user
	logger.Info("User token added successfully")
	return nil
}

func (a *AuthTokenPool) RemoveTokenFromPool(logger log.CustomLogger, token string) error {
	_, found := a.tokenPool[token]
	if !found {
		logger.Error("Invalid Token, not present in pool.")
		return fmt.Errorf("user not logged in")
	}
	delete(a.tokenPool, token)
	logger.Info("User token removed successfully")
	return nil
}

func (a *AuthTokenPool) IsValid(logger log.CustomLogger, token string) bool {
	_, found := a.tokenPool[token]
	logger.Info("Token validated")
	return found
}

func (a *AuthTokenPool) GenerateToken(logger log.CustomLogger, userID int, username string, jwtSecret string) (string, error) {
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
		logger.Error(fmt.Sprintf("Failed to generate jwt token: %s", err.Error()))
		return "", err
	}
	logger.Info("JWT token generated successfully")

	return tokenString, nil
}
