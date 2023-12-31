package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	response "github.com/ssonumkar/weather-report-api/internal/http_response"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

type AuthMiddleware struct {
	secretKey string
}

func NewAuthMiddleware(secretKey string) *AuthMiddleware {
	return &AuthMiddleware{secretKey}
}

func (m *AuthMiddleware) Authenticate(tokenPool encrypt.IAuthTokenPool, logger log.CustomLogger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.UpdateEndpoint(log.Auth)
		logger.Info("Authorizing..")
		tokenString := r.Header.Get("Authorization")
		logger.Debug(tokenString)
		if tokenString == "" {
			logger.Error("Token cannot be null")
			response.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		if !tokenPool.IsValid(logger, tokenString) {
			logger.Error("token invalid, not present in pool")
			response.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.secretKey), nil
		})
		if err != nil || !token.Valid {
			logger.Error(fmt.Sprintf("Failed to authenticate token: %s", err))
			response.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		logger.Info("Authorized !")
		next.ServeHTTP(w, r)
	}
}
