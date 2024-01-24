package domain

import (
	"errors"
)

type ShoppingCart struct {
	Items         map[ProductCode]ShoppingCartItem
	DiscountRules map[ProductCode]DiscountRule
}

func NewShoppingCart(discountRules map[ProductCode]DiscountRule) (ShoppingCart, error) {
	cart := ShoppingCart{
		Items:         make(map[ProductCode]ShoppingCartItem),
		DiscountRules: discountRules,
	}

	return cart, nil
}

func (c ShoppingCart) AddProduct(product Product, quantity int) error {
	if quantity != 0 {
		if item, ok := c.Items[product.Code]; ok {
			if quantity < 0 && item.Quantity < (-1*quantity) {
				return errors.New("invalid quantity")
			}
			item.Quantity += quantity
			c.Items[product.Code] = item
		} else {
			item := ShoppingCartItem{
				Product:  product,
				Quantity: quantity,
			}
			if rule, ok := c.DiscountRules[product.Code]; ok {
				item.DiscountRule = &rule
			}
			c.Items[product.Code] = item
		}
	}

	return nil
}

func (c ShoppingCart) Total() (float64, error) {
	total := float64(0)

	for _, item := range c.Items {
		itemTotal, err := item.Total()
		if err != nil {
			return 0, err
		}
		total += itemTotal
	}

	return total, nil
}

func (c ShoppingCart) Reset() {
	for k := range c.Items {
		delete(c.Items, k)
	}
}
