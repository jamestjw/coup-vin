package controllers

import (
	"net/http"
)

func (server *Server) Heartbeat(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, "Je suis bien vivant!")
}
