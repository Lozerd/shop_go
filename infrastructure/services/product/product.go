package product

import (
	"github.com/lozerd/shop_go/domain/product"
	dp "github.com/lozerd/shop_go/interfaces/dto/product"
)

type ProductService struct {
	repo product.IProductRepository
}

func NewProductService(r product.IProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) List() ([]product.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(dto dp.ProductInDTO) (*product.Product, error) {
	return s.repo.Save(&product.Product{Name: dto.Name})
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
