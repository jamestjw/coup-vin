package models

import (
	"log"
	"os"
	"testing"

	"github.com/jamestjw/coup-vin/config"
	"github.com/spf13/viper"
)

var db *DB

func initialiseTestDatabase() {
	var err error
	db, err = InitialiseDatabase("test")
	if err != nil {
		log.Fatalf("Error loading database %v\n", err)
	}
}

func TestMain(m *testing.M) {
	initialiseConfig()
	initialiseTestDatabase()
	os.Exit(m.Run())
}

// refreshTable will drop a table and auto migrate it (which will then recreate it)
// Pass in a pointer to an instance of a model, e.g. refreshTable(&User{})
func refreshTable(table interface{}) error {
	// TODO: Consider just wiping table instead of dropping and re-creating
	err := db.Migrator().DropTable(table)
	if err != nil {
		return err
	}
	err = db.AutoMigrate(table)
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func initialiseConfig() {
	viper.AddConfigPath("../config")
	config.InitialiseConfig()
}
