package workItems

import (
	"encoding/json"
	"fmt"

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
			b, _ := json.Marshal(workItem.Fields)
			fmt.Println(string(b))
		}

		fmt.Println(len(workItems))

		return nil
	}
}
