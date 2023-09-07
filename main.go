package main

import (
	"fmt"
	"log"
	"net/http"

	"websocket_1/server/registration"
	"websocket_1/server/socket-server"
)

func main() {
	fs := http.FileServer(http.Dir("client/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", socket.GetRoot)
	http.HandleFunc("/ws", socket.WebsocketHandler)
	go socket.HandleMessages()

	fmt.Println(registration.Registration("ABDRA SUKA"))

	port := "8080"
	fmt.Printf("Server started on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
