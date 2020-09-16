package controllers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jamestjw/coup-vin/models"
)

type Server struct {
	Router *mux.Router
	DB     models.Datastore
}

// Initialize the server object given a particular env (production or test)
func (server *Server) Initialize(env string) {
	var err error
	server.DB, err = models.InitialiseDatabase(env)
	if err != nil {
		log.Fatalf("Error loading database %v\n", err)
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)

	log.Info(fmt.Sprintf("Listening on addr: %s", addr))
	log.Fatal(http.ListenAndServe(addr, cors(server.Router)))
}
