package db

import (
	"log"
	"testing"

	a "github.com/Lozerd/shop_go/internal/domain/aggregates"
	e "github.com/Lozerd/shop_go/internal/domain/entities"
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = NewConnection()
	DB.AutoMigrate(a.ProductModel{}, e.Order{})
}

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Config.GetDBUrl()), &gorm.Config{})
	if err != nil {
		log.Panic("Couldn't open postgres connection")
	}

	return db
}

func SetupTestDB(t *testing.T) func(t *testing.T) {
	return func(t *testing.T) {
	}
}
