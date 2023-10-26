package recoverpassword

import (
	"encoding/json"
	"log"
	"net/http"
	"websocket_1/server/models"
	"websocket_1/server/validation"
	postgresdb "websocket_1/server/database"
)

func recoverPassword(w http.ResponseWriter, r *http.Request) {
	db := postgresdb.GetInstanceDB().DB
	var request models.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := request.Username

	if !validation.IsValidUsername(username) {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		log.Println("Error checking existence:", err)
		return
	}

	if exists {
		err := db.QueryRow("UPDATE TABLE users SET password_hash = '' WHERE username = $1", username)
		if err != nil {
			log.Println("Error updating password:", err)
			return
		}

		// TODO:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Password is reset"))
	}
	
}
