package models

type User struct {
	Username string `json:"username"`
	PasswordHash string `json:"password"`
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}