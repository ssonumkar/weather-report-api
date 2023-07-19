package auth

import (
	"fmt"

	"github.com/ssonumkar/weather-report-api/internal/auth"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

type MockSuccessUserRepository struct {
}

func (m *MockSuccessUserRepository) CreateUser(user auth.User, logger log.CustomLogger) error{
	return nil
}
func (m *MockSuccessUserRepository) GetUserByUsername(username string, logger log.CustomLogger) (auth.User, error){
	return auth.User{ID: 12, Username: "user2", Password: "pass", DOB: "11-2-2021"}, nil
}

type MockFailUserRepository struct {
}

func (m *MockFailUserRepository) CreateUser(user auth.User, logger log.CustomLogger) error{
	return fmt.Errorf("")
}
func (m *MockFailUserRepository) GetUserByUsername(username string, logger log.CustomLogger) (auth.User, error){
	return auth.User{}, fmt.Errorf("")
}