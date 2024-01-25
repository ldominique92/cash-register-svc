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
		fmt.Println("reset called")
	},
}

func init() {
	if RootCmd == nil {
		return
	}

	RootCmd.AddCommand(resetCmd)
}
