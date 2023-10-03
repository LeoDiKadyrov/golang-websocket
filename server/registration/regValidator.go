package registration

import (
	"encoding/json"
	"net/http"

	"websocket_1/server/models"
	"websocket_1/server/validation"
)

func RegValidator(w http.ResponseWriter, r *http.Request) {
	var request models.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := request.Username
	password := request.PasswordHash

	if !validation.IsValidPassword(password) || !validation.IsValidUsername(username) {
		http.Error(w, "Invalid registration input", http.StatusBadRequest)
		return
	}

	// Perform registration
	if err := RegisterUser(username, password); err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	// Registration successful // TODO: PUT THIS INTO REGISTRATION.GO
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))
}
