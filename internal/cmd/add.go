package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add products to your cart",
	Long:  `Usage: add PRODUCT_CODE QUANTITY`,
	Run: func(cmd *cobra.Command, args []string) {
		if cashRegisterApp == nil {
			fmt.Println("not implemented")
			return
		}

		if len(args) < 1 {
			fmt.Println("product code is mandatory")
			return
		}
		productCode := args[0]

		quantity := 1
		if len(args) >= 2 {
			if q, err := strconv.Atoi(args[1]); err == nil {
				quantity = q
				if quantity <= 0 {
					fmt.Println("product quantity should be bigger than 0")
					return
				}
			} else {
				fmt.Println("quantity should be an integer")
				return
			}
		}

		err := cashRegisterApp.AddProductToCart(productCode, quantity)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	if RootCmd == nil {
		fmt.Println("not implemented")
		return
	}

	RootCmd.AddCommand(addCmd)
}
