package encrypt

import (
	"fmt"

	"github.com/ssonumkar/weather-report-api/internal/log"
	"golang.org/x/crypto/bcrypt"
)

type IPasswordManager interface {
	ComparePasswords(hashedPassword string, plainPassword string, logger log.CustomLogger) error
	EncryptPassword(plainPassword string, logger log.CustomLogger) (string, error)
}

type PasswordManager struct {
}

func NewPasswordManager() *PasswordManager {
	return &PasswordManager{}
}
func (p *PasswordManager) ComparePasswords(hashedPassword, plainPassword string, logger log.CustomLogger) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		logger.Error(fmt.Sprintf("Passwords do not Match : %s", err.Error()))
	}
	logger.Debug("Password matched")
	return err
}

func (p *PasswordManager) EncryptPassword(plainPassword string, logger log.CustomLogger) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(fmt.Sprintf("Error encrypting password: %s", err.Error()))
		return "", err
	}
	logger.Debug("Password encryption successful")
	return string(hashedPassword), nil
}
