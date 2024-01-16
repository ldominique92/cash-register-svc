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
	applyDiscountRules map[string][]string,
) (CashRegisterApp, error) {
	validRules := make(map[domain.ProductCode]domain.DiscountRuleName)

	for productCode, ruleNames := range applyDiscountRules {
		for _, ruleName := range ruleNames {
			//TODO: check if product IDs exist
			validRules[domain.ProductCode(productCode)] = domain.DiscountRuleName(ruleName)
		}
	}

	cart, err := domain.NewShoppingCart(validRules)
	if err != nil {
		return CashRegisterApp{}, err
	}

	a := CashRegisterApp{
		ProductRepository: productRepository,
		ProductCache:      productCache,
		ShoppingCart:      cart,
	}

	products, err := a.ProductRepository.GetProducts()
	if err != nil {
		return CashRegisterApp{}, err
	}

	err = a.ProductCache.Load(products)
	if err != nil {
		return CashRegisterApp{}, err
	}

	return a, nil
}

func (a CashRegisterApp) AddProductToCart(productCode string, quantity int64) error {
	product, err := a.ProductCache.GetProduct(domain.ProductCode(productCode))
	if err != nil {
		return err
	}

	if len(product.Code) == 0 {
		return errors.New(fmt.Sprintf("product with code %s not found", productCode))
	}

	a.ShoppingCart.AddProduct(product, quantity)

	return nil
}

func (a CashRegisterApp) GetTotal() float64 {
	return a.ShoppingCart.GetTotal()
}
