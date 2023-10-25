package authentication

import (
	"database/sql"
	"errors"
	"log"

	postgresdb "websocket_1/server/database"
	"websocket_1/server/models"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(username, password string) error {
	var user models.User
	db := postgresdb.GetInstanceDB().DB

	stmt, err := db.Prepare("SELECT username, password_hash FROM users WHERE username = $1;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
