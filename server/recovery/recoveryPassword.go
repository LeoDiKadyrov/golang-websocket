package recovery

import (
	"encoding/json"
	"log"
	"net/http"
	postgresdb "websocket_1/server/database"
	"websocket_1/server/models"
	"websocket_1/server/validation"
)

func RecoverPassword(w http.ResponseWriter, r *http.Request) {
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
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists) // put into model
	if err != nil {
		log.Println("Error checking existence:", err)
		return
	}

	if exists {
		_, err := db.Exec("UPDATE users SET password_hash = '' WHERE username = $1", username) // put into model
		if err != nil {
			log.Println("Error updating password:", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Password is reset"))

	} else {
		http.Error(w, "Username doesn't exist", http.StatusBadRequest)
		return
	}
}
