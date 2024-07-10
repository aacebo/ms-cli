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
	var adoClient *ado.Client
	adoUrl := os.Getenv("AZURE_DEVOPS_URL")
	adoToken := os.Getenv("AZURE_DEVOPS_TOKEN")

	if adoUrl != "" {
		v, err := ado.New(adoUrl, adoToken)

		if err != nil {
			log.Fatal(err)
		}

		adoClient = v
	}

	cmd := cli.NewCommand("ag").
		WithDescription("MS CLI (Unoffcial)").
		WithCommand(version.New()).
		WithCommand(
			_ado.New(adoClient).
				WithDisabled(adoClient == nil),
		)

	if err := cmd.Run(os.Args[1:]...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
