package domain

import (
	"errors"
)

type ShoppingCart struct {
	Items map[ProductCode]ShoppingCartItem
}

func NewShoppingCart() ShoppingCart {
	return ShoppingCart{
		Items: make(map[ProductCode]ShoppingCartItem),
	}
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
			c.Items[product.Code] = ShoppingCartItem{
				Product:  product,
				Quantity: quantity,
			}
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
