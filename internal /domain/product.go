package domain

type ProductCode string

type Product struct {
	Code  ProductCode
	Name  string
	Price float64
}