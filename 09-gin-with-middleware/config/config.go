package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Server provides the system configuration.
	Server struct {
		Port string `envconfig:"GIN_SERVER_PORT" default:"8088"`
	}

	// Logs provides the system configuration.
	Logs struct {
		Pretty bool `envconfig:"LOGS_PRETTY" default:"true"`
		Color  bool `envconfig:"LOGS_COLOR" default:"true"`
	}

	setting struct {
		Server
		Logs
		Debug bool `envconfig:"GIN_SERVER_DEBUG" default:"true"`
	}
)

var (
	// Setting config
	Setting = &setting{}
)

func init() {
	envconfig.MustProcess("", Setting)
}
