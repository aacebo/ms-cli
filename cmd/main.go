package main

import (
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
	cmd := cli.NewCommand("ag").
		WithDescription("MS CLI (Unoffcial)").
		WithCommand(version.New())

	if adoUrl != "" {
		v, err := ado.New(adoUrl, adoToken)

		if err != nil {
			log.Fatal(err)
		}

		adoClient = v
	}

	cmd = cmd.WithCommand(
		_ado.New(adoClient).
			WithDisabled(adoClient == nil),
	)

	if err := cmd.Run(os.Args[1:]...); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
