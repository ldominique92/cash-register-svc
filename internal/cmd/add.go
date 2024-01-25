package cmd

import (
	"errors"
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

		productCode, quantity, err := getProductCodeAndQuantityFromArgs(args)
		if err != nil {
			fmt.Println(err)
		}

		err = cashRegisterApp.AddProductToCart(productCode, quantity)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func getProductCodeAndQuantityFromArgs(args []string) (string, int64, error) {
	if len(args) < 1 {
		return "", 0, errors.New("product code is mandatory")
	}
	productCode := args[0]

	quantity := int64(1)
	if len(args) >= 2 {
		if q, err := strconv.Atoi(args[1]); err == nil {
			quantity = int64(q)
			if quantity <= 0 {
				return "", 0, errors.New("product quantity should be bigger than 0")
			}
		} else {
			return "", 0, errors.New("quantity should be an integer")
		}
	}
	return productCode, quantity, nil
}

func init() {
	if RootCmd == nil {
		fmt.Println("not implemented")
		return
	}

	RootCmd.AddCommand(addCmd)
}
