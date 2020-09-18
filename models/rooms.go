package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Password  string         `json:"-"`
	MinUsers  int            `json:"min_users"`
	MaxUsers  int            `json:"max_users"`
	Name      string         `json:"name"`
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

	return &room, err
}
