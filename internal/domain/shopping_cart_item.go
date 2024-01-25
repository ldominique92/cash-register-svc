package domain

import "github.com/shopspring/decimal"

type ShoppingCartItem struct {
	Product  Product
	Quantity int64
}

// TODO: test this
func (i ShoppingCartItem) Total() (decimal.Decimal, error) {
	discount, err := i.TotalDiscount()
	if err != nil {
		return decimal.Zero, err
	}
	return decimal.NewFromInt(i.Quantity).Mul(i.Product.Price).Sub(discount), nil
}

func (i ShoppingCartItem) TotalDiscount() (decimal.Decimal, error) {
	return i.Product.DiscountRule.TotalDiscount(i.Quantity, i.Product.Price)
}
