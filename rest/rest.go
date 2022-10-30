package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetMcpResponse(baseAPI string, authToken string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not contact MCP. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Harvest CLI (https://github.com/FinarxGmbH/harvest-cli)")
	request.Header.Add("Authorization", "Bearer "+authToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

func PostMcpResponse(baseAPI string, body string, authToken string) []byte {
	request, err := http.NewRequest(
		http.MethodPost,         //method
		baseAPI,                 //url
		strings.NewReader(body), //body
	)

	if err != nil {
		log.Printf("Could not contact MCP. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Harvest CLI (https://github.com/FinarxGmbH/harvest-cli)")
	request.Header.Add("Authorization", "Bearer "+authToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}
	fmt.Println("response:" + string(responseBytes))
	return responseBytes
}
