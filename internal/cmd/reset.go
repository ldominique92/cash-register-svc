package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset shopping cart",
	Long:  `Reset shopping cart`,
	Run: func(cmd *cobra.Command, args []string) {
		if cashRegisterApp == nil {
			fmt.Println("not implemented")
			return
		}

		cashRegisterApp.ResetShoppingCart()
	},
}

func init() {
	if RootCmd == nil {
		return
	}

	RootCmd.AddCommand(resetCmd)
}
