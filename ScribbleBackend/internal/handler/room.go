package handler

import (
	"scribble-backend/internal/constants"

	"github.com/google/uuid"
)

func CreateRoom() {

	uuid := uuid.New().String()
	room := constants.NewRoom(uuid)
	hub.Rooms[uuid] = room
}
