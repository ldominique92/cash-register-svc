package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyOneGetOneFreeDiscountRule_TotalDiscount(t *testing.T) {
	newRule := domain.BuyOneGetOneFreeDiscountRule{}
	discount := newRule.TotalDiscount(10, 25.5)
	assert.Equal(t, discount, 127.5)

	discount = newRule.TotalDiscount(5, 8.30)
	assert.Equal(t, discount, 16.6)
}
