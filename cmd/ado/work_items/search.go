package workItems

import (
	"fmt"

	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
)

func search(client *ado.Client) func(cli.Args) error {
	return func(args cli.Args) error {
		workItems, err := client.WorkItems(args["project"].String()).Search("AI")

		if err != nil {
			return err
		}

		fmt.Println(len(workItems))
		return nil
	}
}
