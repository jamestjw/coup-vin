package controllers

// func seedOneUser(username string, password string) (*models.User, error) {

// 	err := refreshTable(&models.User{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	user, err := server.DB.CreateUser(username, password)
// 	return user, err
// }

// func refreshTable(table interface{}) error {
// 	// TODO: Consider just wiping table instead of dropping and re-creating
// 	err := server.DB.Migrator().DropTable(table)
// 	if err != nil {
// 		return err
// 	}
// 	err = server.DB.AutoMigrate(table)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Successfully refreshed table")
// 	return nil
// }

// func TestSigninHandler(t *testing.T) {
// 	username := "testUsername"
// 	password := "password"

// 	user, err := seedOneUser(username, password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	examples := []struct {
// 		inputJSON  string
// 		statusCode int
// 	}{
// 		{
// 			inputJSON:  `{"username": "testUsername", "password": "password"}`,
// 			statusCode: 200,
// 		},
// 		{
// 			inputJSON:  `{"username": "testUsername", "password": "wrong password"}`,
// 			statusCode: 401,
// 		},
// 		{
// 			inputJSON:  `{"username": "wrong username", "password": "password"}`,
// 			statusCode: 401,
// 		},
// 	}

// 	for _, v := range examples {
// 		req, err := http.NewRequest("POST", "/auth/signin", bytes.NewBufferString(v.inputJSON))
// 		if err != nil {
// 			t.Errorf("request failed with error: %v", err)
// 		}

// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.Signin)
// 		handler.ServeHTTP(rr, req)

// 		assert.Equal(t, rr.Code, v.statusCode)
// 		if v.statusCode == 200 {
// 			assert.NotEqual(t, rr.Body.String(), "")
// 		}
// 	}
// }
