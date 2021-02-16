package transaction

import (
	"fmt"

	"github.com/jasonpilz/crypto-cli/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var NewCmd = &cobra.Command{
	Run:   runNewCmd,
	Use:   "new",
	Short: "New transaction",
	Long: `Use this command to create a new transaction.
Supports targeting specific portfolios using the --portfolio or -p argument.

Examples:
  crypto transaction new
  crypto transaction new --portfolio hodl`,
}

func runNewCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Using portfolio", viper.GetString("portfolio"))

	db := database.GetDb().
		LogMode(true)
	defer db.Close()

	transaction := database.Transaction{
		CoinName:   "BTC",
		CoinAmount: 0.01040429,
		CoinPrice:  9390.36,
		Fee:        2.30,
		Total:      100.00,
	}

	db.Create(&transaction)
}
