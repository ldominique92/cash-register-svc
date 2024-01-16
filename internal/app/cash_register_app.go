package app

import "cash-register-svc/internal/domain"

type CashRegisterApp struct {
	ProductRepository domain.ProductRepository
	ShoppingCart      domain.ShoppingCart
	products          []domain.Product
}

func NewCashRegisterApp(
	productRepository domain.ProductRepository,
	applyDiscountRules map[string][]string,
) (CashRegisterApp, error) {
	validRules := make(map[domain.ProductCode][]domain.DiscountRuleName)

	for productCode, ruleNames := range applyDiscountRules {
		for _, ruleName := range ruleNames {
			//TODO: check if product IDs exist
			validRules[domain.ProductCode(productCode)] =
				append(validRules[domain.ProductCode(productCode)], domain.DiscountRuleName(ruleName))
		}
	}

	cart, err := domain.NewShoppingCart(validRules)
	if err != nil {
		return CashRegisterApp{}, err
	}

	a := CashRegisterApp{
		ProductRepository: productRepository,
		ShoppingCart:      cart,
	}

	products, err := a.ProductRepository.GetProducts()
	if err != nil {
		return CashRegisterApp{}, err
	}
	a.products = products

	return a, nil
}
