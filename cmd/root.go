/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile   string
	mcpHost   string
	authToken string

	// see https://github.com/spf13/cobra-cli/blob/main/README.md for adding new commands

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "harvest",
		Short: "Operate the and develop against the OCP harvester cluster",
		Long: `The harvester cluster runs in a k8s cluster and consists of harvester
nodes, trigger-exporter and the harvesters Master Control Program (MCP).

Basic Commands:
  help:    Prints this output, meant to help. 
           Also see https://github.com/FinarxGmbH/harvest-cli
  config:  Configure -m default and credentials to reduce number of cmdline arguments required.
           Stores configuration in ~/.harvest/cli.yml

Operations Commands:	
  status:  Prints status only of mcp and nodes (harvester, trigger-exporter etc.)
  ls:      Prints cluster or node infos depending on whether -m or -t is given.


Developer Commands:
  jpagen:  Generates JPA Code useful to store entities in Spring and JEE/Quarkus apps.
  pggen:   Generates PostgreSQL code to create entity tables natively.
  h2gen:   Generates H2 Database code to create entity tables natively.

USAGE harvest [-t <targetHarvesterHost>] [-m <masterControlProgramHost>]  <command> :
  -v                             Verbose output.
  -t <targetHarvesterHost>       Directly adress or filter one harvester node
  -m <masterControlProgramHost>  Override the MCP host given in ~/.harvest/cli.yaml

	.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.harvest/cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("verbose", "v", false, "Generate verbose output what is going on")

	rootCmd.PersistentFlags().StringVarP(&mcpHost, "mcpHost", "m", "http://localhost:7878/", "Host adress for Master Control Program of the cluster")
	rootCmd.PersistentFlags().StringVarP(&authToken, "authToken", "a", "somesecrettoken", "Secret token to access Master Control Program of the cluster")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	//rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("mcpHost", rootCmd.PersistentFlags().Lookup("mcpHost"))
	viper.BindPFlag("authToken", rootCmd.PersistentFlags().Lookup("authToken"))
	//viper.BindPFlag("viper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("authToken", "somesecrettoken")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".harvest/cli" (without extension).
		viper.AddConfigPath(home + "/.harvest")
		viper.SetConfigType("yaml")
		viper.SetConfigName("cli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
