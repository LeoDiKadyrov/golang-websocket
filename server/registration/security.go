package registration

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(password string) (string, []byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", nil, err
	}
	combined := append([]byte(password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	return string(hashedPassword), salt, nil
}
