package auth

import (
	"fmt"

	"github.com/ssonumkar/weather-report-api/internal/log"
)

type MockSuccessTokenPool struct {
}

func (m *MockSuccessTokenPool) GenerateToken(logger log.CustomLogger, userID int, username string, jwtSecret string) (string, error) {
	return "dummy_token", nil
}
func (m *MockSuccessTokenPool) AddTokenToPool(logger log.CustomLogger, user string, token string) error {
	return nil
}
func (m *MockSuccessTokenPool) IsValid(logger log.CustomLogger, token string) bool {
	return true
}
func (m *MockSuccessTokenPool) RemoveTokenFromPool(logger log.CustomLogger, token string) error {
	return nil
}

type MockFailTokenPool struct {
}

func (m *MockFailTokenPool) GenerateToken(logger log.CustomLogger, userID int, username string, jwtSecret string) (string, error) {
	return "", fmt.Errorf("test error")
}
func (m *MockFailTokenPool) AddTokenToPool(logger log.CustomLogger, user string, token string) error {
	return fmt.Errorf("test error")
}
func (m *MockFailTokenPool) IsValid(logger log.CustomLogger, token string) bool {
	return false
}
func (m *MockFailTokenPool) RemoveTokenFromPool(logger log.CustomLogger, token string) error {
	return fmt.Errorf("test error")
}
