/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure ~/.harvest/cli.yml",
	Long: `Configure -m default and credentials to reduce number of cmdline arguments required.
Stores configuration in ~/.harvest/cli.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().StringP("key", "k", "", "The key for the key value set to add to the configuration.")
	configCmd.PersistentFlags().StringP("value", "v", "", "The value for the key value set to add to the configuration.")
}
