package domain

type ShoppingCart struct {
	Items         map[ProductCode]ShoppingCartItem
	DiscountRules map[ProductCode][]DiscountRule
}

func NewShoppingCart(discountRules []DiscountRule) {
}
