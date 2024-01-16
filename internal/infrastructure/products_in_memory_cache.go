package infrastructure

import "cash-register-svc/internal/domain"

type ProductsInMemoryCache struct {
	products map[domain.ProductCode]domain.Product
}

func NewProductsInMemoryCache() *ProductsInMemoryCache {
	return &ProductsInMemoryCache{}
}

func (c *ProductsInMemoryCache) Load(products []domain.Product) error {
	c.products = make(map[domain.ProductCode]domain.Product)

	for _, p := range products {
		c.products[p.Code] = p
	}

	return nil
}

func (c *ProductsInMemoryCache) GetProduct(code domain.ProductCode) (domain.Product, error) {
	if p, ok := c.products[code]; ok {
		return p, nil
	}

	return domain.Product{}, nil
}
