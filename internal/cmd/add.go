package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add products to your cart",
	Long: `Use the flags for:
-p PRODUCT_NAME 
-q QUANTITY`,
	Run: func(cmd *cobra.Command, args []string) {
		if cashRegisterApp == nil {
			fmt.Println("not implemented")
			return
		}

		productCode, _ := cmd.Flags().GetString("p")
		quantity, _ := cmd.Flags().GetInt("q")

		if len(productCode) == 0 {
			fmt.Println("product code is mandatory")
			return
		}

		if quantity <= 0 {
			fmt.Println("product quantity should be bigger than 0")
			return
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

	addCmd.Flags().String("p", "", "Product code")
	addCmd.Flags().Int("q", 0, "Product quantity")
}
