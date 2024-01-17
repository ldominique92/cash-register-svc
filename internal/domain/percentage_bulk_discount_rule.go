package domain

const (
	PercentageDiscountRuleName = "BULK"
	MinimumPercentageBulkSize  = 3
	DiscountPercentage         = float64(1) / float64(3)
)

type PercentageBulkDiscountRule struct {
}

func (t PercentageBulkDiscountRule) Name() string {
	return PercentageDiscountRuleName
}

func (t PercentageBulkDiscountRule) TotalDiscount(quantity int64, price float64) float64 {
	if quantity >= MinimumPercentageBulkSize {
		return DiscountPercentage * price
	}

	return 0
}
