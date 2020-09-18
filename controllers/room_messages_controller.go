package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateMessageForRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID, err := strconv.Atoi(vars["id"])

	// TODO: Get message params and throw error

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	room, err := server.DB.FindRoomByID(uint(roomID))

	if err != nil {
		recordNotFoundResponse(w)
		return
	}

	// TODO: Add message to room

	jsonResponse(w, http.StatusOK, room)
}
