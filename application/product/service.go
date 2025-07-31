package product

import (
	"github.com/lozerd/shop_go/domain/product"
	dp "github.com/lozerd/shop_go/interfaces/dto/product"
)

type IProductService interface {
	List() ([]product.Product, error)
	Create(dto dp.ProductDTO) (*product.Product, error)
	Delete(id int) error
}
