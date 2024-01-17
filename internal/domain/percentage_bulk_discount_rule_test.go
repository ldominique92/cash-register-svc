package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCashBulkDiscountRule_TotalDiscount(t *testing.T) {
	newRule := domain.CashBulkDiscountRule{}
	discount := newRule.TotalDiscount(2, 6.00)
	assert.Equal(t, discount, float64(0))

	discount = newRule.TotalDiscount(5, 6.00)
	assert.Equal(t, discount, 2.50)
}
