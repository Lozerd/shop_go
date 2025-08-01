package product

import "github.com/lozerd/shop_go/domain/product"

type ProductInDTO struct {
	Name string `json:"name" form:"name" binding:"required,min=10"`
}

type ProductOutDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (dto *ProductOutDTO) From(product *product.Product) *ProductOutDTO {
	dto.ID = product.ID
	dto.Name = product.Name
	return dto
}
