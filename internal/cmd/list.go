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
			fmt.Println("cart")
		case "discounts":
			fmt.Println("discounts")
		default:
			fmt.Println(message)
		}
	},
}

func init() {
	if rootCmd == nil {
		return
	}
	rootCmd.AddCommand(listCmd)
}

func listProducts() {
	if rootCmd == nil {
		fmt.Println("not implemented")
	}

	products := rootCmd.app.GetProducts()
	fmt.Println("|             Products                 |")
	fmt.Println("| Code | Name                 | Price  |")

	for _, p := range products {
		fmt.Printf("| %-4s | %-20s | %5.2f€ |\n", p.Code, p.Name, p.Price)
	}

}
