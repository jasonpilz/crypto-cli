package commands

import (
	"github.com/jasonpilz/crypto-cli/config"
	"github.com/jasonpilz/crypto-cli/prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Run:   runConfigureCmd,
	Use:   "configure",
	Short: "Manage configuration settings",
	Long: `Use this command to create an initial configuration file and assign configuration values.
Can be used to update/overwrite existing values if config previously exists.

Configuration Settings:
  CMC API Key - The API key used to retreive crypto data on-demand via Coin Market Cap.
              - See https://pro.coinmarketcap.com/signup/

Examples:
  crypto configure`,
}

func runConfigureCmd(cmd *cobra.Command, args []string) {
	config.Init()

	// Set Coin Market Cap API Key
	var cmcKey string
	existingKey := viper.GetString("CMC_PRO_API_KEY")
	if existingKey != "" {
		cmcKey = prompt.LinerPromptSuggest("CMC API Key: ", existingKey)
	} else {
		cmcKey = prompt.LinerPrompt("CMC API Key: ")
	}
	viper.Set("CMC_PRO_API_KEY", cmcKey)

	viper.WriteConfig()
}
