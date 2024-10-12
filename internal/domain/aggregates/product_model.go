package aggregates

type ProductModel struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
    Title string  `gorm:"size:255"`
}
