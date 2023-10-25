package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil
}
