package domain

import (
	"errors"

	"github.com/shopspring/decimal"
)

type ShoppingCart struct {
	Items map[ProductCode]ShoppingCartItem
}

func NewShoppingCart() ShoppingCart {
	return ShoppingCart{
		Items: make(map[ProductCode]ShoppingCartItem),
	}
}

func (c ShoppingCart) AddProduct(product Product, quantity int64) error {
	if quantity == 0 {
		return nil
	}

	if item, ok := c.Items[product.Code]; ok {
		newQuantity := item.Quantity + quantity
		if newQuantity < 0 {
			return errors.New("invalid quantity")
		}

		if newQuantity == 0 {
			delete(c.Items, product.Code)
			return nil
		}

		item.Quantity = newQuantity
		c.Items[product.Code] = item
	} else {
		c.Items[product.Code] = ShoppingCartItem{
			Product:  product,
			Quantity: quantity,
		}
	}

	return nil
}

func (c ShoppingCart) Total() (decimal.Decimal, error) {
	total := decimal.Zero

	for _, item := range c.Items {
		itemTotal, err := item.Total()
		if err != nil {
			return decimal.Zero, err
		}
		total = total.Add(itemTotal)
	}

	return total, nil
}

func (c ShoppingCart) Reset() {
	for k := range c.Items {
		delete(c.Items, k)
	}
}
