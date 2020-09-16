package models

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Datastore interface {
}

type DB struct {
	*gorm.DB
}

// InitialiseDatabase creates a database connection based on env
// env: production or test
func InitialiseDatabase(env string) (*DB, error) {
	newLogger := logger.New(
		log.New(), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	var err error

	dbDirectoryKey := fmt.Sprintf("database.%s.directory", env)
	dbDirectory := viper.GetString(dbDirectoryKey)

	log.Info(dbDirectoryKey)
	log.Info(fmt.Sprintf("Loading DB file from directory: %s for env: %s.", dbDirectory, env))

	db, err := gorm.Open(sqlite.Open(dbDirectory), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Room{})

	log.Info(fmt.Sprintf("Migrated DB schemas for %s env.", env))

	return &DB{db}, nil
}
