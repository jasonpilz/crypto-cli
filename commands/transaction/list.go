package transaction

import (
	"fmt"
	// "os"

	// "github.com/jasonpilz/crypto-cli/database"
	// "github.com/olekukonko/tablewriter"
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

	// db := database.GetDb().
	// 	LogMode(true)
	// defer db.Close()

	// portfolios := []database.Portfolio{}
	// db.Find(&portfolios)

	// data := [][]string{}
	// for _, p := range portfolios {
	// 	data = append(data, []string{p.Name, p.Description})
	// }

	// table := tablewriter.NewWriter(os.Stdout)
	// table.SetHeader([]string{"Name", "Description"})
	// table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	// table.SetCenterSeparator("|")
	// table.AppendBulk(data)
	// table.Render()
}
