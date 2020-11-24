package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Server ...
	Server struct {
		Port string `envconfig:"GIN_SERVER_PORT" default:"8088"`
	}

	// Database ...
	Database struct {
		Port string `envconfig:"GIN_DATABASE_PORT" default:"5432"`
	}

	setting struct {
		Server   Server
		Database Database
	}
)

// Setting config
var Setting = &setting{}

// Load load config
func Load() {
	envconfig.MustProcess("", Setting)
}
