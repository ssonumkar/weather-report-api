package encrypt

import (
	"testing"

	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	"golang.org/x/crypto/bcrypt"
)

var testPasswordManager encrypt.IPasswordManager = encrypt.NewPasswordManager()

func TestSuccessComparePasswords(t *testing.T) {
	//given
	plainPassword := "plaintextpassowrd"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	//when
	err = testPasswordManager.ComparePasswords(string(hashedPassword), plainPassword, logger)
	//then
	if err != nil {
		t.Errorf("Expected no error but got err as %s", err.Error())
	}
}

func TestShouldFailComparePasswords(t *testing.T) {
	//given
	plainPassword := "plaintextpassowrd"
	hashedPassword := "hashedPassword"
	//when
	err := testPasswordManager.ComparePasswords(string(hashedPassword), plainPassword, logger)
	//then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}
