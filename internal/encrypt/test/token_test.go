package encrypt

import (
	"testing"

	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

var TesttokenPool map[string]bool
var tokenPool encrypt.IAuthTokenPool = encrypt.NewAuthTokenPool()
var logger log.CustomLogger = *log.NewCustomLogger()

// Positive scenarios
func TestShouldSucceedForGenerateToken(t *testing.T) {
	//Given
	userID := 1
	userName := "sa1"
	jwtSecret := "secret1"

	//When
	_, err := tokenPool.GenerateToken(logger, userID, userName, jwtSecret)
	if err != nil {
		t.Errorf("Expected error was nil, but got %s", err.Error())
	}
}

func TestShouldSucceedForAddTokenToPool(t *testing.T) {
	//Given
	token := "jwt_token"
	user := "user1"
	//when
	err := tokenPool.AddTokenToPool(logger, user, token)
	//then
	if err != nil {
		t.Errorf("Expected error was nil, but got %s", err.Error())
	}
}

func TestShouldSucceedForRemoveTokenFromPool(t *testing.T) {
	token := "jwt_token"
	user := "dummy"
	//when
	tokenPool.AddTokenToPool(logger, user, token)
	err := tokenPool.RemoveTokenFromPool(logger, token)
	//then
	if err != nil {
		t.Errorf("Expected error was nil, but got %s", err.Error())
	}
}

func TestShouldSucceedForIsLoggedIn(t *testing.T) {
	token := "jwt_token"
	user := "dummy"
	exPectedStatus := true
	//when
	tokenPool.AddTokenToPool(logger, user, token)
	status := tokenPool.IsValid(logger, token)
	//then
	if status != exPectedStatus {
		t.Errorf("Expected status was %v, but got %v", exPectedStatus, status)
	}
}

// Negative Scenarios
func TestShouldFailForAddTokenToPool(t *testing.T) {
	//given
	token := "jwt_token"
	user := "dummy"

	//when
	tokenPool.AddTokenToPool(logger, user, token)
	err := tokenPool.AddTokenToPool(logger, user, token)
	if err == nil {
		t.Errorf("Expected error was 'user already logged in', but got nil")
	}
}

func TestShouldFailForRemoveTokenFromPool(t *testing.T) {
	token := "fake_token"

	//when
	err := tokenPool.RemoveTokenFromPool(logger, token)
	//then
	if err == nil {
		t.Errorf("Expected error was 'user not logged in', but got nil")
	}
}

func TestShouldFailForIsLoggedIn(t *testing.T) {
	token := "fake_token"
	exPectedStatus := false
	//when
	status := tokenPool.IsValid(logger, token)
	//then
	if status != exPectedStatus {
		t.Errorf("Expected login status was %v, but got %v", exPectedStatus, status)
	}
}
