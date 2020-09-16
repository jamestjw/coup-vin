package models

import (
	"errors"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Password string
	MinUsers int
	MaxUsers int
	Name     string
}

func (db *DB) AllJoinableRooms() ([]Room, error) {
	rooms := new([]Room)
	err := db.Find(rooms).Error
	if err != nil {
		return []Room{}, err
	}
	return *rooms, nil
}

func (db *DB) FindRoomByID(id uint) (*Room, error) {
	room := Room{}
	err := db.First(&room, "id = ?", id).Error
	if err != nil {
		return &Room{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Room{}, errors.New("Room Not Found")
	}
	return &room, err
}
