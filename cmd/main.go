package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
	_ado "github.com/aacebo/ms-cli/cmd/ado"
	"github.com/aacebo/ms-cli/cmd/version"
)

func main() {
	adoClient, err := ado.New(
		os.Getenv("AZURE_DEVOPS_URL"),
		os.Getenv("AZURE_DEVOPS_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = cli.NewCommand("ag").
		WithDescription("MS CLI (Unoffcial)").
		WithCommand(version.New()).
		WithCommand(_ado.New(adoClient)).
		Run(os.Args[1:]...)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
