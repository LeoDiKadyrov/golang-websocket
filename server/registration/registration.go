package registration

import (
	"encoding/json"
	"log"
	"net/http"

	postgresdb "websocket_1/server/database"
	"websocket_1/server/models"
	"websocket_1/server/security"
	"websocket_1/server/validation"
)

func RegHandler(w http.ResponseWriter, r *http.Request) {
	db := postgresdb.GetInstanceDB().DB
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

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		log.Println("Error checking existence:", err)
		return
	}

	if exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	} 

	hashedFinalPassword, err := security.HashAndSaltPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hashedFinalPassword)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))
}
