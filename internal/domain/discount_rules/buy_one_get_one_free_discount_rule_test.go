package discount_rules_test

import (
	"cash-register-svc/internal/domain/discount_rules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyOneGetOneFreeDiscountRule_TotalDiscount(t *testing.T) {
	newRule := discount_rules.BuyOneGetOneFreeDiscountRule{}
	discount := newRule.TotalDiscount(10, 25.5)
	assert.Equal(t, discount, 127.5)

	discount = newRule.TotalDiscount(5, 8.30)
	assert.Equal(t, discount, 16.6)
}
