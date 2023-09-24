package registration

import (
	"fmt"
	"log"

	postgresdb "websocket_1/server/database"
)

func RegisterUser(username, password string) error {
	db := postgresdb.GetInstanceDB().DB

	fmt.Println("username in registration.go: ", username)
	fmt.Println("password in registration.go: ", password)
	fmt.Println("DB: ", db)

	hashedFinalPassword, finalSalt, err := HashAndSaltPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hashedFinalPassword)
	_, err = db.Exec("INSERT INTO users (username, password_hash, salt) VALUES ($1, $2, $3)", username, hashedFinalPassword, finalSalt)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

/* TODO:
- Send registration successfull message to frontend and show it from javascript to users
- Redirect to chat main page (right now it's "/")
- Authentication handler func
- Forgot password handler
- Logout handler func
- Protect Websocket Endpoint (check user's JWT before allowing connection)
- Session management through gorilla/sessions
- Create models folder with Message model and User model ???
- Rate Limiting and Brute Force Protection
- Secure Password Recovery

x Handle input values on js
x Validate them
x Send POST request to registration.go
x Put DB connection in another file for singleton
x User Struct
x Database Setup (ORM or GORM) for username + password
x Registration handler func
*/
