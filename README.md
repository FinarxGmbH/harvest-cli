## Harvest CLI



### How started and to learn
https://divrhino.com/articles/build-command-line-tool-go-cobra/

* snap install po
* edit .profile
* go install github.com/spf13/cobra-cli@latest
* go mod init github.com/FinarxGmbH/harvest-cli
* cobra-cli init --viper
* cobra-cli add config
* cobra-cli add create -p 'configCmd'
* go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
* oapi-codegen -package rest harvest-oapi.yaml >rest.go
