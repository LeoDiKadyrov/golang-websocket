package registration

import (
	"fmt"
	"log"

	postgresdb "websocket_1/server/database"
)

func RegisterUser(username, password string) error {
	fmt.Println("username in registration.go: ", username)
	fmt.Println("password in registration.go: ", password)

	hashedFinalPassword, finalSalt, err := HashAndSaltPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hashedFinalPassword)
	_, err = postgresdb.PostgresDB.Exec("INSERT INTO users (username, password_hash, salt) VALUES ($1, $2, $3)", username, hashedFinalPassword, finalSalt)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

/* TODO:
x Handle input values on js
x Validate them
x Send POST request to registration.go
x Put DB connection in another file for singleton

x User Struct
x Database Setup (ORM or GORM) for username + password
- Registration handler func
- Authentication handler func
- Logout handler func
- Protect Websocket Endpoint (check user's JWT before allowing connection)
- Session management through gorilla/sessions
- Create models folder with Message model and User model ???
- Rate Limiting and Brute Force Protection
- Secure Password Recovery
*/
