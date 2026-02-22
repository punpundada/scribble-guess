package handler

import (
	"log"
	"net/http"
	"time"

	"scribble-backend/internal/constants"
	ws "scribble-backend/pkg/websocket"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WS struct {
	Hub *ws.Hub
}

// func init() {
// 	go hub.Run()
// }

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (ws *WS) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	roomId := r.URL.Query().Get("roomId")

	room := ws.Hub.GetOrCreateRoom(roomId)
	clientId := uuid.New().String()

	log.Printf("Client %s is trying to join room %s", clientId, roomId)

	client := &constants.Client{
		Conn:   conn,
		Send:   make(chan []byte, 1024),
		RoomID: room.ID,
		Id:     clientId,
	}
	room.Register <- client
	room.BroadcastAll <- constants.NewBroadcastMessage(
		[]byte(`{"message": "A new user has joined the room", "roomId": "`+room.ID+`"}`),
		client.Id)

	go writePump(client)
	readPump(room, client)
}

func readPump(room *constants.Room, c *constants.Client) {

	defer func() {
		if recover() != nil {
			log.Printf("Recovered from panic in readPump for client %s in room %s", c.Id, c.RoomID)
		}
		room.Unregister <- c
		c.Conn.Close()
	}()

	const maxMessageSize = 1024 * 8 // 8KB limit
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(10 * time.Minute))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(10 * time.Minute))
		return nil
	})

	for {
		_, raw, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error for client %s in room %s: %v", c.Id, c.RoomID, err)
			}
			break
		}
		log.Printf("client %s in room %s sent message: %s", c.Id, c.RoomID, string(raw))
		room.Broadcast <- constants.NewBroadcastMessage(raw, c.Id)
	}
}

func writePump(c *constants.Client) {
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
