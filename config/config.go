package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = ""

/*
NOTE TO SELF - Public properties must start with CAPITAL
*/

type Configuration struct {
	TwitchConfiguration
	HTTPServer
}

type TwitchConfiguration struct {
	Url  string `envconfig:"TWITCH_WS_URL" default:"irc-ws.chat.twitch.tv:443"`
	Nick string `envconfig:"TWITCH_NICK" default:"justinfan821"`
	Pass string `envconfig:"TWITCH_PASS" default:"ANONAUTH"`
}

type HTTPServer struct {
	IdleTimeout  time.Duration `envconfig:"HTTP_SERVER_IDLE_TIMEOUT" default:"60s"`
	ReadTimeout  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"1s"`
	Port         int           `envconfig:"PORT" default:"5000"`
	WriteTimeout time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"2s"`
}

func Load() (Configuration, error) {
	var cfg Configuration
	err := envconfig.Process(envPrefix, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
