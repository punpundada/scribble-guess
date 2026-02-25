package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"scribble-backend/pkg/websocket"
)

type RoomHandler struct {
	Hub *websocket.Hub
}

func (r *RoomHandler) CreateRoom(w http.ResponseWriter, req *http.Request) {
	room := r.Hub.GetOrCreateRoom("")
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

func (r *RoomHandler) GetRandomRoomId(w http.ResponseWriter, req *http.Request) {
	rommLen := len(r.Hub.Rooms)
	if rommLen == 0 {
		room := r.Hub.GetOrCreateRoom("")
		w.Write([]byte(`{"roomId":"` + room.ID + `"}`))
		return
	}
	keys := make([]string, 0, len(r.Hub.Rooms))
	for k := range r.Hub.Rooms {
		keys = append(keys, k)
	}
	randomKey := keys[rand.Intn(len(keys))]
	json.NewEncoder(w).Encode(map[string]string{
		"roomId": randomKey,
	})
}
