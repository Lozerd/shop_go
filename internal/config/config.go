package config

import (
	"github.com/Netflix/go-env"
	"github.com/rs/zerolog/log"
)

type DB struct {
	Name     string `env:"DB_NAME,required=true"`
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=5432"`
	User     string `env:"DB_USER,required=true"`
	Password string `env:"DB_PASSWORD,required=true"`
}

type Config struct {
	GinMode string `env:"GIN_MODE,default=release"`
	Port    int    `env:"PORT,default=8000"`
	TrustedProxies []string `env:"TRUSTED_PROXIES"`
	DB
}

func validateGinMode(mode string) bool {
	return mode == "debug" || mode == "release"
}

func Load() (*Config, error) {
	var config Config
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		msg := "Could not parse environment variables: %s"
		log.Error().Err(err).Msg(msg)
		return nil, err
	}
	if !validateGinMode(config.GinMode) {
		msg := "Invalid gin mode: %s, should be one of [debug, release]"
		err := ErrInvalidGinMode{value: config.GinMode}
		log.Error().Err(err).Msgf(msg, config.GinMode)
		return nil, err
	}
	return &config, nil
}
