package cfg

import (
	"gitlab.cloud.n-ask.com/n-ask/env"
	"log"
)

var Config struct {
	Verbose       bool   `env:"VERBOSE"`
	ListenAddress string `env:"LISTEN_ADDRESS"`
	Port          int    `env:"PORT"`

	DatabaseURL string `env:"DATABASE_URL"`
}

func init() {
	if err := env.Load(&Config); err != nil {
		log.Fatal(err)
	}
	if Config.Port <= 0 {
		Config.Port = 5454
	}

	if len(Config.DatabaseURL) == 0 {
		log.Fatal("missing DATABASE_URL...")
	}
}
