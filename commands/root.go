package commands

import (
	"fmt"
	"os"

	"github.com/jasonpilz/crypto-cli/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto",
	Short: "crypto is a command line interface for cryptocurrency prices and data",
	Long:  "crypto is a command line interface for ...",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func init() {
	config.Init()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(portfolioCmd)
	rootCmd.AddCommand(transactionCmd)
}
