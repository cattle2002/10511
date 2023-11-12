package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		message := string(p)
		fmt.Printf("Received message: %s\n", message)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main22() {
	http.HandleFunc("/api/websocket/1", func(w http.ResponseWriter, r *http.Request) {
		handleConnection(w, r)
	})

	fmt.Println("WebSocket server is running on :8080")
	http.ListenAndServe(":9999", nil)
}
