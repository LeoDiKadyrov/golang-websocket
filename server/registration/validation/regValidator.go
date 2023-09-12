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

	err := registration.RegisterUser(username, password)
	if err != nil {
		http.Error(w, "Failed to send username + password", http.StatusInternalServerError)
		return
	}
}
