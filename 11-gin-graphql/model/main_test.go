package model

import (
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB connection
var DB *gorm.DB
var err error

// SetupDatabase connection
func SetupDatabase() error {
	// github.com/mattn/go-sqlite3
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	sqlDB, err := DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	initialDB()

	return err
}

func initialDB() {
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal().Err(err).Msg("can't AutoMigrate db")
	}
}
