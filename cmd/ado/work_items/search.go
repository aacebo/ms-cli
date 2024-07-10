package workItems

import (
	"fmt"
	"strings"

	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
)

func search(client *ado.Client) func(cli.Args) error {
	return func(args cli.Args) error {
		workItems, err := client.WorkItems(args["project"].String()).Search("Plat_Ailib")

		if err != nil {
			return err
		}

		for _, workItem := range workItems {
			id := (*workItem.Fields)["system.id"]
			title := (*workItem.Fields)["system.title"]
			t := (*workItem.Fields)["system.workitemtype"]
			state := (*workItem.Fields)["system.state"]
			tags := strings.Split((*workItem.Fields)["system.tags"], "; ")

			fmt.Printf(
				"- [%s (%s)]:\n\t- Type => %s\n\t- State => %s\n\t- Tags => %s\n\n",
				title,
				id,
				t,
				state,
				strings.Join(tags, ", "),
			)
		}

		return nil
	}
}
