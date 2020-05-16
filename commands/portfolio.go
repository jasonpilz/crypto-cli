package commands

import (
	"github.com/jasonpilz/crypto-cli/commands/portfolio"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Manage portfolios",
	Long: `Root command for managing transaction portfolios.
Requires use of sub-command.

Examples:
  crypto portfolio list`,
}

func init() {
	portfolioCmd.AddCommand(portfolio.ListCmd)
}
