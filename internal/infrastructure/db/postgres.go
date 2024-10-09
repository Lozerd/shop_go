package db

import (
	"log"

	c "github.com/Lozerd/shop_go/internal/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    DB = NewConnection()
}

func NewConnection() *gorm.DB {
    db, err := gorm.Open(postgres.Open(c.Config.GetDBUrl()), &gorm.Config{})
    if err != nil {
        log.Panic("Couldn't open postgres connection")
    }

    return db
}
