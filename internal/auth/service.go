package auth

import (
	"fmt"

	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

// AuthService handles authentication-related operations
type AuthService struct {
	userRepository IUserRepository
	secretKey string
}

var blackListTokens map[string]string
// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepository IUserRepository, secretKey string) *AuthService {
	return &AuthService{userRepository, secretKey}
}

// Login performs user login
func (s *AuthService) Login(username, password string, logger log.CustomLogger) (string, error) {
	
	user, err := s.userRepository.GetUserByUsername(username, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Error finding user: %s",err.Error()))
		return "", err
	}
	logger.Info("User found")
	// Compare the provided password with the stored password
	err =encrypt.ComparePasswords(user.Password, password)
	if err != nil {
		logger.Error(fmt.Sprintf("Error comparing passwords: %s",err.Error()))
		return "", fmt.Errorf("incorrect Password")
	}
	// Generate a JWT token
	token, err :=encrypt.GenerateToken(user.ID, user.Username, s.secretKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Error generating token: %s",err.Error()))
		return "", err
	}
	encrypt.AddTokenToPool(token)
	logger.Debug(fmt.Sprintf("Token added to pool: %s", token))
	return token, nil
}

// Logout performs user logout
func (s *AuthService) Logout(token string, logger log.CustomLogger) error {
	err  :=encrypt.RemoveTokenFromPool(token)
	if err != nil{
		logger.Error("User not logged in")
		return err
	}
	return nil;
}

func (s *AuthService) RegisterUser(user User, logger log.CustomLogger) (error){
	var err error
	user.Password, err =encrypt.EncryptPassword(user.Password)
	if err != nil{
		logger.Error(fmt.Sprint("could not encrypt the password", err.Error()))
		return err
	}
	logger.Debug("Password encrypted successfully")
	err = s.userRepository.CreateUser(user, logger)
	if err != nil{
		return err
	}
	return nil
}
