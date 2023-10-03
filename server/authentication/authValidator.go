package authentication

import (
	"fmt"
	"net/http"
	"encoding/json"

	"websocket_1/server/models"
	"websocket_1/server/validation"
)

func AuthValidator(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I'm an auth validator")

	var request models.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := request.Username
	password := request.PasswordHash

	if !validation.IsValidPassword(password) || !validation.IsValidUsername(username) {
		http.Error(w, "Invalid authentication input", http.StatusBadRequest)
		return
	}

	if err := AuthenticateUser(username, password); err != nil {
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Authentication successful"))
}
