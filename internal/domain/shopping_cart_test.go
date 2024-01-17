package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCart_NewShoppingCart(t *testing.T) {
	// Case 1: When rules are invalid
	testRules := map[domain.ProductCode]domain.DiscountRuleName{"PRD1": "TEST_RULE"}
	cart, err := domain.NewShoppingCart(testRules)
	assert.NotNil(t, err)
	assert.Nil(t, cart.Items)
	assert.Nil(t, cart.DiscountRules)

	// Case 2: When rules valid
	testRules = map[domain.ProductCode]domain.DiscountRuleName{
		"PRD1": "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE",
		"PRD2": "CASH_BULK_DISCOUNT_RULE",
	}

	cart, err = domain.NewShoppingCart(testRules)
	assert.Nil(t, err)
	assert.NotNil(t, cart.Items)
	assert.Empty(t, cart.Items)
	assert.NotNil(t, cart.DiscountRules)
	assert.Len(t, cart.DiscountRules, 2)
	assert.Equal(t, cart.DiscountRules["PRD1"].Name(), "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE")
	assert.Equal(t, cart.DiscountRules["PRD2"].Name(), "CASH_BULK_DISCOUNT_RULE")
}

func TestShoppingCart_AddProduct(t *testing.T) {
	testRules := map[domain.ProductCode]domain.DiscountRuleName{"PRD1": "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE"}
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
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(2))

	cart.AddProduct(product1, 1)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 1)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(3))

	cart.AddProduct(product2, 5)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(3))
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, int64(5))

	cart.AddProduct(product2, -2)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(3))
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, int64(3))
	// TODO: quantity can't be negative
}

func TestShoppingCart_GetTotal(t *testing.T) {
	discountRules := map[domain.ProductCode]domain.DiscountRuleName{
		"GR1": "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE",
		"SR1": "CASH_BULK_DISCOUNT_RULE",
		"CF1": "PERCENTAGE_BULK_DISCOUNT_RULE",
	}
	cart, err := domain.NewShoppingCart(discountRules)
	assert.Nil(t, err)

	cart.AddProduct(domain.Product{
		Code:  "GR1",
		Name:  "Green Tea",
		Price: 3.11,
	}, 2)
	assert.Equal(t, cart.GetTotal(), 3.11)

	cart.Reset() // TODO: test

	cart.AddProduct(domain.Product{
		Code:  "SR1",
		Name:  "Strawberries",
		Price: 5.00,
	}, 3)
	cart.AddProduct(domain.Product{
		Code:  "GR1",
		Name:  "Green Tea",
		Price: 3.11,
	}, 1)
	assert.Equal(t, cart.GetTotal(), 16.61)

	cart.Reset()

	cart.AddProduct(domain.Product{
		Code:  "GR1",
		Name:  "Green Tea",
		Price: 3.11,
	}, 1)
	cart.AddProduct(domain.Product{
		Code:  "CF1",
		Name:  "Coffee",
		Price: 11.23,
	}, 3)
	cart.AddProduct(domain.Product{
		Code:  "SR1",
		Name:  "Strawberries",
		Price: 5.00,
	}, 1)
	assert.Equal(t, cart.GetTotal(), 30.57)

}
