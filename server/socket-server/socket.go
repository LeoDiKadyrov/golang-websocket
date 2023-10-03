package socket

// WE"RE OFFICIALY IN DEV
import (
	"fmt"
	"log"
	"net/http"

	"websocket_1/server/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading JSON: %v", err)
			delete(clients, conn)
			return
		}

		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("err writing message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	http.ServeFile(w, r, "client/chat.html")
}
