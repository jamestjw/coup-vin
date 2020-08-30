package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jamestjw/coup-vin/auth"
	"github.com/jamestjw/coup-vin/middlewares"
	"github.com/jamestjw/coup-vin/models"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// Heartbeat API
	r.HandleFunc("/heartbeat", heartbeatHandler).Methods("GET")

	// Static content
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signin", signinHandler).Methods("POST")
	authRouter.HandleFunc("/refresh", refreshHandler).Methods("POST")
	authRouter.HandleFunc("/signup", signupHandler).Methods("POST")

	protectedRouter := r.PathPrefix("/").Subrouter()
	protectedRouter.Use(auth.Middleware)
	protectedRouter.HandleFunc("/rooms", roomsHandler).Methods("GET")
	protectedRouter.HandleFunc("/rooms/{id}/messages", addMessageHandler).Methods("POST")

	r.Use(middlewares.LoggingMiddleware)
	http.ListenAndServe(":8080", r)
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Implemented")
})

var heartbeatHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Je suis bien vivant!")
}

var roomsHandler = func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(models.DefaultRooms)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

var addMessageHandler = func(w http.ResponseWriter, r *http.Request) {
	// var message models.Message
	var room models.Room
	vars := mux.Vars(r)
	roomID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	for _, r := range models.DefaultRooms {
		if r.ID == roomID {
			room = r
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if room.ID != 0 {
		payload, _ := json.Marshal(room)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func init() {
	initialiseConfig()
	models.InitialiseDatabase()
}
