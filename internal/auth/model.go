package auth

import "github.com/ssonumkar/weather-report-api/internal/log"

// User represents a user entity
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	DOB string 	`json:"dob"`
}

type IUserRepository interface{
	CreateUser(user User, logger log.CustomLogger) error
	GetUserByUsername(username string, logger log.CustomLogger) (User, error)
}

type IAuthService interface{
	Login(username, password string,logger log.CustomLogger) (LoginResponse, error) 
	Logout(token string, logger log.CustomLogger) error
	RegisterUser(user User, logger log.CustomLogger) (error) 
}

type LoginResponse struct{
	ID int `json:"id"`
	Username string `json:"username"`
	JwtToken string `json:"jwt_token"`
}