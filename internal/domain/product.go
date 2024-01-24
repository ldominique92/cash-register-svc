package domain

type ProductCode string

type Product struct {
	Code         ProductCode  `json:"code"`
	Name         string       `json:"name"`
	Price        float64      `json:"price"`
	DiscountRule DiscountRule `json:"discount_rule"`
}

type ProductRepository interface {
	GetProducts() ([]Product, error)
}

type ProductCache interface {
	Load([]Product) error
	GetProduct(code ProductCode) (Product, error)
	ListProducts() []Product
}
