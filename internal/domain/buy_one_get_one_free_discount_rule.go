package domain

const BuyOneGetOneFreeDiscountRuleName = "BUY_ONE_GET_ONE_FREE"

type BuyOneGetOneFreeDiscountRule struct{}

func (t BuyOneGetOneFreeDiscountRule) Name() string {
	return BuyOneGetOneFreeDiscountRuleName
}

func (t BuyOneGetOneFreeDiscountRule) TotalDiscount(quantity int64, price float64) float64 {
	pairs := quantity / 2
	return float64(pairs) * price
}
