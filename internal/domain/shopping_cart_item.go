package domain

type ShoppingCartItem struct {
	Product      Product
	Quantity     int
	DiscountRule *DiscountRule
}

// TODO: test this
func (i ShoppingCartItem) Total() (float64, error) {
	if i.DiscountRule != nil {
		discount, err := i.DiscountRule.TotalDiscount(i.Quantity, i.Product.Price)
		if err != nil {
			return 0, err
		}
		return (float64(i.Quantity) * i.Product.Price) - discount, nil
	}
	return float64(i.Quantity) * i.Product.Price, nil
}
