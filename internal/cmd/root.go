package cmd

import (
	"os"

	"cash-register-svc/internal/app"

	"github.com/spf13/cobra"
)

type CashRegisterRootCommand struct {
	cmd             cobra.Command
	cashRegisterApp app.CashRegisterApp
}

func (c *CashRegisterRootCommand) Execute() error {
	return c.cmd.Execute()
}

func (c *CashRegisterRootCommand) AddCommand(cmd *cobra.Command) {
	c.cmd.AddCommand(cmd)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd *CashRegisterRootCommand

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(cashRegisterApp app.CashRegisterApp) {
	rootCmd = &CashRegisterRootCommand{
		cmd: cobra.Command{
			Use:   "cash-register-svc",
			Short: "###  Cash Register Amenitiz ###",
			Long:  `###  Cash Register Amenitiz ###`,
		},
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
