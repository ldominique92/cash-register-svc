package discount_rules

const (
	CashBulkDiscountRuleName = "CASH_BULK_DISCOUNT_RULE"
	MinimumCashBulkSize      = 3
	DiscountPerItem          = 0.50
)

type CashBulkDiscountRule struct {
}

func (t CashBulkDiscountRule) Name() string {
	return CashBulkDiscountRuleName
}

func (t CashBulkDiscountRule) TotalDiscount(quantity int, _ float64) float64 {
	if quantity >= MinimumCashBulkSize {
		return DiscountPerItem * float64(quantity)
	}

	return 0
}
