package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCashRegisterApp(t *testing.T) {
	discountRules := map[string][]string{
		"PRD1": {"BUY_ONE_GET_ONE_FREE"},
		"PRD2": {"BUY_ONE_GET_ONE_FREE"}, // TODO: change when more rules are implemented
	}

	testApp, err := NewCashRegisterApp(discountRules)
	assert.Nil(t, err)
	assert.Len(t, testApp.ShoppingCart.DiscountRules, 2)
}
