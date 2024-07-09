package projects

import (
	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
)

func New(client *ado.Client) cli.Command {
	return cli.NewCommand("projects").
		WithDescription("Azure Devops Projects").
		WithCommand(
			cli.NewCommand("list").
				WithDescription("list all projects").
				WithHandler(list(client)),
		)
}
