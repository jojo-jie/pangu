# Server docker initialization command

## Compile
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o init.sh main.go
```
## Install Kratos
```
Usage:
   [command]

Available Commands:
  clean       Uninstall old Docker versions
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  install     Docker installation and environment initialization

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```
## Pre-Installed Software
1. docker 
2. docker-compose
3. portainer

