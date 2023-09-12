package regValidator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"websocket_1/server/registration"
)

type UserRegistration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegValidator(w http.ResponseWriter, r *http.Request) {
	var request UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := request.Username
	password := request.Password

	fmt.Println("username in validator.go: ", username)
	fmt.Println("password in validator.go: ", password)

	// if !isValid(username, password) {
	// 	http.Error(w, "Invalid input", http.StatusBadRequest)
	// 	return
	// }

	// Perform registration
	if err := registration.RegisterUser(username, password); err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	// Registration successful
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))
}

// func isValid(username, password string) bool {
// 	if len(password) > 8 {
// 		return true
// 	}
// 	return false
// }
