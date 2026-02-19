package handler

import (
	"log"
	"net/http"

	"scribble-backend/internal/constants"
	ws "scribble-backend/pkg/websocket"

	"github.com/google/uuid"
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

	uuid := uuid.New().String()
	hub.Rooms[uuid] = constants.NewRoom(uuid)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			room, exists := hub.GetRoomFromId(uuid)
			if exists {
				room.Mu.Lock()
				delete(room.Clients, conn)
				room.Mu.Unlock()
			}
			conn.Close()
			break
		}

		hub.Broadcast <- msg
	}
}
