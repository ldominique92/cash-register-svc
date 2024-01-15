package domain

type DiscountRule interface {
	Name() string
	AppliesTo() ProductCode
	TotalDiscount() float64
}
