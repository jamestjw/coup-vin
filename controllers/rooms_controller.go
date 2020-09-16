package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := server.DB.AllJoinableRooms()

	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}
	jsonResponse(w, http.StatusOK, rooms)
}

func (server *Server) GetRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
	}

	room, err := server.DB.FindRoomByID(uint(roomID))

	if err != nil {
		recordNotFoundResponse(w)
		return
	}
	jsonResponse(w, http.StatusOK, room)
}
