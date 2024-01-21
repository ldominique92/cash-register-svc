package domain

import "cash-register-svc/internal/domain/discount_rules"

type DiscountRuleName string

type DiscountRule interface {
	TotalDiscount(quantity int, price float64) float64
	Name() string
}

var DiscountRules = map[DiscountRuleName]DiscountRule{
	discount_rules.BuyOneGetOneFreeDiscountRuleName: discount_rules.BuyOneGetOneFreeDiscountRule{},
	discount_rules.CashBulkDiscountRuleName:         discount_rules.CashBulkDiscountRule{},
	discount_rules.PercentageBulkDiscountRuleName:   discount_rules.PercentageBulkDiscountRule{},
}
