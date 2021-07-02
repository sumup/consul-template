package main // import "github.com/sumup/consul-template"

import (
	"github.com/sumup/consul-template/cli"
	"os"
)

func main() {
	cliInstance := cli.NewCLI(os.Stdout, os.Stderr)
	os.Exit(cliInstance.Run(os.Args))
}
