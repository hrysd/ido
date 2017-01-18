package main

import (
	"github.com/hrysd/ido/internal/idobata"
)

func Post(content string) {
	room := idobata.GetRooms("esm", "hrysd").Rooms[0]

	idobata.Post("hi from golang", room.Id)
}
