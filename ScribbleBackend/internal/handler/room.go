package handler

import (
	"log"
	"net/http"
	"scribble-backend/internal/constants"
	"scribble-backend/pkg/websocket"

	"github.com/google/uuid"
)

type RoomHandler struct {
	Hub *websocket.Hub
}

func (r *RoomHandler) CreateRoom(w http.ResponseWriter, req *http.Request) {
	uuid := uuid.New().String()
	room := constants.NewRoom(uuid)
	r.Hub.Rooms[uuid] = room
	log.Printf("new room created with id %s", room.ID)
	w.Write([]byte(`{"roomId":"` + room.ID + `"}`))
}

func (r *RoomHandler) DeleteRoom(w http.ResponseWriter, req *http.Request) {
	roomId := req.URL.Query().Get("roomId")
	_, exist := r.Hub.GetRoomFromId(roomId)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status":"room not found"}`))
		return
	}
	delete(r.Hub.Rooms, roomId)
	w.Write([]byte(`{"status":"ok"}`))
}
