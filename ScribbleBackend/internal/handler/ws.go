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
	log.Println("roomId", roomId)
	room := ws.Hub.GetOrCreateRoom(roomId)
	clientId := uuid.New().String()

	log.Printf("Client %s is trying to join room %s", clientId, roomId)

	client := &constants.Client{
		Conn:   conn,
		Send:   make(chan *constants.BroadcastMessage, 1024),
		RoomID: room.ID,
		Id:     clientId,
	}
	room.Register <- client

	room.BroadcastAll <- constants.NewBroadcastMessage("message", &constants.Meta{
		SenderId: client.Id,
		RoomId:   room.ID,
		Time:     time.Now().Unix(),
	}, "User "+client.Id+" has joined the room")

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

	const maxMessageSize = 1024 * 16 // 8KB limit
	c.Conn.SetReadLimit(maxMessageSize)
	t := time.Now().Add(10 * time.Minute)
	c.Conn.SetReadDeadline(t)
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(t)
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
		meta := constants.Meta{
			SenderId: c.Id,
			RoomId:   c.RoomID,
			Time:     time.Now().Unix(),
		}
		room.Broadcast <- constants.NewBroadcastMessage("message", &meta, string(raw))
	}
}

func writePump(c *constants.Client) {
	for msg := range c.Send {
		err := c.Conn.WriteJSON(msg)
		// err := c.Conn.(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
