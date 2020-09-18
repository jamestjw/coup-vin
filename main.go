package main

import (
	"fmt"
	"net/http"

	"github.com/jamestjw/coup-vin/config"
	"github.com/jamestjw/coup-vin/models"
)

func main() {
	runServer()
	// authRouter := r.PathPrefix("/auth").Subrouter()
	// authRouter.HandleFunc("/signin", signinHandler).Methods("POST")
	// authRouter.HandleFunc("/refresh", refreshHandler).Methods("POST")
	// authRouter.HandleFunc("/signup", signupHandler).Methods("POST")

	// protectedRouter := r.PathPrefix("/").Subrouter()
	// protectedRouter.Use(auth.Middleware)
	// protectedRouter.HandleFunc("/rooms", roomsHandler).Methods("GET")
	// protectedRouter.HandleFunc("/rooms/{id}/messages", addMessageHandler).Methods("POST")

	// r.Use(middlewares.LoggingMiddleware)
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Implemented")
})

func init() {
	config.InitialiseConfig()
	models.InitialiseDatabase("production")
}
