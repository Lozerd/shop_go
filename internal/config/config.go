package config

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/lozerd/shop_go/infrastructure/logging"
)

var logger = logging.NewLogger("config", false)

type DB struct {
	Name     string `env:"DB_NAME,required=true"`
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=5432"`
	User     string `env:"DB_USER,required=true"`
	Password string `env:"DB_PASSWORD,required=true"`
}

type Config struct {
	Debug          bool     `env:"DEBUG,default=false"`
	GinMode        string   `env:"GIN_MODE,default=release"`
	ApiPrefix      string   `env:"API_PREFIX,default=api"`
	Port           int      `env:"PORT,default=8000"`
	TrustedProxies []string `env:"TRUSTED_PROXIES"`
	DB
}

func (c *Config) GetApiPrefix() string {
	return fmt.Sprintf("/%s", c.ApiPrefix)
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf(":%d", c.Port)
}

func validateGinMode(mode string) bool {
	return mode == "debug" || mode == "release"
}

var config *Config

func Load() (*Config, error) {
	var c Config
	if _, err := env.UnmarshalFromEnviron(&c); err != nil {
		logger.Error().Err(err)
		return nil, err
	}
	if !validateGinMode(c.GinMode) {
		err := ErrInvalidGinMode{value: c.GinMode}
		logger.Error().Err(err).Msgf(err.Error(), c.GinMode)
		return nil, err
	}
	return &c, nil
}

func GetConfig() *Config {
	if config == nil {
		c, err := Load()
		if err != nil {
			logger.Fatal().Err(err).Msg("Could not load config")
		}
		config = c
	}
	return config
}
