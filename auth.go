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
