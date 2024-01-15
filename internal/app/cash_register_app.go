package app

import "cash-register-svc/internal/domain"

type CashRegisterApp struct {
	ShoppingCart domain.ShoppingCart
}

func NewCashRegisterApp(applyDiscountRules map[string][]string) (CashRegisterApp, error) {
	validRules := make(map[domain.ProductCode][]domain.DiscountRuleName)

	for productCode, ruleNames := range applyDiscountRules {
		for _, ruleName := range ruleNames {
			// TODO: check if product IDs exist
			validRules[domain.ProductCode(productCode)] =
				append(validRules[domain.ProductCode(productCode)], domain.DiscountRuleName(ruleName))
		}
	}

	cart, err := domain.NewShoppingCart(validRules)
	if err != nil {
		return CashRegisterApp{}, err
	}

	return CashRegisterApp{
		ShoppingCart: cart,
	}, nil
}
