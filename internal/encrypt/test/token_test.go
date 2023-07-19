package encrypt

import (
	"testing"

	"github.com/ssonumkar/weather-report-api/internal/encrypt"
)

var TesttokenPool map[string]bool

//Positive scenarios
func TestShouldSucceedForGenerateToken(t *testing.T){
	//Given
	userID := 1
	userName:= "sa1"
	jwtSecret:= "secret1"

	//When
	_, err := encrypt.GenerateToken(userID, userName, jwtSecret)
	if err != nil{
		t.Errorf("Expected error was nil, but got %s", err.Error())
	}
}


func TestShouldSucceedForAddTokenToPool(t *testing.T){
	//Given
	token := "jwt_token"

	//when
	err := encrypt.AddTokenToPool(token)
	//then
	if err != nil{
		t.Errorf("Expected error was nil, but got %s", err.Error())
	}
}

func TestShouldSucceedForRemoveTokenFromPool(t *testing.T){
	token := "jwt_token"

	//when
	encrypt.AddTokenToPool(token)
	err := encrypt.RemoveTokenFromPool(token)
	//then
	if err != nil{
		t.Errorf("Expected error was nil, but got %s", err.Error())
	}
}

func TestShouldSucceedForIsLoggedIn(t *testing.T){
	token := "jwt_token"
	exPectedStatus := true
	//when
	encrypt.AddTokenToPool(token)
	status := encrypt.IsLoggedIn(token)
	//then
	if status != exPectedStatus{
		t.Errorf("Expected status was %v, but got %v", exPectedStatus, status)
	}
}

//Negative Scenarios
func TestShouldFailForAddTokenToPool(t *testing.T){
	//given
	token := "jwt_token"
	//when
	encrypt.AddTokenToPool(token)
	err := encrypt.AddTokenToPool(token)
	if err == nil{
		t.Errorf("Expected error was 'user already logged in', but got nil")
	}
}

func TestShouldFailForRemoveTokenFromPool(t *testing.T){
	token := "jwt_token"

	//when
	err := encrypt.RemoveTokenFromPool(token)
	//then
	if err == nil{
		t.Errorf("Expected error was 'user not logged in', but got nil")
	}
}

func TestShouldFailForIsLoggedIn(t *testing.T){
	token := "jwt_token"
	exPectedStatus := false
	//when
	status := encrypt.IsLoggedIn(token)
	//then
	if status != exPectedStatus{
		t.Errorf("Expected login status was %v, but got %v", exPectedStatus, status)
	}
}

