## Harvest CLI



### How started and to learn
https://divrhino.com/articles/build-command-line-tool-go-cobra/

* sudo apt update
* sudo apt-get install build-essential
* snap install go
* edit .profile
* go install -v golang.org/x/tools/gopls@latest
* go install github.com/spf13/cobra-cli@latest
* go mod init github.com/FinarxGmbH/harvest-cli
* cobra-cli init --viper
* cobra-cli add config
* cobra-cli add create -p 'configCmd'
* go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
* oapi-codegen -package rest harvest-oapi.yaml >rest.go
