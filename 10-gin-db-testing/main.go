package main

import (
	"net/http"
	"os"
	"time"

	"gin-http-server/config"
	"gin-http-server/model"
	"gin-http-server/router"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	setLogger()
	if err := model.SetupDatabase(); err != nil {
		log.Fatal().Err(err).Msg("can't create db")
	}

	model.CreateUser(&model.User{
		Email: "testmodel@gmail.com",
		Name:  "Person02",
	})

	server := &http.Server{
		Addr:    ":" + config.Setting.Server.Port,
		Handler: router.Handler(),
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
