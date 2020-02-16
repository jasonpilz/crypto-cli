package commands

import (
	"fmt"

	"github.com/jasonpilz/crypto-cli/spec"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version",
	Long:  fmt.Sprintf("Show the current version of %s", spec.Repo),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(spec.AppVersion.Complete(&spec.GithubLatestVersioner{}))
	},
}
