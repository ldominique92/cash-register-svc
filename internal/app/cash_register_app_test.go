package app_test

import (
	"cash-register-svc/internal/app"
	"cash-register-svc/internal/domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCashRegisterApp_NewCashRegisterApp(t *testing.T) {
	productRepositoryMock := new(ProductRepositoryMock)
	productCacheMock := new(ProductCacheMock)

	// Case 1: repository throws error to fetch products
	discountRules := map[string]string{"PRD1": "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE"}

	repositoryMockCall := productRepositoryMock.On("GetProducts").
		Return([]domain.Product{}, errors.New("unexpected repo error"))

	testApp, err := app.NewCashRegisterApp(productRepositoryMock, productCacheMock, discountRules)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unexpected repo error")
	productRepositoryMock.AssertExpectations(t)
	repositoryMockCall.Unset()

	// Case 2 : repository returns list of products but cache returns error on loading
	product := domain.Product{
		Code:  "PRD1",
		Name:  "Coffee",
		Price: 10,
	}
	productList := []domain.Product{product}
	productRepositoryMock.On("GetProducts").Return(productList, nil)
	cacheMockCall := productCacheMock.On("Load", productList).Return(errors.New("unexpected cache error"))

	testApp, err = app.NewCashRegisterApp(productRepositoryMock, productCacheMock, discountRules)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unexpected cache error")

	productRepositoryMock.AssertExpectations(t)
	productCacheMock.AssertExpectations(t)
	cacheMockCall.Unset()

	// Case 3: cache loaded successfully, but cache return error on checking if product exists
	productCacheMock.On("Load", productList).Return(nil)
	cacheMockCall = productCacheMock.On("GetProduct", domain.ProductCode("PRD1")).Return(domain.Product{}, errors.New("unexpected cache error"))

	testApp, err = app.NewCashRegisterApp(productRepositoryMock, productCacheMock, discountRules)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unexpected cache error")

	productRepositoryMock.AssertExpectations(t)
	productCacheMock.AssertExpectations(t)
	cacheMockCall.Unset()

	// Case 4: discount rule contains unknown product
	discountRules = map[string]string{
		"XPTO": "INVALID",
	}

	cacheMockCall = productCacheMock.On("GetProduct", domain.ProductCode("XPTO")).Return(domain.Product{}, nil)

	testApp, err = app.NewCashRegisterApp(productRepositoryMock, productCacheMock, discountRules)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "product with code XPTO not found")

	productRepositoryMock.AssertExpectations(t)
	productCacheMock.AssertExpectations(t)
	cacheMockCall.Unset()

	// Case 5: invalid discount rules
	discountRules = map[string]string{
		"PRD1": "INVALID",
	}
	cacheMockCall = productCacheMock.On("GetProduct", domain.ProductCode("PRD1")).Return(product, nil)

	testApp, err = app.NewCashRegisterApp(productRepositoryMock, productCacheMock, discountRules)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Discount rule INVALID not implemented")

	productRepositoryMock.AssertExpectations(t)
	productCacheMock.AssertExpectations(t)

	// Case 5 : valid discount rules
	discountRules = map[string]string{
		"PRD1": "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE",
	}
	testApp, err = app.NewCashRegisterApp(productRepositoryMock, productCacheMock, discountRules)
	assert.Nil(t, err)
	assert.Len(t, testApp.ShoppingCart.DiscountRules, 1)

	productRepositoryMock.AssertExpectations(t)
	productCacheMock.AssertExpectations(t)
}

func TestCashRegisterApp_AddProductToCart(t *testing.T) {
	productCacheMock := new(ProductCacheMock)
	productRepositoryMock := new(ProductRepositoryMock)
	productCode := "PRD1"
	discountRules := map[string]string{"PRD1": "BUY_ONE_GET_ONE_FREE_DISCOUNT_RULE"}
	product := domain.Product{
		Code:  "PRD1",
		Name:  "Coffee",
		Price: 10,
	}
	productList := []domain.Product{product}
	productRepositoryMock.On("GetProducts").Return(productList, nil)
	productCacheMock.On("Load", productList).Return(nil)
	testApp, err := app.NewCashRegisterApp(productRepositoryMock, productCacheMock, discountRules)

	// Case 1: cache returns error
	cacheMockCall := productCacheMock.
		On("GetProduct", domain.ProductCode(productCode)).
		Return(domain.Product{}, errors.New("unexpected cache error"))

	err = testApp.AddProductToCart(productCode, 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unexpected cache error")
	productCacheMock.AssertExpectations(t)
	cacheMockCall.Unset()

	// Case 2: cache returns empty product
	cacheMockCall = productCacheMock.
		On("GetProduct", domain.ProductCode(productCode)).
		Return(domain.Product{}, nil)

	err = testApp.AddProductToCart(productCode, 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "product with code PRD1 not found")
	productCacheMock.AssertExpectations(t)
	cacheMockCall.Unset()

	// Case 3: product is added
	productCacheMock.On("GetProduct", domain.ProductCode(productCode)).Return(product, nil)
	err = testApp.AddProductToCart(productCode, 1)
	assert.Nil(t, err)
	assert.Equal(t, testApp.ShoppingCart.Items[domain.ProductCode(productCode)].Product, product)
	productCacheMock.AssertExpectations(t)
}

type ProductRepositoryMock struct {
	mock.Mock
}

func (r ProductRepositoryMock) GetProducts() ([]domain.Product, error) {
	args := r.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

type ProductCacheMock struct {
	mock.Mock
}

func (c ProductCacheMock) Load(products []domain.Product) error {
	args := c.Called(products)
	return args.Error(0)
}

func (c ProductCacheMock) GetProduct(code domain.ProductCode) (domain.Product, error) {
	args := c.Called(code)
	return args.Get(0).(domain.Product), args.Error(1)
}
