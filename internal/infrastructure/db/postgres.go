package db

import (
	"log"

	i "github.com/Lozerd/shop_go/internal/infrastructure"
    e "github.com/Lozerd/shop_go/internal/domain/entities"
    a "github.com/Lozerd/shop_go/internal/domain/aggregates"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    DB = NewConnection()
    DB.AutoMigrate(a.ProductModel{}, e.Order{})
}

func NewConnection() *gorm.DB {
    db, err := gorm.Open(postgres.Open(i.Config.GetDBUrl()), &gorm.Config{})
    if err != nil {
        log.Panic("Couldn't open postgres connection")
    }

    return db
}
