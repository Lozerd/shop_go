package config

import (
	"errors"
	"fmt"

	// "log"
	"os"
	"strings"

	u "github.com/Lozerd/shop_go/pkg/utils"
	"github.com/joho/godotenv"
	// "github.com/joho/godotenv"
)

type Server struct {
	Port string
	Host string
}

type Database struct {
	Name string
	User string
	Host string
	Pass string
	Port string
}

type Configuration struct {
	Server     Server
	Database   Database
	ApiVersion string
	ApiPrefix  string
}

var Config *Configuration

func (c *Configuration) GetAddr() string {
	return strings.Join([]string{c.Server.Host, c.Server.Port}, ":")
}

func (c *Configuration) GetDBUrl() string {
	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s", // sslmode=disable TimeZone=Asia/Shanghai",
		c.Database.Host, c.Database.User, c.Database.Pass,
		c.Database.Name, c.Database.Port,
	)
}

func GetApiPrefix() string {
	return fmt.Sprintf("/%s", Config.ApiPrefix)
}

func GetApiVersion() string {
	return fmt.Sprintf("/%s", Config.ApiVersion)
}

func GetApiBasePath() string {
	return fmt.Sprintf("/%s/%s", Config.ApiPrefix, Config.ApiVersion)
}

func GetConfig() *Configuration {
	return Config
}

func LoadEnv() (err error) {
	app_env := u.StringOrDefault(os.Getenv("APP_ENV"), ".env")

	if app_env == "" {
		return errors.New("filepath argument must not be empty")
	}

    if _, err := os.Stat(app_env); errors.Is(err, os.ErrNotExist) {
        return fmt.Errorf("Filepath do not exist\nOriginal: %s", err)
	}

	err = godotenv.Load(app_env)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't load %s file", app_env))
	}

	return nil
}

func LoadConfig() {
	// LoadEnv()

	HOST := u.StringOrDefault(os.Getenv("HOST"), "0.0.0.0")
	PORT := u.StringOrDefault(os.Getenv("PORT"), "8080")

	dbName := u.StringOrDefault(os.Getenv("POSTGRES_NAME"), "dev_shop")
	dbHost := u.StringOrDefault(os.Getenv("POSTGRES_HOST"), "localhost")
	dbPort := u.StringOrDefault(os.Getenv("POSTGRES_PORT"), "5432")
	dbUser := u.StringOrDefault(os.Getenv("POSTGRES_USER"), "dev_shop")
	dbPass := u.StringOrDefault(os.Getenv("POSTGRES_PASSWORD"), "password")

	apiVersion := u.StringOrDefault(os.Getenv("API_VERSION"), "v1")
	apiPrefix := u.StringOrDefault(os.Getenv("API_PREFIX"), "api")

	Config = &Configuration{
		Server: Server{
			Host: HOST,
			Port: PORT,
		},
		Database: Database{
			Name: dbName,
			User: dbUser,
			Host: dbHost,
			Pass: dbPass,
			Port: dbPort,
		},
		ApiVersion: apiVersion,
		ApiPrefix:  apiPrefix,
	}
}
