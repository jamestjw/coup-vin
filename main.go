package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jamestjw/coup-vin/models"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.Handle("/heartbeat", heartbeatHandler).Methods("GET")
	r.Handle("/rooms", roomsHandler).Methods("GET")
	r.Handle("/rooms/{id}/messages", addMessageHandler).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", r)
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var heartbeatHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Je suis bien vivant!"))
})

var roomsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(models.DefaultRooms)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var addMessageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
})
