package main

import (
	"net/http"
	"os"
	"time"

	"gin-http-server/config"
	"gin-http-server/model"
	"gin-http-server/router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	setLogger()
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	server := &http.Server{
		Addr:    ":" + config.Setting.Server.Port,
		Handler: router.Handler(db),
		// for upload data
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info().Msg("Server is running")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("can't start server")
	}
}

func setLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Setting.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if config.Setting.Logs.Pretty {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: !config.Setting.Logs.Color,
			},
		)
	}
}
