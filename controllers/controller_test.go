package controllers

import (
	"os"
	"testing"

	"github.com/jamestjw/coup-vin/config"
	"github.com/jamestjw/coup-vin/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var server = Server{}

func TestMain(m *testing.M) {
	initialiseConfig()
	initialiseTestDatabase()
	os.Exit(m.Run())
}

func initialiseTestDatabase() {
	var err error
	server.DB, err = models.InitialiseDatabase("test")
	if err != nil {
		log.Fatalf("Error loading database %v\n", err)
	}
}

func initialiseConfig() {
	viper.AddConfigPath("../config")
	config.InitialiseConfig()
}
