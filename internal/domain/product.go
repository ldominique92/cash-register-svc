package domain

type ProductCode string

type Product struct {
	Code  ProductCode
	Name  string
	Price float64
}

type ProductRepository interface {
	GetProducts() ([]Product, error)
}
