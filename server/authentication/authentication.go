package authentication

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	postgresdb "websocket_1/server/database"
	customError "websocket_1/server/error"
	"websocket_1/server/models"
	"websocket_1/server/validation"

	"golang.org/x/crypto/bcrypt"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := authenticateUser(username, password); err != nil {
		if err.Error() == "password is incorrect" {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			customError.SendCustomHttpError(w, http.StatusInternalServerError, "Authentication failed")
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Authentication successful"))
}

func authenticateUser(username, password string) error {
	var user models.User
	db := postgresdb.GetInstanceDB().DB

	stmt, err := db.Prepare("SELECT username, password_hash FROM users WHERE username = $1;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}
