package cmd

import (
	"fmt"

	configMgmt "github.com/FinarxGmbH/harvest-cli/configMgt"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "The 'delete' subcommand removes a key value pair from the configuration file.",
	Long:  `The 'delete' subcommand removes a key value pair from the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		fmt.Printf("\n\n    **** Deleting key: %s ****\n\n", key)
		configMgmt.ConfigKeyValuePairDelete(key)
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)
}
