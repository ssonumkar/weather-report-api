package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	response "github.com/ssonumkar/weather-report-api/internal/http_response"
)

// AuthMiddleware is a middleware for authentication
type AuthMiddleware struct {
	secretKey string
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware(secretKey string) *AuthMiddleware {
	return &AuthMiddleware{secretKey}
}

// Authenticate is the middleware function for authentication
func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			response.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		fmt.Println(tokenString)

		if !encrypt.IsLoggedIn(tokenString) {
			log.Println("token invalid, not present in pool")
			response.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.secretKey), nil
		})
		if err != nil || !token.Valid {
			log.Println("Failed to authenticate token:", err)
			response.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	}
}
