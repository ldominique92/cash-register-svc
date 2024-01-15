package domain

type DiscountRuleName string

type DiscountRule interface {
	TotalDiscount(quantity int64, price float64) float64
	Name() string
}

var DiscountRules = map[DiscountRuleName]DiscountRule{
	BuyOneGetOneFreeDiscountRuleName: BuyOneGetOneFreeDiscountRule{},
}
