package cmd

import (
	configMgmt "github.com/FinarxGmbH/harvest-cli/configMgt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "The 'add' subcommand will add a passed in key value pair to the application configuration file.",
	Long: `The 'add' subcommand adds a key value pair to the application configuration file. For example:
'<cmd> config add --key theKey --value "the value can be a bunch of things."'.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		configMgmt.ConfigKeyValuePairAdd(key, value)
	},
}

func init() {
	configCmd.AddCommand(addCmd)
}
