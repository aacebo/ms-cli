package projects

import (
	"fmt"

	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
)

func list(client *ado.Client) func(cli.Args) error {
	return func(args cli.Args) error {
		projects, err := client.Projects().List()

		if err != nil {
			return err
		}

		for _, project := range projects {
			fmt.Printf(
				cli.Text("[%s (%s)]").CyanForeground().String()+" => %s\n",
				cli.Text(*project.Name).Bold(),
				project.Id.String(),
				cli.Text(*project.Url).Underline(),
			)
		}

		return nil
	}
}
