package entities

type Product struct {
	ID    uint `gorm:"primaryKey;autoIncrement;"`
    Title string `gorm:"size:255";index:"idx_name,type:btree"`
}

// Table "entity_product" as Product {
//   "id" int [not null, pk, increment]
//   "title" varchar(255) [not null]
//   "price_id" int [null, ref: > Price.id]
// }
