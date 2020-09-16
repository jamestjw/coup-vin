package models

import (
	"fmt"
	"log"
	"testing"

	"github.com/jamestjw/coup-vin/app/models"
	"github.com/stretchr/testify/assert"
)

func TestAllJoinableRooms(t *testing.T) {
	rooms, err := seedRooms()
	if err != nil {
		log.Fatal(err)
	}

	foundRooms, err := db.AllJoinableRooms()
	if err != nil {
		t.Fatal("unable to fetch all joinable rooms")
	}
	assert.Equal(t, len(foundRooms), len(rooms))
}

func TestFindRoomByID(t *testing.T) {
	rooms, err := seedRooms()
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		fmt.Println("searching for room...")
		fmt.Println(room)
		foundRoom, err := db.FindRoomByID(room.ID)
		if err != nil {
			t.Fatalf("unable to fetch room by ID: %v", err)
		}
		assert.Equal(t, foundRoom.ID, room.ID)
		assert.Equal(t, foundRoom.Name, room.Name)
	}
}

func seedRooms() ([]*models.Room, error) {
	err := refreshTable(&models.Room{})
	if err != nil {
		log.Fatal(err)
	}

	rooms := []*models.Room{
		&models.Room{
			Name: "room_1",
		},
		&models.Room{
			Name: "room_2",
		},
	}

	for _, room := range rooms {
		err := db.Model(&models.Room{}).Create(room).Error
		if err != nil {
			return nil, err
		}
	}
	return rooms, nil
}
