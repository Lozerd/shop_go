package db

import (
	"log"

	i "github.com/Lozerd/shop_go/internal/infrastructure"
    e "github.com/Lozerd/shop_go/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    DB = NewConnection()
    DB.AutoMigrate(e.ProductModel{}, e.Order{})
}

func NewConnection() *gorm.DB {
    db, err := gorm.Open(postgres.Open(i.Config.GetDBUrl()), &gorm.Config{})
    if err != nil {
        log.Panic("Couldn't open postgres connection")
    }

    return db
}
