/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var rawtopicCmd = &cobra.Command{
	Use:   "rawtopic",
	Short: "Set of rawtopic commands",
	Long:  `Set of rawtopic commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rawtopic called")

		domainname, _ := cmd.Flags().GetString("domain")
		version, _ := cmd.Flags().GetString("major")

		fmt.Println("update-schema for raw topic " + domainname + "-V" + version)
	},
}

func init() {
	rootCmd.AddCommand(rawtopicCmd)

	rawtopicCmd.PersistentFlags().StringP("domain", "d", "", "The domainName for rawtopic subcommands.")
	rawtopicCmd.PersistentFlags().StringP("major", "j", "", "The majorVersion for the rawtopic subcommands.")

}
