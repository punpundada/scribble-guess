package handler

import (
	"log"
	"net/http"

	ws "scribble-backend/pkg/websocket"

	"github.com/gorilla/websocket"
)

var hub = ws.NewHub()

func init() {
	go hub.Run()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	hub.Clients[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(hub.Clients, conn)
			conn.Close()
			break
		}

		hub.Broadcast <- msg
	}
}
