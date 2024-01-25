package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove products to your cart",
	Long:  `Usage: remove PRODUCT_CODE QUANTITY`,
	Run: func(cmd *cobra.Command, args []string) {
		if cashRegisterApp == nil {
			fmt.Println("not implemented")
			return
		}

		productCode, quantity, err := getProductCodeAndQuantityFromArgs(args)
		if err != nil {
			fmt.Println(err)
		}

		err = cashRegisterApp.AddProductToCart(productCode, -quantity)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	if RootCmd == nil {
		return
	}

	RootCmd.AddCommand(removeCmd)
}
