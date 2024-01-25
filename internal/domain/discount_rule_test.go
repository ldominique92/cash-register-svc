package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDiscountRule_TotalDiscount(t *testing.T) {
	buyOneGetOneFreeDiscountRule := domain.DiscountRule{
		MinimumQuantity:      2,
		IsAppliedToBatches:   true,
		BatchSize:            2,
		IsPercentageDiscount: true,
		DiscountPercentage:   decimal.NewFromInt(1),
		DiscountInEuro:       decimal.Zero,
	}

	discount, err := buyOneGetOneFreeDiscountRule.TotalDiscount(10, decimal.NewFromFloat(25.5))
	assert.Nil(t, err)
	assert.Equal(t, discount, decimal.NewFromFloat(127.5))

	discount, err = buyOneGetOneFreeDiscountRule.TotalDiscount(5, decimal.NewFromFloat(8.30))
	assert.Nil(t, err)
	assert.Equal(t, discount, decimal.NewFromFloat(16.6))

	buyTreeStrawberriesOrMoreAndGet50CentsDiscount := domain.DiscountRule{
		MinimumQuantity:      3,
		IsAppliedToBatches:   false,
		BatchSize:            0,
		IsPercentageDiscount: false,
		DiscountPercentage:   decimal.Zero,
		DiscountInEuro:       decimal.NewFromFloat(0.50),
	}
	discount, err = buyTreeStrawberriesOrMoreAndGet50CentsDiscount.TotalDiscount(2, decimal.NewFromInt(6.00))
	assert.Nil(t, err)
	assert.Equal(t, discount, decimal.Zero)

	discount, err = buyTreeStrawberriesOrMoreAndGet50CentsDiscount.TotalDiscount(5, decimal.NewFromInt(6.00))
	assert.Nil(t, err)
	assert.Equal(t, discount, decimal.NewFromFloat(2.50))

	buyTreeCoffeesOrMoreAndPayTwoThirds := domain.DiscountRule{
		MinimumQuantity:      3,
		IsAppliedToBatches:   false,
		BatchSize:            0,
		IsPercentageDiscount: true,
		DiscountPercentage:   decimal.NewFromInt(1).Div(decimal.NewFromInt(3)),
		DiscountInEuro:       decimal.Zero,
	}

	discount, err = buyTreeCoffeesOrMoreAndPayTwoThirds.TotalDiscount(2, decimal.NewFromInt(6))
	assert.Nil(t, err)
	assert.Equal(t, discount, decimal.Zero)

	discount, err = buyTreeCoffeesOrMoreAndPayTwoThirds.TotalDiscount(5, decimal.NewFromInt(6))
	assert.Nil(t, err)
	assert.Equal(t, discount.StringFixed(2), "10.00")
}
