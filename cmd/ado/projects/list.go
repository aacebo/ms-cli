package projects

import (
	"fmt"

	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
)

func list(client *ado.Client) func(...string) error {
	return func(args ...string) error {
		projects, err := client.Projects().List()

		if err != nil {
			return err
		}

		for _, project := range projects {
			fmt.Printf(
				"[%s (%s)] => %s\n",
				cli.Text().RedForeground(*project.Name),
				project.Id.String(),
				*project.Url,
			)
		}

		return nil
	}
}
