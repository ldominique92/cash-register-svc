package infrastructure

import (
	"cash-register-svc/internal/domain"
	"encoding/json"
	"io"
	"os"
)

// TODO: should be a parameter so we can have a different file for testing
const ProductsSourceFile = "products.json"

type ProductFileRepository struct{}

func (r ProductFileRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product

	jsonFile, err := os.Open(ProductsSourceFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	fileBytes, err := io.ReadAll(jsonFile)
	err = json.Unmarshal(fileBytes, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
