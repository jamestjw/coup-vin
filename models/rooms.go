package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Password string
	MinUsers int
	MaxUsers int
	Name     string
}

func AllJoinableRooms() (*[]Room, error) {
	rooms := new([]Room)
	result := DB.Find(rooms)
	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}
	return rooms, nil
}

func FindRoomByID(id int) *Room {
	room := Room{}
	DB.First(&room, "id = ?", id)
	if room == (Room{}) {
		return nil
	}
	return &room
}
