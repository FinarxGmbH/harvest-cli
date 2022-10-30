/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a default configuration",
	Long:  `Create a config file with vars commented out`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create config file " + viper.GetViper().ConfigFileUsed())

		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		} else {
			configdir := home + "/.harvest"
			err := os.MkdirAll(configdir, 0700)
			if err != nil {
				log.Fatal(err)
			} else {
				configfile := configdir + "/cli.yaml"
				data := []byte(`# config file generated
mcpHost=http://localhost:7877/
authToken=somesecrettoken
`)
				os.WriteFile(configfile, data, 0600)
				fmt.Println(configfile + " created successfully. Use harvest config add --key foo --value bar to change.")
			}
		}
		//viper.GetViper().WriteConfig()
	},
}

func init() {
	configCmd.AddCommand(createCmd)
}
