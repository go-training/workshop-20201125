package main

import (
	"net/http"
	"os"
	"time"

	"gin-http-server/config"
	"gin-http-server/router"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
}

func main() {
	config.Load()
	setLogger()
	hander := router.Handler()

	s := &http.Server{
		Addr:         ":" + config.Setting.Server.Port,
		Handler:      hander,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info().Msg("start the api server")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("can't start the server")
	}
}
