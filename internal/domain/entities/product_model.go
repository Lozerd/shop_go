package entities

type ProductModel struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
    Title string 
}
