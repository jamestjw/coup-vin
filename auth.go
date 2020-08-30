package main

import (
	"encoding/json"
	"net/http"

	"github.com/jamestjw/coup-vin/auth"
	"github.com/jamestjw/coup-vin/models"
)

type signinParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type refreshParams struct {
	RefreshToken string `json:"refresh_token"`
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var creds signinParams
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := models.FindUserByUsername(creds.Username)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokens, err := auth.GenerateTokenPair(user)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, _ := json.Marshal(tokens)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	params := refreshParams{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil || params.RefreshToken == "" {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := auth.RefreshToken(params.RefreshToken)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, _ := json.Marshal(tokens)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
