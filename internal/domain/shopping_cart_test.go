package domain_test

import (
	"cash-register-svc/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCart_NewShoppingCart(t *testing.T) {
	cart := domain.NewShoppingCart()
	assert.NotNil(t, cart.Items)
	assert.Empty(t, cart.Items)
}

func TestShoppingCart_AddProduct(t *testing.T) {
	cart := domain.NewShoppingCart()

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

	err := cart.AddProduct(product1, 2)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 1)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 2)

	err = cart.AddProduct(product1, 1)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 1)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 3)

	err = cart.AddProduct(product2, 5)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 3)
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, 5)

	err = cart.AddProduct(product2, -2)
	assert.Nil(t, err)
	assert.NotEmpty(t, cart.Items)
	assert.Equal(t, len(cart.Items), 2)
	assert.Equal(t, cart.Items["PRD1"].Product, product1)
	assert.Equal(t, cart.Items["PRD1"].Quantity, 3)
	assert.Equal(t, cart.Items["PRD2"].Product, product2)
	assert.Equal(t, cart.Items["PRD2"].Quantity, 3)
	// TODO: quantity can't be negative
}

func TestShoppingCart_Total(t *testing.T) {
	greenTeaDiscountRule := domain.DiscountRule{
		MinimumQuantity:      2,
		IsAppliedToBatches:   true,
		BatchSize:            2,
		IsPercentageDiscount: true,
		DiscountPercentage:   1,
		DiscountInEuro:       0,
	}
	strawberriesDiscountRule := domain.DiscountRule{
		MinimumQuantity:      3,
		IsAppliedToBatches:   false,
		BatchSize:            0,
		IsPercentageDiscount: false,
		DiscountPercentage:   0,
		DiscountInEuro:       0.50,
	}
	coffeeDiscountRule := domain.DiscountRule{
		MinimumQuantity:      3,
		IsAppliedToBatches:   false,
		BatchSize:            0,
		IsPercentageDiscount: true,
		DiscountPercentage:   float64(1) / float64(3),
		DiscountInEuro:       0,
	}

	cart := domain.NewShoppingCart()

	err := cart.AddProduct(domain.Product{
		Code:         "GR1",
		Name:         "Green Tea",
		Price:        3.11,
		DiscountRule: greenTeaDiscountRule,
	}, 2)
	assert.Nil(t, err)
	total, err := cart.Total()
	assert.Equal(t, total, 3.11)
	assert.Nil(t, err)

	cart.Reset() // TODO: test

	err = cart.AddProduct(domain.Product{
		Code:         "SR1",
		Name:         "Strawberries",
		Price:        5.00,
		DiscountRule: strawberriesDiscountRule,
	}, 3)
	assert.Nil(t, err)

	err = cart.AddProduct(domain.Product{
		Code:         "GR1",
		Name:         "Green Tea",
		Price:        3.11,
		DiscountRule: greenTeaDiscountRule,
	}, 1)
	assert.Nil(t, err)

	total, err = cart.Total()
	assert.Equal(t, total, 16.61)
	assert.Nil(t, err)

	cart.Reset()

	err = cart.AddProduct(domain.Product{
		Code:         "GR1",
		Name:         "Green Tea",
		Price:        3.11,
		DiscountRule: greenTeaDiscountRule,
	}, 1)
	assert.Nil(t, err)

	err = cart.AddProduct(domain.Product{
		Code:         "CF1",
		Name:         "Coffee",
		Price:        11.23,
		DiscountRule: coffeeDiscountRule,
	}, 3)
	assert.Nil(t, err)

	err = cart.AddProduct(domain.Product{
		Code:         "SR1",
		Name:         "Strawberries",
		Price:        5.00,
		DiscountRule: strawberriesDiscountRule,
	}, 1)
	assert.Nil(t, err)

	total, err = cart.Total()
	assert.Equal(t, total, 30.57)
	assert.Nil(t, err)
}
