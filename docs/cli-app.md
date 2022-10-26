# Create Cli App in Golang

## installation

```shell
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
```

## Use Cli to generator code

```shell
cobra-cli  -h
```

```shell

> cobra-cli  -h
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
Usage:
cobra-cli [command]

Available Commands:
add         Add a command to a Cobra Application
completion  Generate the autocompletion script for the specified shell
help        Help about any command
init        Initialize a Cobra Application

Flags:
-a, --author string    author name for copyright attribution (default "YOUR NAME")
--config string    config file (default is $HOME/.cobra.yaml)
-h, --help             help for cobra-cli
-l, --license string   name of license for the project
--viper            use Viper for configuration

Use "cobra-cli [command] --help" for more information about a command.
```