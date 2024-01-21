package discount_rules

const (
	PercentageBulkDiscountRuleName = "PERCENTAGE_BULK_DISCOUNT_RULE"
	MinimumPercentageBulkSize      = 3
	DiscountPercentage             = float64(1) / float64(3)
)

type PercentageBulkDiscountRule struct {
}

func (t PercentageBulkDiscountRule) Name() string {
	return PercentageBulkDiscountRuleName
}

func (t PercentageBulkDiscountRule) TotalDiscount(quantity int64, price float64) float64 {
	if quantity >= MinimumPercentageBulkSize {
		a := DiscountPercentage * price * float64(quantity)
		return a
	}

	return 0
}
