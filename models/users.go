package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string
	Username string `gorm:"uniqueIndex"`
}

var DefaultUsers = []User{
	{Username: "john_doe", Password: "notasecret"},
}

func FindUserByID(id int) *User {
	user := User{}
	DB.First(&user, "id = ?", id)
	if user == (User{}) {
		return nil
	}
	return &user
}

func FindUserByUsername(username string) *User {
	user := User{}
	DB.First(&user, "username = ?", username)
	if user == (User{}) {
		return nil
	}
	return &user
}

func UsernameExists(username string) (exists bool) {
	user := FindUserByUsername(username)
	return user != nil
}

func CreateUser(username string, password string) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Password: string(encryptedPassword),
	}

	res := DB.Create(user)
	err = res.Error
	return user, err
}

func (u *User) MatchesPassword(p string) bool {
	// Passwords match if there is no error
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) == nil
}
