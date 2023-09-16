package registration

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(password string) (string, error) {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        return "", err
    }

	combined := append([]byte(password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

    // Store the salt and hashed password in the database
    saltAndPassword := append(salt, hashedPassword...)
    return string(saltAndPassword), nil
}