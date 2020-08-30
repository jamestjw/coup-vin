package models

type Room struct {
	Password string
	MinUsers int
	MaxUsers int
	Name     string
	ID       int
}

var DefaultRooms = []Room{
	{Name: "World of Authcraft", MinUsers: 3, MaxUsers: 6, ID: 1},
	{Name: "Ocean Explorer", MinUsers: 3, MaxUsers: 6, ID: 2},
	{Name: "Dinosaur Park", MinUsers: 3, MaxUsers: 6, ID: 3},
	{Name: "Cars VR", MinUsers: 3, MaxUsers: 6, ID: 4},
	{Name: "Robin Hood", MinUsers: 3, MaxUsers: 6, ID: 5},
	{Name: "Real World VR", MinUsers: 3, MaxUsers: 6, ID: 6},
}
