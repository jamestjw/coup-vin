package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitialiseDatabase() {
	newLogger := logger.New(
		log.New(), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	var err error

	dbDirectory := viper.GetString("database.directory")
	DB, err = gorm.Open(sqlite.Open(dbDirectory), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&User{}, &Room{})

	log.Info("Migrated DB schemas.")
}
