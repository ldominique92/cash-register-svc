package domain

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type DiscountRule struct {
	MinimumQuantity      int64           `json:"minimum_quantity"`
	IsAppliedToBatches   bool            `json:"is_applied_to_batches"`
	BatchSize            int64           `json:"batch_size"`
	IsPercentageDiscount bool            `json:"is_percentage_discount"`
	DiscountPercentage   decimal.Decimal `json:"discount_percentage"`
	DiscountInEuro       decimal.Decimal `json:"discount_in_euro"`
}

func (d DiscountRule) TotalDiscount(quantity int64, price decimal.Decimal) (decimal.Decimal, error) {
	if d.IsAppliedToBatches && d.BatchSize <= 0 {
		return decimal.Zero, errors.New("batch size is mandatory for discount rules applied in batches")
	}

	if d.IsPercentageDiscount && d.DiscountPercentage == decimal.Zero {
		return decimal.Zero, errors.New("discount percentage is mandatory for percentage discount")
	}

	if !d.IsPercentageDiscount && d.DiscountInEuro == decimal.Zero {
		return decimal.Zero, errors.New("discount in euro is mandatory for value discount")
	}

	if quantity < d.MinimumQuantity {
		return decimal.Zero, nil
	}

	numberOfItems := quantity
	if d.IsAppliedToBatches {
		numberOfItems = quantity / d.BatchSize
	}

	if d.IsPercentageDiscount {
		return price.Mul(decimal.NewFromInt(numberOfItems).Mul(d.DiscountPercentage)), nil
	}

	return decimal.NewFromInt(numberOfItems).Mul(d.DiscountInEuro), nil
}

func (d DiscountRule) Description() any {
	var desc string

	if d.MinimumQuantity > 0 {
		desc = fmt.Sprintf("when buying %d or more ", d.MinimumQuantity)
	}
	if d.IsAppliedToBatches {
		desc = fmt.Sprintf(
			"%s, for every %d products apply %v percent discount",
			desc,
			d.BatchSize,
			d.DiscountPercentage.StringFixed(2))
	} else {
		desc = fmt.Sprintf(
			"%s, for every %d products apply €%v discount",
			desc,
			d.BatchSize,
			d.DiscountInEuro.StringFixed(2))
	}

	return desc
}
