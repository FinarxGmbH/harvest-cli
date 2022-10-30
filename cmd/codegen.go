/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"

	rest "github.com/FinarxGmbH/harvest-cli/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var codegenCmd = &cobra.Command{
	Use:   "codegen",
	Short: "Generate source code for developer's use.",
	Long:  `Generate source code for developer's use.`,
	Run: func(cmd *cobra.Command, args []string) {

		domainname, _ := cmd.Flags().GetString("domain")
		major, _ := cmd.Flags().GetString("major")
		kind, _ := cmd.Flags().GetString("kind")

		fmt.Println("using config file " + viper.GetViper().ConfigFileUsed())
		fmt.Println("codegen kind " + kind + " for domain " + domainname + "-V" + major)

		doMcpCodeGen(domainname, major, kind)
	},
}

func init() {
	rootCmd.AddCommand(codegenCmd)

	codegenCmd.PersistentFlags().StringP("domain", "d", "", "The domainName for rawtopic subcommands.")
	codegenCmd.PersistentFlags().StringP("major", "j", "", "The majorVersion for the rawtopic subcommands.")
	codegenCmd.PersistentFlags().StringP("kind", "k", "", "The kind of source code to generate (e.g. jpa).")

}

type McpCodeGenRequest struct {
	Subfolder    string `json:"subfolder"`
	Packagename  string `json:"packagename"`
	Kind         string `json:"kind"`
	DomainName   string `json:"domainName"`
	MajorVersion string `json:"majorVersion"`
}

type McpCodeGenResponse struct {
	Subfolder    string `json:"subfolder"`
	Filename     string `json:"filename"`
	ErrorMessage string `json:"errorMsg"`
	ErrorCode    int    `json:"errorCode"`
	Source       string `json:"source"`
}

func doMcpCodeGen(domainname string, major string, kind string) {
	//fmt.Println("Get random dad joke :P")
	url := mcpHost + "/codegen"

	request := McpCodeGenRequest{
		Kind:         kind,
		DomainName:   domainname,
		MajorVersion: major,
	}
	body, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseBytes := rest.PostMcpResponse(url, string(body), authToken)
	response := McpCodeGenResponse{}

	if err := json.Unmarshal(responseBytes, &response); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	fmt.Println(response.Subfolder + ":" + string(response.Filename))
	fmt.Println(response.Source)
}
