package domain

import (
	"errors"
	"fmt"
)

type DiscountRule struct {
	MinimumQuantity      int
	IsAppliedToBatches   bool
	BatchSize            int
	IsPercentageDiscount bool
	DiscountPercentage   float64
	DiscountInEuro       float64 // TODO: better to work with decimal
}

func (d DiscountRule) TotalDiscount(quantity int, price float64) (float64, error) {
	if d.IsAppliedToBatches && d.BatchSize <= 0 {
		return 0, errors.New("batch size is mandatory for discount rules applied in batches")
	}

	if d.IsPercentageDiscount && d.DiscountPercentage == 0 {
		return 0, errors.New("discount percentage is mandatory for percentage discount")
	}

	if !d.IsPercentageDiscount && d.DiscountInEuro == 0 {
		return 0, errors.New("discount in euro is mandatory for value discount")
	}

	if quantity <= d.MinimumQuantity {
		return 0, nil
	}

	numberOfItems := quantity
	if d.IsAppliedToBatches {
		numberOfItems = quantity / d.BatchSize
	}

	if d.IsPercentageDiscount {
		return price * float64(numberOfItems) * d.DiscountPercentage, nil
	}

	return price * float64(numberOfItems) * d.DiscountInEuro, nil
}

func (d DiscountRule) Description() any {
	var desc string

	if d.MinimumQuantity > 0 {
		desc = fmt.Sprintf("when buying %d or more ", d.MinimumQuantity)
	}
	if d.IsAppliedToBatches {
		desc = fmt.Sprintf(
			"%s, for every %d products apply %.2f percent discount",
			desc,
			d.BatchSize,
			d.DiscountPercentage)
	} else {
		desc = fmt.Sprintf("%s, for every %d products apply €%.2f discount", desc, d.BatchSize, d.DiscountInEuro)
	}

	return desc
}
