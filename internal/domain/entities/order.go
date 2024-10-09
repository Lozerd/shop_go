package entities

import "github.com/jinzhu/gorm"

type Order struct {
    gorm.Model
    ID int
}
