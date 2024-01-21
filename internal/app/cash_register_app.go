package app

import (
	"cash-register-svc/internal/domain"
	"errors"
	"fmt"
	"strings"
)

type CashRegisterApp struct {
	ProductRepository domain.ProductRepository
	ShoppingCart      domain.ShoppingCart
	ProductCache      domain.ProductCache
}

func NewCashRegisterApp(
	productRepository domain.ProductRepository,
	productCache domain.ProductCache,
	applyDiscountRules map[string]string,
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

	validRules := make(map[domain.ProductCode]domain.DiscountRuleName)
	for productCode, ruleName := range applyDiscountRules {
		_, err := a.getProductFromCache(strings.ToUpper(productCode))
		if err != nil {
			return CashRegisterApp{}, err
		}
		validRules[domain.ProductCode(productCode)] = domain.DiscountRuleName(ruleName)
	}

	cart, err := domain.NewShoppingCart(validRules)
	if err != nil {
		return CashRegisterApp{}, err
	}
	a.ShoppingCart = cart

	return a, nil
}

func (a CashRegisterApp) AddProductToCart(productCode string, quantity int) error {
	product, err := a.getProductFromCache(productCode)
	if err != nil {
		return err
	}

	a.ShoppingCart.AddProduct(product, quantity)
	return nil
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

func (a CashRegisterApp) GetTotal() float64 {
	return a.ShoppingCart.GetTotal()
}

func (a CashRegisterApp) GetProducts() []domain.Product {
	return a.ProductCache.ListProducts()
}
