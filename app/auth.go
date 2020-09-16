package main

type signinParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type refreshParams struct {
	RefreshToken string `json:"refresh_token"`
}

type signupParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// func signinHandler(w http.ResponseWriter, r *http.Request) {
// 	var creds signinParams
// 	// Get the JSON body and decode into credentials
// 	err := json.NewDecoder(r.Body).Decode(&creds)
// 	if err != nil {
// 		// If the structure of the body is wrong, return an HTTP error
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	user := models.FindUserByUsername(creds.Username)

// 	if user == nil || !user.MatchesPassword(creds.Password) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	tokens, err := auth.GenerateTokenPair(user)

// 	if err != nil {
// 		// If there is an error in creating the JWT return an internal server error
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	payload, _ := json.Marshal(tokens)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(payload))
// }

// func refreshHandler(w http.ResponseWriter, r *http.Request) {
// 	params := refreshParams{}
// 	err := json.NewDecoder(r.Body).Decode(&params)
// 	if err != nil || params.RefreshToken == "" {
// 		// If the structure of the body is wrong, return an HTTP error
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	tokens, err := auth.RefreshToken(params.RefreshToken)

// 	if err != nil {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		// TODO: Write a helper to handle errors
// 		writeJSONResponse(w, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSONResponse(w, tokens)
// }

// func signupHandler(w http.ResponseWriter, r *http.Request) {
// 	var params signupParams
// 	err := json.NewDecoder(r.Body).Decode(&params)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	exists := models.UsernameExists(params.Username)

// 	if exists {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		writeJSONResponse(w, map[string]string{"error": "Username is taken"})
// 		return
// 	}

// 	user, err := models.CreateUser(params.Username, params.Password)

// 	if err != nil {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		writeJSONResponse(w, map[string]string{"error": "Failed to create user"})
// 		return
// 	}

// 	tokens, err := auth.GenerateTokenPair(user)

// 	if err != nil {
// 		// If there is an error in creating the JWT return an internal server error
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	writeJSONResponse(w, tokens)
// }
