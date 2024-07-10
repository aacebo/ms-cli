package workItems

import (
	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
)

func New(client *ado.Client) cli.Command {
	return cli.NewCommand("work-items").
		WithDescription("Azure Devops Work Items").
		WithCommand(
			cli.NewCommand("search").
				WithDescription("search project work items").
				WithParams(cli.Params{
					"project": cli.String().Required(),
				}).
				WithHandler(search(client)),
		)
}
