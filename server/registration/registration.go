package registration

import (
	"log"

	postgresdb "websocket_1/server/database"
	"websocket_1/server/security"
)

func RegisterUser(username, password string) error {
	db := postgresdb.GetInstanceDB().DB

	hashedFinalPassword, err := security.HashAndSaltPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hashedFinalPassword)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}