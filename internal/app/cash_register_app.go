package app

import (
	"cash-register-svc/internal/domain"
	"errors"
	"fmt"
)

type CashRegisterApp struct {
	ProductRepository domain.ProductRepository
	ShoppingCart      domain.ShoppingCart
	ProductCache      domain.ProductCache
}

func NewCashRegisterApp(
	productRepository domain.ProductRepository,
	productCache domain.ProductCache,
) (CashRegisterApp, error) {
	a := CashRegisterApp{
		ProductRepository: productRepository,
		ProductCache:      productCache,
	}

	products, err := a.ProductRepository.GetProducts()
	if err != nil {
		return CashRegisterApp{}, err
	}

	err = a.ProductCache.Load(products)
	if err != nil {
		return CashRegisterApp{}, err
	}

	a.ShoppingCart = domain.NewShoppingCart()

	return a, nil
}

func (a CashRegisterApp) AddProductToCart(productCode string, quantity int) error {
	product, err := a.getProductFromCache(productCode)
	if err != nil {
		return err
	}

	return a.ShoppingCart.AddProduct(product, quantity)
}

func (a CashRegisterApp) getProductFromCache(productCode string) (domain.Product, error) {
	product, err := a.ProductCache.GetProduct(domain.ProductCode(productCode))
	if err != nil {
		return domain.Product{}, err
	}

	if len(product.Code) == 0 {
		return domain.Product{}, errors.New(fmt.Sprintf("product with code %s not found", productCode))
	}
	return product, nil
}

func (a CashRegisterApp) GetTotal() (float64, error) {
	return a.ShoppingCart.Total()
}

func (a CashRegisterApp) GetProducts() []domain.Product {
	return a.ProductCache.ListProducts()
}
