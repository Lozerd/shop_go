package db

import (
	"log"
	"testing"

	a "github.com/Lozerd/shop_go/internal/domain/aggregates"
	e "github.com/Lozerd/shop_go/internal/domain/entities"
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

type InitDeps struct {
	dig.In

	DB *Connection
}

func NewConnection(c *config.Configuration) *Connection {
	db, err := gorm.Open(postgres.Open(c.GetDBUrl()), &gorm.Config{})
	if err != nil {
		log.Panic("Couldn't open postgres connection")
	}

	return &Connection{db: db}
}

func Init(d InitDeps) {
	Migrate(d.DB)
}

func Migrate(c *Connection) {
	if err := c.db.AutoMigrate(a.ProductModel{}, e.Order{}); err != nil {
		log.Panic("Couldn't migrate database models")
	}
}

func SetupTestDB(t *testing.T) func(t *testing.T) {
	return func(t *testing.T) {
	}
}
