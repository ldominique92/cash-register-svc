package domain

import (
	"errors"
	"fmt"
)

type ShoppingCart struct {
	Items         map[ProductCode]ShoppingCartItem
	DiscountRules map[ProductCode][]DiscountRule
}

func NewShoppingCart(discountRules map[ProductCode][]DiscountRuleName) (ShoppingCart, error) {
	cart := ShoppingCart{
		Items:         make(map[ProductCode]ShoppingCartItem),
		DiscountRules: make(map[ProductCode][]DiscountRule),
	}

	for productCode, ruleNames := range discountRules {
		var rules []DiscountRule
		for _, ruleName := range ruleNames {
			if rule, ok := DiscountRules[ruleName]; ok {
				rules = append(rules, rule)
			} else {
				return ShoppingCart{}, errors.New(fmt.Sprintf("Discount rule %s not implemented", ruleName))
			}

			cart.DiscountRules[productCode] = rules
		}
	}

	return cart, nil
}

func (c ShoppingCart) AddProduct(product Product, quantity int) {
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
