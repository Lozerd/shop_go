package entities

import "github.com/Lozerd/shop_go/internal/infrastructure/db/base"

type Category struct {
    base.Base
    Title    string `gorm:"size:255,index:idx_name,type:btree"`
    ParentID *uint
	Parent   *Category
}

// Table "entity_category" as Category {
//   "id" int [not null, pk, increment]
//   "title" varchar(255) [not null]
//   "parent_id" int [null, ref: > Category.id]
//
//   Indexes {
//     title [name:"idx_category_title", type: btree]
//     id [unique]
//   }
// }
