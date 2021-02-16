package commands

import (
	"github.com/jasonpilz/crypto-cli/commands/transaction"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Manage transactions",
	Long: `Root command for managing cryptocurrency transactions.
Requires use of sub-command.

Examples:
  crypto transaction new
  crypto transaction list`,
}

func init() {
	transactionCmd.AddCommand(transaction.ListCmd)
	transactionCmd.AddCommand(transaction.NewCmd)

	transactionCmd.PersistentFlags().StringP("portfolio", "p", "default", "transaction portfolio")

	viper.BindPFlag("portfolio", transactionCmd.PersistentFlags().Lookup("portfolio"))
}
