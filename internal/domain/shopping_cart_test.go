package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShoppingCart(t *testing.T) {
	// With invalid rules
	testRules := map[domain.ProductCode][]domain.DiscountRuleName{
		"PRD1": {"TEST_RULE"},
	}

	cart, err := domain.NewShoppingCart(testRules)
	assert.NotNil(t, err)
	assert.Nil(t, cart.Items)
	assert.Nil(t, cart.DiscountRules)

	// With valid rules
	testRules = map[domain.ProductCode][]domain.DiscountRuleName{
		"PRD1": {"BUY_ONE_GET_ONE_FREE"},
		"PRD2": {"BUY_ONE_GET_ONE_FREE"}, // TODO: add new rules when implemented
	}

	cart, err = domain.NewShoppingCart(testRules)
	assert.Nil(t, err)
	assert.NotNil(t, cart.Items)
	assert.Empty(t, cart.Items)
	assert.NotNil(t, cart.DiscountRules)
	assert.Len(t, cart.DiscountRules, 2)
	assert.Equal(t, cart.DiscountRules["PRD1"][0].Name(), "BUY_ONE_GET_ONE_FREE")
	assert.Equal(t, cart.DiscountRules["PRD2"][0].Name(), "BUY_ONE_GET_ONE_FREE")
}
