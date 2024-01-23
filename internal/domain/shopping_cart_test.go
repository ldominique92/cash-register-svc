package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCart_NewShoppingCart(t *testing.T) {
	rule := domain.DiscountRule{}
	testRules := map[domain.ProductCode]domain.DiscountRule{
		"PRD1": rule,
		"PRD2": rule,
	}

	cart, err := domain.NewShoppingCart(testRules)
	assert.Nil(t, err)
	assert.NotNil(t, cart.Items)
	assert.Empty(t, cart.Items)
	assert.NotNil(t, cart.DiscountRules)
	assert.Len(t, cart.DiscountRules, 2)
	assert.Equal(t, cart.DiscountRules["PRD1"], rule)
	assert.Equal(t, cart.DiscountRules["PRD2"], rule)
}

func TestShoppingCart_AddProduct(t *testing.T) {
	rule := domain.DiscountRule{}
	testRules := map[domain.ProductCode]domain.DiscountRule{"PRD1": rule}
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

	err = cart.AddProduct(product1, 2)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 1)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(2))

	err = cart.AddProduct(product1, 1)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 1)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(3))

	err = cart.AddProduct(product2, 5)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(3))
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, int64(5))

	err = cart.AddProduct(product2, -2)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, int64(3))
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, int64(3))
	// TODO: quantity can't be negative
}

func TestShoppingCart_GetTotal(t *testing.T) {
	discountRules := map[domain.ProductCode]domain.DiscountRule{
		"GR1": {
			MinimumQuantity:      2,
			IsAppliedToBatches:   true,
			BatchSize:            2,
			IsPercentageDiscount: true,
			DiscountPercentage:   0.50,
			DiscountInEuro:       0,
		},
		"SR1": {
			MinimumQuantity:      3,
			IsAppliedToBatches:   false,
			BatchSize:            0,
			IsPercentageDiscount: false,
			DiscountPercentage:   0,
			DiscountInEuro:       0.50,
		},
		"CF1": {
			MinimumQuantity:      3,
			IsAppliedToBatches:   false,
			BatchSize:            0,
			IsPercentageDiscount: true,
			DiscountPercentage:   float64(1) / float64(3),
			DiscountInEuro:       0,
		},
	}
	cart, err := domain.NewShoppingCart(discountRules)
	assert.Nil(t, err)

	err = cart.AddProduct(domain.Product{
		Code:  "GR1",
		Name:  "Green Tea",
		Price: 3.11,
	}, 2)
	assert.Nil(t, err)
	total, err := cart.GetTotal()
	assert.Equal(t, total, 3.11)
	assert.Nil(t, err)

	cart.Reset() // TODO: test

	err = cart.AddProduct(domain.Product{
		Code:  "SR1",
		Name:  "Strawberries",
		Price: 5.00,
	}, 3)
	assert.Nil(t, err)

	err = cart.AddProduct(domain.Product{
		Code:  "GR1",
		Name:  "Green Tea",
		Price: 3.11,
	}, 1)
	assert.Nil(t, err)

	total, err = cart.GetTotal()
	assert.Equal(t, total, 16.61)
	assert.Nil(t, err)

	cart.Reset()

	err = cart.AddProduct(domain.Product{
		Code:  "GR1",
		Name:  "Green Tea",
		Price: 3.11,
	}, 1)
	assert.Nil(t, err)

	err = cart.AddProduct(domain.Product{
		Code:  "CF1",
		Name:  "Coffee",
		Price: 11.23,
	}, 3)
	assert.Nil(t, err)

	err = cart.AddProduct(domain.Product{
		Code:  "SR1",
		Name:  "Strawberries",
		Price: 5.00,
	}, 1)
	assert.Nil(t, err)

	total, err = cart.GetTotal()
	assert.Equal(t, total, 30.57)
	assert.Nil(t, err)
}
