package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCart_NewShoppingCart(t *testing.T) {
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

func TestShoppingCart_AddProduct(t *testing.T) {
	testRules := map[domain.ProductCode][]domain.DiscountRuleName{"PRD1": {"BUY_ONE_GET_ONE_FREE"}}
	cart, err := domain.NewShoppingCart(testRules)
	assert.Nil(t, err)

	product1 := domain.Product{
		Code:  "PRD1",
		Name:  "Coffee",
		Price: 10,
	}

	product2 := domain.Product{
		Code:  "PRD2",
		Name:  "Coffee",
		Price: 15,
	}

	assert.Empty(t, cart.Items)

	cart.AddProduct(product1, 2)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 1)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 2)

	cart.AddProduct(product1, 1)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 1)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 3)

	cart.AddProduct(product2, 5)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 3)
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, 5)

	cart.AddProduct(product2, -2)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 3)
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, 3)
}
