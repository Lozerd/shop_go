package config

import (
	"errors"
	"fmt"

	"os"
	"strings"

	. "github.com/Lozerd/shop_go/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

type Server struct {
	Port string
	Host string
}

type Client struct {
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
	Server      Server
	Client      Client
	CorsConfig  cors.Config
	Database    Database
	ApiVersion  string
	ApiPrefix   string
	initialized bool
}

var config *Configuration

func (c *Configuration) GetServerAddr() string {
	return strings.Join([]string{c.Server.Host, c.Server.Port}, ":")
}

func (c *Configuration) GetClientAddr() string {
	return "localhost:8000"
}

func (c *Configuration) GetDBUrl() string {
	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s", // sslmode=disable TimeZone=Asia/Shanghai",
		c.Database.Host, c.Database.User, c.Database.Pass,
		c.Database.Name, c.Database.Port,
	)
}

func (c *Configuration) GetApiPrefix() string {
	return fmt.Sprintf("/%s", c.ApiPrefix)
}

func (c *Configuration) GetApiVersion() string {
	return fmt.Sprintf("/%s", c.ApiVersion)
}

func (c *Configuration) GetApiBasePath() string {
	return fmt.Sprintf("/%s/%s", c.ApiPrefix, c.ApiVersion)
}

func GetConfig() *Configuration {
	if config == nil || !config.initialized {
		LoadConfig()
	}
	return config
}

func LoadEnv() (err error) {
	app_env := StringOrDefault(os.Getenv("APP_ENV"), ".env")

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
	LoadEnv()

	SERVER_HOST := StringOrDefault(os.Getenv("SERVER_HOST"), "0.0.0.0")
	SERVER_PORT := StringOrDefault(os.Getenv("SERVER_PORT"), "8080")

	CLIENT_HOST := StringOrDefault(os.Getenv("CLIENT_HOST"), "0.0.0.0")
	CLIENT_PORT := StringOrDefault(os.Getenv("CLIENT_PORT"), "8000")

    var CORS_ORIGINS []string
	if origins := os.Getenv("CORS_ORIGINS"); origins != "" {
		CORS_ORIGINS = strings.Split(origins, ",")
	}

	dbName := StringOrDefault(os.Getenv("POSTGRES_NAME"), "dev_shop")
	dbHost := StringOrDefault(os.Getenv("POSTGRES_HOST"), "localhost")
	dbPort := StringOrDefault(os.Getenv("POSTGRES_PORT"), "5432")
	dbUser := StringOrDefault(os.Getenv("POSTGRES_USER"), "dev_shop")
	dbPass := StringOrDefault(os.Getenv("POSTGRES_PASSWORD"), "password")

	apiVersion := StringOrDefault(os.Getenv("API_VERSION"), "v1")
	apiPrefix := StringOrDefault(os.Getenv("API_PREFIX"), "api")

	config = &Configuration{
		Server: Server{
			Host: SERVER_HOST,
			Port: SERVER_PORT,
		},
		Client: Client{
			Host: CLIENT_HOST,
			Port: CLIENT_PORT,
		},
		CorsConfig: cors.Config{
			AllowOrigins: CORS_ORIGINS,
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
		// initialized: true,
	}
}
