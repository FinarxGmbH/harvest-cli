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
var rawtopicIntialLoadCmd = &cobra.Command{
	Use:   "initial-load",
	Short: "Initial load a raw topic.",
	Long:  `Load a raw topic initially with all data from A-view.`,
	Run: func(cmd *cobra.Command, args []string) {
		domainname, _ := cmd.Flags().GetString("domain")
		major, _ := cmd.Flags().GetString("major")

		fmt.Println("using config file " + viper.GetViper().ConfigFileUsed())
		fmt.Println("update-schema for raw topic " + domainname + "-V" + major)

		doMcpInitialLoad(domainname, major)
	},
}

func init() {
	rawtopicCmd.AddCommand(rawtopicIntialLoadCmd)
}

type McpInitialLoadRequest struct {
	SystemCategory string `json:"systemCategory"`
	DomainName     int    `json:"domainName"`
	MajorVersion   string `json:"majorVersion"`
}

type McpInitialLoadResponse struct {
	Topicname     string `json:"topicname"`
	NumEntities   int    `json:"numEntities"`
	SecondsToLoad int    `json:"secondsToLoad"`
	ErrorMessage  string `json:"errorMsg"`
	ErrorCode     int    `json:"errorCode"`
}

func doMcpInitialLoad(domainname string, major string) {
	//fmt.Println("Get random dad joke :P")
	url := mcpHost + "initial-load"
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
	response := McpInitialLoadResponse{}

	if err := json.Unmarshal(responseBytes, &response); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return
	}

	fmt.Println(response.Topicname + ": loaded " + fmt.Sprint(response.NumEntities) + " in " + fmt.Sprint(response.SecondsToLoad) + " seconds")
}
