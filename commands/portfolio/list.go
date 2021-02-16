package portfolio

import (
	"os"

	"github.com/jasonpilz/crypto-cli/database"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
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
	// TODO: Extract to pre-run
	db := database.GetDb().
		LogMode(true)
	defer db.Close()

	portfolios := []database.Portfolio{}
	db.Find(&portfolios)

	data := [][]string{}
	for _, p := range portfolios {
		data = append(data, []string{p.Name, p.Description})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
}

// TODO: Save for portfolio new command
// portfolio := database.Portfolio{Name: "hodl", Description: "life savings"}
// db.Create(&portfolio)
