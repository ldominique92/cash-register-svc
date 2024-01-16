package infrastructure_test

import (
	"cash-register-svc/internal/domain"
	"cash-register-svc/internal/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsInMemoryCache(t *testing.T) {
	product1 := domain.Product{
		Code:  "PRD1",
		Name:  "Coffee",
		Price: 10,
	}

	productsList := []domain.Product{product1}
	cache := infrastructure.NewProductsInMemoryCache()

	err := cache.Load(productsList)
	assert.Nil(t, err)

	product, err := cache.GetProduct("PRD1")
	assert.Nil(t, err)
	assert.Equal(t, product, product1)

	product, err = cache.GetProduct("PRD2")
	assert.Nil(t, err)
	assert.Equal(t, product.Code, domain.ProductCode(""))
}
