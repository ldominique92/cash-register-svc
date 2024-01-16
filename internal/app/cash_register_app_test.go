package app_test

import (
	"cash-register-svc/internal/app"
	"cash-register-svc/internal/domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r ProductRepositoryMock) GetProducts() ([]domain.Product, error) {
	args := r.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func TestNewCashRegisterApp(t *testing.T) {
	productRepositoryMock := new(ProductRepositoryMock)

	// With invalid discount rules
	discountRules := map[string][]string{
		"PRD1": {"INVALID"},
	}
	testApp, err := app.NewCashRegisterApp(productRepositoryMock, discountRules)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Discount rule INVALID not implemented")

	// With valid discount rules
	discountRules = map[string][]string{
		"PRD1": {"BUY_ONE_GET_ONE_FREE"},
		"PRD2": {"BUY_ONE_GET_ONE_FREE"}, // TODO: change when more rules are implemented
	}

	// When repository returns error to get products
	mockCall := productRepositoryMock.On("GetProducts").Return([]domain.Product{}, errors.New("unexpected error"))
	testApp, err = app.NewCashRegisterApp(productRepositoryMock, discountRules)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unexpected error")
	productRepositoryMock.AssertExpectations(t)
	mockCall.Unset()

	// When repository returns list of products
	productRepositoryMock.On("GetProducts").Return([]domain.Product{}, nil)
	testApp, err = app.NewCashRegisterApp(productRepositoryMock, discountRules)
	assert.Nil(t, err)
	assert.Len(t, testApp.ShoppingCart.DiscountRules, 2)
	productRepositoryMock.AssertExpectations(t)
}
