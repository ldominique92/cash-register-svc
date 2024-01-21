package infrastructure

import (
	"cash-register-svc/internal/domain"
	"encoding/json"
	"io"
	"os"
)

type ProductFileRepository struct {
	sourceFile string
}

func NewProductFileRepository(sourceFile string) ProductFileRepository {
	return ProductFileRepository{
		sourceFile: sourceFile,
	}
}

func (r ProductFileRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product

	jsonFile, err := os.Open(r.sourceFile)
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
