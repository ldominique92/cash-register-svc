package discount_rules

const BuyOneGetOneFreeDiscountRuleName = "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE"

type BuyOneGetOneFreeDiscountRule struct{}

func (t BuyOneGetOneFreeDiscountRule) Name() string {
	return BuyOneGetOneFreeDiscountRuleName
}

func (t BuyOneGetOneFreeDiscountRule) TotalDiscount(quantity int, price float64) float64 {
	pairs := quantity / 2
	return float64(pairs) * price
}
