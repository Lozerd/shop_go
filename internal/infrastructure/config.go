package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	u "github.com/Lozerd/shop_go/pkg/utils"
	"github.com/joho/godotenv"
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

func (c *Configuration) GetApiPrefix() string {
	return fmt.Sprintf("/%s", c.ApiPrefix)
}

func (c *Configuration) GetApiVersion() string {
	return fmt.Sprintf("/%s", c.ApiVersion)
}

func (c *Configuration) GetApiBasePath() string {
	return fmt.Sprintf("/%s/%s", c.ApiPrefix, c.ApiVersion)
}

func LoadEnv() {
	app_env := u.StringOrDefault(os.Getenv("APP_ENV"), ".env")

	err := godotenv.Load(app_env)
	if err != nil {
		log.Panic(fmt.Sprintf("Couldn't load %s file", app_env))
	}
}

func LoadConfig() {
	LoadEnv()

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
