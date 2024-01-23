package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available products, discounts or the content of your cart",
	Long: `Use one of the following parameters to determine what you want to list:
- products
- discounts
- cart`,
	Run: func(cmd *cobra.Command, args []string) {
		message := "list requires a parameter: products, discounts or cart"
		if len(args) < 1 {
			fmt.Println(message)
			return
		}

		p := args[0]
		switch p {
		case "products":
			listProducts()
		case "cart":
			listCart()
		case "discounts":
			listDiscounts()
		default:
			fmt.Println(message)
		}
	},
}

func init() {
	if RootCmd == nil {
		return
	}
	RootCmd.AddCommand(listCmd)
}

func listProducts() {
	if cashRegisterApp == nil {
		fmt.Println("not implemented")
	}

	products := cashRegisterApp.GetProducts()
	fmt.Println("|             Products                 |")
	fmt.Println("| Code | Name                 | Price  |")

	for _, p := range products {
		fmt.Printf("| %-4s | %-20s | %5.2f€ |\n", p.Code, p.Name, p.Price)
	}

}

func listCart() {
	if cashRegisterApp == nil {
		fmt.Println("not implemented")
	}

	fmt.Println("|                     Cart                        |")
	fmt.Println("| Code | Name                 | Quantity | Price  |")

	for _, i := range cashRegisterApp.ShoppingCart.Items { // TODO: create getter
		fmt.Printf("| %-4s | %-20s | %2d | %5.2f€ |\n", i.Product.Code, i.Product.Name, i.Quantity, i.Product.Price)
	}

}

func listDiscounts() {
	if cashRegisterApp == nil {
		fmt.Println("not implemented")
	}

	fmt.Println("|                     Discount                       |")
	fmt.Println("| Product | Description                              |")

	for p, r := range cashRegisterApp.ShoppingCart.DiscountRules {
		fmt.Printf("| %-4s | %-40s |\n", p, r.Description())
	}

}
