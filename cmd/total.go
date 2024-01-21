package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// totalCmd represents the total command
var totalCmd = &cobra.Command{
	Use:   "total",
	Short: "Compute cart total",
	Long:  `Compute cart total`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("total called")
	},
}

func init() {
	rootCmd.AddCommand(totalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// totalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// totalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
