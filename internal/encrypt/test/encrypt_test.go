package encrypt

import (
	"testing"

	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	"golang.org/x/crypto/bcrypt"
)

func TestSuccessComparePasswords(t *testing.T) {
	//given
	plainPassword := "plaintextpassowrd"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	//when
	err = encrypt.ComparePasswords(string(hashedPassword), plainPassword)
	//then
	if err != nil{
		t.Errorf("Expected no error but got err as %s", err.Error())
	}
}

func TestShouldFailComparePasswords(t *testing.T){
	//given
	plainPassword := "plaintextpassowrd"
	hashedPassword:= "hashedPassword"
	//when
	err := encrypt.ComparePasswords(string(hashedPassword), plainPassword)
	//then
	if err != nil{
		t.Errorf("Expected no error but got err as %s", err.Error())
	}	
}
