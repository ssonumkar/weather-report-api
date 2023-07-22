package auth

import (
	"fmt"

	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

// AuthService handles authentication-related operations
type AuthService struct {
	userRepository  IUserRepository
	tokenPool       encrypt.IAuthTokenPool
	passwordManager encrypt.IPasswordManager
	secretKey       string
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepository IUserRepository, tokenPool encrypt.IAuthTokenPool, passwordManager encrypt.IPasswordManager, secretKey string) *AuthService {
	return &AuthService{userRepository, tokenPool, passwordManager, secretKey}
}

// Login performs user login
func (s *AuthService) Login(username, password string, logger log.CustomLogger) (LoginResponse, error) {

	user, err := s.userRepository.GetUserByUsername(username, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Error finding user: %s", err.Error()))
		return LoginResponse{}, err
	}
	logger.Info("User found")
	// Compare the provided password with the stored password
	err = s.passwordManager.ComparePasswords(user.Password, password, logger)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("incorrect Password")
	}
	// Generate a JWT token
	token, err := s.tokenPool.GenerateToken(logger, user.ID, user.Username, s.secretKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Error generating token: %s", err.Error()))
		return LoginResponse{}, err
	}
	err = s.tokenPool.AddTokenToPool(logger, username, token)
	if err != nil {
		logger.Error(fmt.Sprintf("Error adding token to pool: %s", err.Error()))
		return LoginResponse{}, fmt.Errorf("user already has a session logged in")
	}
	logger.Debug(fmt.Sprintf("Token added to pool: %s", token))
	return LoginResponse{user.ID, user.Username, token}, nil
}

// Logout performs user logout
func (s *AuthService) Logout(token string, logger log.CustomLogger) error {
	err := s.tokenPool.RemoveTokenFromPool(logger, token)
	if err != nil {
		logger.Error("User not logged in")
		return err
	}
	return nil
}

func (s *AuthService) RegisterUser(user User, logger log.CustomLogger) error {
	var err error
	user.Password, err = s.passwordManager.EncryptPassword(user.Password, logger)
	if err != nil {
		return err
	}
	err = s.userRepository.CreateUser(user, logger)
	if err != nil {
		return err
	}
	return nil
}
