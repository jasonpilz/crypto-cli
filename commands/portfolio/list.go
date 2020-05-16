package portfolio

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Run:   runListCmd,
	Use:   "list",
	Short: "List portfolios",
	Long: `Use this command to view transaction portfolios.

Examples:
  crypto portfolio list`,
}

func runListCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Called portfolio list")
}
