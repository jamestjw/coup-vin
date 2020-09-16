package models

import (
	"log"
	"testing"

	"github.com/jamestjw/coup-vin/app/models"
	"github.com/jamestjw/coup-vin/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByID(t *testing.T) {
	err := refreshTable(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := db.FindUserByID(user.ID)
	if err != nil {
		t.Errorf("error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Username, user.Username)
}

func TestFindUserByUsername(t *testing.T) {
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := db.FindUserByUsername(user.Username)
	if err != nil {
		t.Errorf("error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Username, user.Username)
}

func TestUsernameExists(t *testing.T) {
	users, err := seedUsers()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	type example struct {
		inputUsername string
		exists        bool
	}

	var examples []example

	// Add in usernames that exist
	for _, user := range users {
		examples = append(examples, example{inputUsername: user.Username, exists: true})
	}
	// Add in username that does not exist
	randomName, err := utils.RandomHex(20)
	if err != nil {
		log.Fatalf("cannot generate random name: %v", err)
	}
	examples = append(examples, example{inputUsername: randomName, exists: false})

	for _, example := range examples {
		res := db.UsernameExists(example.inputUsername)
		assert.Equal(t, res, example.exists)
	}
}
func TestCreateUser(t *testing.T) {

	err := refreshTable(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	randomName, err := utils.RandomHex(20)
	if err != nil {
		log.Fatalf("cannot generate random name: %v", err)
	}

	user, err := db.CreateUser(randomName, "password")
	if err != nil {
		t.Errorf("error creating new user: %v\n", err)
	}
	// Check that ID was assigned
	assert.Greater(t, user.ID, uint(0))
}

// Helper functions for user test

func seedOneUser() (models.User, error) {

	refreshTable(&models.User{})

	user := models.User{
		Username: "test_username",
		Password: "password",
	}

	err := db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() ([]*models.User, error) {
	refreshTable(&models.User{})

	users := []*models.User{
		&models.User{
			Username: "test_username1",
			Password: "password",
		},
		&models.User{
			Username: "test_username2",
			Password: "password",
		},
	}

	for _, user := range users {
		err := db.Model(&models.User{}).Create(user).Error
		if err != nil {
			return users, err
		}
	}
	return users, nil
}
