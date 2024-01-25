package domain

type ShoppingCartItem struct {
	Product  Product
	Quantity int
}

// TODO: test this
func (i ShoppingCartItem) Total() (float64, error) {
	discount, err := i.TotalDiscount()
	if err != nil {
		return 0, err
	}
	return (float64(i.Quantity) * i.Product.Price) - discount, nil
}

func (i ShoppingCartItem) TotalDiscount() (float64, error) {
	return i.Product.DiscountRule.TotalDiscount(i.Quantity, i.Product.Price)
}
