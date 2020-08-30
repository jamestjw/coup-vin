package models

import "fmt"

type User struct {
	Password string
	Username string
	Name     string
	ID       int
}

var DefaultUsers = []User{
	{Name: "John Doe", Username: "john_doe", Password: "notasecret", ID: 1},
}

func FindUserByID(id int) (User, error) {
	for _, user := range DefaultUsers {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

func FindUserByUsername(username string) (User, error) {
	for _, user := range DefaultUsers {
		if user.Username == username {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}
