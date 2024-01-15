package domain

const AllProducts = "*"

type DiscountRule interface {
	Name() string
	AppliesTo() []ProductCode
	TotalDiscount() float64
}
