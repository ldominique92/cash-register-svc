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
			fmt.Println("products")
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
	rootCmd.AddCommand(listCmd)
}
