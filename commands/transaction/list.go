package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Run:   runListCmd,
	Use:   "list",
	Short: "List transactions",
	Long: `Use this command to view cryptocurrency transactions.
Supports targeting specific portfolios using the --portfolio or -p argument.

Examples:
  crypto transaction list
  crypto transaction list --portfolio hodl`,
}

func runListCmd(cmd *cobra.Command, args []string) {

	fmt.Println("Using portfolio", viper.GetString("portfolio"))
}
