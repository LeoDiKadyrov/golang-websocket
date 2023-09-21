package main

import (
	"fmt"
	"log"
	"net/http"

	"websocket_1/server/registration"
	"websocket_1/server/socket-server"
	"websocket_1/server/database"
)

func main() {
	postgresdb.Postgresqdb()
	defer func() {
        if err := postgresdb.PostgresDB.Close(); err != nil {
            log.Println("Error closing database connection:", err)
        } else {
            log.Println("Database connection closed.")
        }
    }()

	fs := http.FileServer(http.Dir("client/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", socket.GetRoot)
	http.HandleFunc("/ws", socket.WebsocketHandler)
	http.HandleFunc("/register", registration.RegValidator)
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/registration.html")
	})

	go socket.HandleMessages()

	port := "8080"
	fmt.Printf("Server started on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
