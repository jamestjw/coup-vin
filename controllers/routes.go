package controllers

import "net/http"

func (s *Server) initializeRoutes() {
	// Serve static content
	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Heartbeat API
	s.Router.HandleFunc("/heartbeat", s.Heartbeat).Methods("GET")

	// Rooms
	s.Router.HandleFunc("/rooms", s.GetRooms).Methods("GET")
	s.Router.HandleFunc("/rooms/{id}", s.GetRoom).Methods("GET")

	// Room Messages
	s.Router.HandleFunc("/rooms/{id}/messages", s.CreateMessageForRoom).Methods("POST")

	// Authentication
	authRouter := s.Router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signin", s.Signin).Methods("POST")
	authRouter.HandleFunc("/signup", s.Signup).Methods("POST")
	authRouter.HandleFunc("/refresh", s.Refresh).Methods("POST")
}
