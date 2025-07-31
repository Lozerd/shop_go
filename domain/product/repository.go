package product

type IProductRepository interface {
	GetAll() ([]Product, error)
	Save(*Product) (*Product, error)
	Delete(id int) error
}
