package domain

const (
	BulkDiscountRuleName = "BULK"
	MinimumBulkSize      = 3
	DiscountPerItem      = 0.50
)

type BulkDiscountRule struct {
}

func (t BulkDiscountRule) Name() string {
	return BulkDiscountRuleName
}

func (t BulkDiscountRule) TotalDiscount(quantity int64, price float64) float64 {
	if quantity >= MinimumBulkSize {
		return DiscountPerItem * float64(quantity)
	}

	return 0
}
