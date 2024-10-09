package entities

type Product struct {
	ID uint `gorm:"primaryKey"`
    Title string 
}
