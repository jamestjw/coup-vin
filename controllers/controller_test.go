package controllers

import (
	"os"
	"testing"

	"github.com/jamestjw/coup-vin/config"
	"github.com/spf13/viper"
)

var server = Server{}

func TestMain(m *testing.M) {
	initialiseConfig()
	os.Exit(m.Run())
}

func initialiseConfig() {
	viper.AddConfigPath("../config")
	config.InitialiseConfig()
}
