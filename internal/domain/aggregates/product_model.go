package aggregates

import (
	. "github.com/Lozerd/shop_go/internal/domain/entities"
	"github.com/Lozerd/shop_go/internal/infrastructure/db/base"
)

type ProductModel struct {
	base.Base
	Title    string   `gorm:"size:255"`

    CategoryID uint
	Category Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    ProductID uint
	Product  Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
}

// Table "aggregate_product_model" as ProductModel {
//   "id" int [not null, pk, increment]
//   "category_id" int [not null, ref: > Category.id]
//   "product_id" int [not null, ref: > Product.id]
//
//   Indexes {
//     id [unique]
//   }
// }
