package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "github.com/jamestjw/coup-vin/mocks/mock_database"
	"github.com/jamestjw/coup-vin/models"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedUsername := "testUsername"
	expectedPassword := "password"

	encrypedPassword, err := bcrypt.GenerateFromPassword([]byte(expectedPassword), 10)
	if err != nil {
		log.Fatal(err)
	}

	expectedUser := &models.User{
		Username: expectedUsername,
		Password: string(encrypedPassword),
	}

	examples := []struct {
		inputUsername string
		inputJSON     string
		statusCode    int
	}{
		{
			inputUsername: "testUsername",
			inputJSON:     `{"username": "testUsername", "password": "password"}`,
			statusCode:    200,
		},
		{
			inputUsername: "testUsername",
			inputJSON:     `{"username": "testUsername", "password": "wrong password"}`,
			statusCode:    401,
		},
		{
			inputUsername: "wrong username",
			inputJSON:     `{"username": "wrong username", "password": "password"}`,
			statusCode:    401,
		},
	}

	for _, example := range examples {
		mockDatabase := mock_database.NewMockDatastore(ctrl)
		if example.inputUsername == expectedUsername {
			mockDatabase.EXPECT().FindUserByUsername(example.inputUsername).Return(expectedUser, nil)
		} else {
			mockDatabase.EXPECT().FindUserByUsername(example.inputUsername).Return(&models.User{}, nil)
		}
		server.DB = mockDatabase

		req, err := http.NewRequest("POST", "/auth/signin", bytes.NewBufferString(example.inputJSON))
		if err != nil {
			t.Errorf("request failed with error: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.Signin)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, example.statusCode)
		if example.statusCode == 200 {
			assert.NotEqual(t, rr.Body.String(), "")
		}
	}
}
