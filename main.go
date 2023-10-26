package main

import (
	"fmt"
	"log"
	"net/http"

	"websocket_1/server/authentication"
	postgresdb "websocket_1/server/database"
	"websocket_1/server/registration"
	"websocket_1/server/socket-server"
)

func main() {
	postgresdb.GetInstanceDB()
	defer postgresdb.Close()

	fs := http.FileServer(http.Dir("client/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", socket.GetRoot)
	http.HandleFunc("/ws", socket.WebsocketHandler)
	http.HandleFunc("/register", registration.RegHandler)
	http.HandleFunc("/authenticate", authentication.AuthHandler)
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/registration.html")
	})
	http.HandleFunc("/authentication", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/authentication.html")
	})
	http.HandleFunc("/recovery", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/recoverPassword.html")
	})
	http.HandleFunc("/newpassword", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/newPassword.html")
	})

	go socket.HandleMessages()

	port := "8080"
	fmt.Printf("Server started on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
