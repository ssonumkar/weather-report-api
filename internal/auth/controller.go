package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	response "github.com/ssonumkar/weather-report-api/internal/http_response"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

// AuthController handles authentication-related requests
type AuthController struct {
	authService IAuthService
	logger log.CustomLogger
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService IAuthService, logger log.CustomLogger) *AuthController {
	return &AuthController{authService, logger}
}

// Login handles the login request
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	c.logger.UpdateEndpoint(log.Login)
	c.logger.Info("-------------------------------------------------")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	c.logger.Debug(fmt.Sprintf("User username received is %s", user.Username))

	// Call the authentication service to handle login, passing the database connection
	token, err := c.authService.Login(user.Username, user.Password, c.logger)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %s ", err))
		return
	}
	c.logger.Info("Login Successful")
	response.RespondWithJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	c.logger.UpdateEndpoint(log.Logout)
	c.logger.Info("-------------------------------------------------")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	err := c.authService.Logout(tokenString, c.logger)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.logger.Info("Logout Successful")
	response.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Logout successful",
	})
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	c.logger.UpdateEndpoint(log.Register)
	c.logger.Info("-------------------------------------------------")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error ")
		return
	}

	err = c.authService.RegisterUser(user, c.logger)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error ")
		return
	}
	c.logger.Info("User creation successful")
	response.RespondWithJSON(w, http.StatusAccepted, map[string]string{
		"message": "status successfull",
	})
}
