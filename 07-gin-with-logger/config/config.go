package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Server ...
type Server struct {
	Debug bool   `envconfig:"GIN_DEBUG"`
	Port  string `envconfig:"GIN_SERVER_PORT"`
}

// Logs ...
type Logs struct {
	Pretty bool `envconfig:"GIN_LOGS_PRETTY"`
}

type setting struct {
	Server Server
	Debug  bool `envconfig:"GIN_DEBUG"`
}

// Setting config
var Setting = &setting{}

// Load load config
func Load() {
	envconfig.MustProcess("", Setting)
}
