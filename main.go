package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jamestjw/coup-vin/app/middlewares"
	"github.com/jamestjw/coup-vin/app/models"
	"github.com/jamestjw/coup-vin/config"
)

func main() {
	r := mux.NewRouter()

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)

	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// Heartbeat API
	r.HandleFunc("/heartbeat", heartbeatHandler).Methods("GET")

	// // Static content
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// authRouter := r.PathPrefix("/auth").Subrouter()
	// authRouter.HandleFunc("/signin", signinHandler).Methods("POST")
	// authRouter.HandleFunc("/refresh", refreshHandler).Methods("POST")
	// authRouter.HandleFunc("/signup", signupHandler).Methods("POST")

	// protectedRouter := r.PathPrefix("/").Subrouter()
	// protectedRouter.Use(auth.Middleware)
	// protectedRouter.HandleFunc("/rooms", roomsHandler).Methods("GET")
	// protectedRouter.HandleFunc("/rooms/{id}/messages", addMessageHandler).Methods("POST")

	r.Use(middlewares.LoggingMiddleware)
	http.ListenAndServe(":8080", cors(r))
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Implemented")
})

var heartbeatHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Je suis bien vivant!")
}

// var roomsHandler = func(w http.ResponseWriter, r *http.Request) {
// 	rooms, err := models.AllJoinableRooms()

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	payload, _ := json.Marshal(rooms)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(payload))
// }

// var addMessageHandler = func(w http.ResponseWriter, r *http.Request) {
// 	// var message models.Message
// 	vars := mux.Vars(r)
// 	roomID, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	room := models.FindRoomByID(roomID)

// 	w.Header().Set("Content-Type", "application/json")
// 	if room != nil {
// 		payload, _ := json.Marshal(room)
// 		w.Write([]byte(payload))
// 	} else {
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

func init() {
	config.InitialiseConfig()
	models.InitialiseDatabase("production")
}
