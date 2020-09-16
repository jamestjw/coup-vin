package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Password  string         `json:"-"`
	Username  string         `gorm:"uniqueIndex" json:"username"`
}

func (db *DB) FindUserByID(id uint) (*User, error) {
	u := &User{}
	err := db.First(u, "id = ?", id).Error
	if err != nil {
		return &User{}, err
	}

	return u, err
}

func (db *DB) FindUserByUsername(username string) (*User, error) {
	u := &User{}
	err := db.First(u, "username = ?", username).Error
	if err != nil {
		return &User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (db *DB) UsernameExists(username string) (exists bool) {
	_, err := db.FindUserByUsername(username)
	return err == nil
}

func (db *DB) CreateUser(username string, password string) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Password: string(encryptedPassword),
	}

	err = db.Save(user).Error

	return user, err
}

func (u *User) MatchesPassword(p string) bool {
	// Passwords match if there is no error
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) == nil
}
