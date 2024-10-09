package infrastructure

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Name string `yaml:"name"`
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Pass string `yaml:"pass"`
		Port string `yaml:"port"`
	} `yaml:"database"`
}

var Config Configuration

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

func LoadConfig() {
	buf, err := os.ReadFile("configs/development.yaml")
	if err != nil {
		log.Panic("Couldn't find development config in configs/development.yaml")
	}

	err = yaml.UnmarshalStrict(buf, &Config)
	if err != nil {
		log.Panic("Couldn't parse config")
	}
}
