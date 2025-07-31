package product

import (
	"github.com/lozerd/shop_go/domain/product"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) Save(product *product.Product) (*product.Product, error) {
	res := r.db.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (r *ProductRepository) Delete(id int) error {
	res := r.db.Delete(&product.Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
