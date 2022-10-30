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
var rawtopicUpdateSchemaCmd = &cobra.Command{
	Use:   "update-schema",
	Short: "Load schema into harvester and print.",
	Long: `Load schema into from Database to the harvester node and 
print the schema as json used in syncher lib.`,
	Run: func(cmd *cobra.Command, args []string) {
		domainname, _ := cmd.Flags().GetString("domain")
		major, _ := cmd.Flags().GetString("major")

		fmt.Println("using config file " + viper.GetViper().ConfigFileUsed())
		fmt.Println("update-schema for raw topic " + domainname + "-V" + major)

		doMcpUpdateSchema(domainname, major)
	},
}

func init() {
	rawtopicCmd.AddCommand(rawtopicUpdateSchemaCmd)
}

type McpUpdateSchemaRequest struct {
	SystemCategory string `json:"systemCategory"`
	DomainName     string `json:"domainName"`
	MajorVersion   string `json:"majorVersion"`
}

type McpUpdateSchemaResponse struct {
	Topicname    string `json:"topicname"`
	Schema       string `json:"schema"`
	ErrorMessage string `json:"errorMsg"`
	ErrorCode    int    `json:"errorCode"`
}

func doMcpUpdateSchema(domainname string, major string) {
	//fmt.Println("Get random dad joke :P")
	url := mcpHost + "schema-update"

	request := McpUpdateSchemaRequest{
		SystemCategory: "core",
		DomainName:     domainname,
		MajorVersion:   major,
	}
	body, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseBytes := rest.PostMcpResponse(url, string(body), authToken)
	response := McpUpdateSchemaResponse{}

	if err := json.Unmarshal(responseBytes, &response); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return
	}

	fmt.Println(response.Topicname + ":" + response.Schema)
}
