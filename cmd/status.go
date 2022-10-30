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
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of mcp and nodes",
	Long:  `Show operations status of mcp, the cluster and list nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("using config file " + viper.GetViper().ConfigFileUsed())
		doMcpLs()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

type McpListResponse struct {
	Name   string    `json:"name"`
	Status int       `json:"status"`
	Nodes  []McpNode `json:"nodes"`
}

type McpNode struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
	BaseUrl string `json:"baseUrl"`
}

func doMcpLs() {
	//fmt.Println("Get random dad joke :P")
	url := mcpHost + "/list"
	responseBytes := rest.GetMcpResponse(url, authToken)
	response := McpListResponse{}

	if err := json.Unmarshal(responseBytes, &response); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	fmt.Println(response.Name + ":" + fmt.Sprint(response.Status))
	for i, node := range response.Nodes {
		fmt.Printf("%3d %s (%d) ", i, node.Name, node.Status)
	}
}
