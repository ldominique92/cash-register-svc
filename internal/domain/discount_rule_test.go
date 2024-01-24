package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscountRule_TotalDiscount(t *testing.T) {
	buyOneGetOneFreeDiscountRule := domain.DiscountRule{
		MinimumQuantity:      2,
		IsAppliedToBatches:   true,
		BatchSize:            2,
		IsPercentageDiscount: true,
		DiscountPercentage:   1,
		DiscountInEuro:       0,
	}

	discount, err := buyOneGetOneFreeDiscountRule.TotalDiscount(10, 25.5)
	assert.Nil(t, err)
	assert.Equal(t, discount, 127.5)

	discount, err = buyOneGetOneFreeDiscountRule.TotalDiscount(5, 8.30)
	assert.Nil(t, err)
	assert.Equal(t, discount, 16.6)

	buyTreeStrawberriesOrMoreAndGet50CentsDiscount := domain.DiscountRule{
		MinimumQuantity:      3,
		IsAppliedToBatches:   false,
		BatchSize:            0,
		IsPercentageDiscount: false,
		DiscountPercentage:   0,
		DiscountInEuro:       0.50,
	}
	discount, err = buyTreeStrawberriesOrMoreAndGet50CentsDiscount.TotalDiscount(2, 6.00)
	assert.Nil(t, err)
	assert.Equal(t, discount, float64(0))

	discount, err = buyTreeStrawberriesOrMoreAndGet50CentsDiscount.TotalDiscount(5, 6.00)
	assert.Nil(t, err)
	assert.Equal(t, discount, 2.50)

	buyTreeCoffeesOrMoreAndPayTwoThirds := domain.DiscountRule{
		MinimumQuantity:      3,
		IsAppliedToBatches:   false,
		BatchSize:            0,
		IsPercentageDiscount: true,
		DiscountPercentage:   float64(1) / float64(3),
		DiscountInEuro:       0,
	}

	discount, err = buyTreeCoffeesOrMoreAndPayTwoThirds.TotalDiscount(2, 6.00)
	assert.Nil(t, err)
	assert.Equal(t, discount, float64(0))

	discount, err = buyTreeCoffeesOrMoreAndPayTwoThirds.TotalDiscount(5, 6.00)
	assert.Nil(t, err)
	assert.Equal(t, discount, float64(10))
}
