package domain

import (
	"errors"
	"fmt"
)

type ShoppingCart struct {
	Items         map[ProductCode]ShoppingCartItem
	DiscountRules map[ProductCode]DiscountRule
}

func NewShoppingCart(discountRules map[ProductCode]DiscountRuleName) (ShoppingCart, error) {
	cart := ShoppingCart{
		Items:         make(map[ProductCode]ShoppingCartItem),
		DiscountRules: make(map[ProductCode]DiscountRule),
	}

	for productCode, ruleName := range discountRules {
		if rule, ok := DiscountRules[ruleName]; ok {
			cart.DiscountRules[productCode] = rule
		} else {
			return ShoppingCart{}, errors.New(fmt.Sprintf("Discount rule %s not implemented", ruleName))
		}
	}

	return cart, nil
}

func (c ShoppingCart) AddProduct(product Product, quantity int64) {
	if quantity == 0 {
		return
	}

	if item, ok := c.Items[product.Code]; ok {
		item.Quantity += quantity
		c.Items[product.Code] = item
	} else {
		c.Items[product.Code] = ShoppingCartItem{
			Product:  product,
			Quantity: quantity,
		}
	}
}

// TODO: test this
func (c ShoppingCart) GetTotal() float64 {
	total := float64(0)

	for productCode, item := range c.Items {
		if rule, ok := c.DiscountRules[productCode]; ok {
			total += (float64(item.Quantity) * item.Product.Price) - rule.TotalDiscount(item.Quantity, item.Product.Price)
		} else {
			total += float64(item.Quantity) * item.Product.Price
		}
	}

	return total
}
