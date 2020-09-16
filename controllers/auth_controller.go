package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jamestjw/coup-vin/auth"
)

type signinParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type signupParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type refreshParams struct {
	RefreshToken string `json:"refresh_token"`
}

func (server *Server) Signin(w http.ResponseWriter, r *http.Request) {
	var creds signinParams
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		badRequestResponse(w)
		return
	}

	user, err := server.DB.FindUserByUsername(creds.Username)

	if err != nil || !user.MatchesPassword(creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokens, err := auth.GenerateTokenPair(user)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(w, http.StatusOK, tokens)
}

func (server *Server) Signup(w http.ResponseWriter, r *http.Request) {
	var params signupParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		badRequestResponse(w)
		return
	}

	exists := server.DB.UsernameExists(params.Username)

	if exists {
		errorResponseWithMessage(w, http.StatusUnprocessableEntity, "Username is taken")
		return
	}

	user, err := server.DB.CreateUser(params.Username, params.Password)

	if err != nil {
		// TODO: Return error for failure
		errorResponseWithMessage(w, http.StatusUnprocessableEntity, "Failed to create user")
		return
	}

	tokens, err := auth.GenerateTokenPair(user)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		internalServerErrorResponse(w)
		return
	}

	jsonResponse(w, http.StatusOK, tokens)
}

func (server *Server) Refresh(w http.ResponseWriter, r *http.Request) {
	params := refreshParams{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil || params.RefreshToken == "" {
		// If the structure of the body is wrong, return an HTTP error
		badRequestResponse(w)
		return
	}

	tokens, err := auth.RefreshToken(params.RefreshToken, server.DB)

	if err != nil {
		errorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	jsonResponse(w, http.StatusOK, tokens)
}
