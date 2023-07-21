package auth

import (
	"testing"

	"github.com/ssonumkar/weather-report-api/internal/auth"
	"github.com/ssonumkar/weather-report-api/internal/log"
)
var mockSuccessUserService auth.IAuthService
var mockFailUserService auth.IAuthService
var secretKey string
var logger log.CustomLogger
func init(){
	secretKey = "test_secret_key"
	mockSuccessUserService = auth.NewAuthService(&MockSuccessUserRepository{}, secretKey)
	mockFailUserService = auth.NewAuthService(&MockFailUserRepository{}, secretKey)
	logger = *log.NewCustomLogger()
}
func TestShouldPassForLogin(t *testing.T) {
	//Given
	username:= "sagar"
	password:= "password"

	expectedToken := "temp_token"
	//When
	loginResp, err := mockSuccessUserService.Login(username, password, logger)
	if err != nil{
		t.Fatal("Error while login ",err)
	}
	//Then
	if loginResp.JwtToken != expectedToken{
		t.Errorf("Expected %v but got %s", expectedToken, loginResp.JwtToken)
	}
}

func TestShouldPassForLogout(t *testing.T) {
	//Given
	token := "temp_token"

	//When
	err := mockSuccessUserService.Logout(token, logger)
	//Then

	if err != nil{
		t.Errorf("Expected no error but got %v", err.Error())
	}
}

func TestShouldPassForRegisterUser(t *testing.T) {
	//Given
	user := auth.User{ID: 1, Username: "username", Password: "password", DOB: ""}
	//When
	err := mockSuccessUserService.RegisterUser(user, logger)
	//Then
	if err != nil{
		t.Errorf("Expected no error but got %v", err.Error())
	}
}
func TestShouldNotPassForLogin(t *testing.T) {
	//Given
	username:= "sagar"
	password:= "password"

	expectedToken := "temp_token"
	//When
	loginResp, err := mockFailUserService.Login(username, password, logger)
	if err != nil{
		t.Fatal("Error while login ",err)
	}
	//Then
	if loginResp.JwtToken == expectedToken{
		t.Errorf("%v should not be same as %s", expectedToken, loginResp.JwtToken)
	}
}

func TestShouldNotPassForLogout(t *testing.T) {
	//Given
	token := "temp_token"

	//When
	err := mockFailUserService.Logout(token, logger)
	//Then

	if err == nil{
		t.Errorf("Expected error but got nil")
	}
}

func TestShouldnotPassForRegisterUser(t *testing.T) {
	//Given
	user := auth.User{ID: 1, Username: "username", Password: "password", DOB: ""}
	//When
	err := mockFailUserService.RegisterUser(user, logger)
	//Then
	if err == nil{
		t.Errorf("Expected error but got nil")
	}
}