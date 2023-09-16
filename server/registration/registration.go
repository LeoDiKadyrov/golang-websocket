package registration

import (
	"fmt"
)

func RegisterUser(username, password string) error {
	fmt.Println("username in registration.go: ", username)
	fmt.Println("password in registration.go: ", password)
	return nil
}

/* TODO:
x Handle input values on js
x Validate them
x Send POST request to registration.go
- Make http server and handler of POST requests here
- Put DB connection in another file for singleton

x User Struct
x Database Setup (ORM or GORM) for username + password
- Registration handler func
- Authentication handler func
- Logout handler func
- Protect Websocket Endpoint (check user's JWT before allowing connection)
- Session management through gorilla/sessions
- Create models folder with Message model and User model ???
*/
