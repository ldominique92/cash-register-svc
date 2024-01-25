package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available products or the content of your cart",
	Long: `Use one of the following parameters to determine what you want to list:
- products
- cart`,
	Run: func(cmd *cobra.Command, args []string) {
		message := "list requires a parameter: products or cart"
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
	fmt.Println("| Products                                                                                                        |")
	fmt.Println("| Code | Name                 | Price  | Discount                                                                 |")

	for _, p := range products {
		fmt.Printf("| %-4s | %-20s | %5.2f€ | %-72s |\n", p.Code, p.Name, p.Price, p.DiscountRule.Description())
	}

}

func listCart() {
	if cashRegisterApp == nil {
		fmt.Println("not implemented")
	}

	fmt.Println("| Cart                                                                |")
	fmt.Println("| Code | Name                 | Quantity | Price  | Discount | Total  | ")

	for _, i := range cashRegisterApp.ShoppingCart.Items {
		discount, err := i.TotalDiscount()
		if err != nil {
			fmt.Println("error calculating item total")
		}

		total, err := i.Total()
		if err != nil {
			fmt.Println("error calculating item total")
		}

		fmt.Printf(
			"| %-4s | %-20s | %-8d | %5.2f€ | %7.2f€ | %5.2f€ |\n",
			i.Product.Code,
			i.Product.Name,
			i.Quantity,
			i.Product.Price,
			discount,
			total,
		)
	}

	total, err := cashRegisterApp.ShoppingCart.Total()
	if err != nil {
		fmt.Println("error calculating cart total")
	}
	fmt.Printf("| Total                                                        %5.2f€ |\n", total)
}
