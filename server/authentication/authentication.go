package authentication

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	postgresdb "websocket_1/server/database"
	"websocket_1/server/models"
	"websocket_1/server/security"
)

func AuthenticateUser(username, password string) error {
	fmt.Println("authentication package")
	var user models.User

	db := postgresdb.GetInstanceDB().DB
	hashedPassword, Salt, err := security.HashAndSaltPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("SELECT username, password_hash, salt FROM users WHERE username = '?';")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.PasswordHash, &user.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username")
		}
		return err
	}

	if hashedPassword != user.PasswordHash || string(Salt) != string(user.Salt) {
		return errors.New("invalid password")
	}

	fmt.Println("authentication successful")

	return nil

	// accept validated data from frontend
	// check in database if users exists
	// if exists - approve and redirect with successfull message
	// if not - disapprove and show red modal
}
