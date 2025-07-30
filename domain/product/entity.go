package product

import (
	"fmt"

	"gorm.io/gorm"
)

type GenderEnum string

const (
	MALE   GenderEnum = "MALE"
	FEMALE GenderEnum = "FEMALE"
)

func (e *GenderEnum) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid str")
	}
	*e = GenderEnum(str)

	return nil
}

func (e GenderEnum) Value() (interface{}, error) {
	return string(e), nil
}

type User struct {
	gorm.Model
	Name   string
	Email  string
	Gender *GenderEnum `gorm:"type:gender"`
}
