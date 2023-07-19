package encrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// ComparePasswords compares a provided password with a stored hashed password
func ComparePasswords(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err
}

func EncryptPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println(string(hashedPassword))
	t := string(hashedPassword)
	err = bcrypt.CompareHashAndPassword([]byte(t), []byte(plainPassword))
    fmt.Println("isMatch: ",err) // nil means it is a match
	return string(hashedPassword), nil
}
