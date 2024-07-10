package ado

import (
	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
	"github.com/aacebo/ms-cli/cmd/ado/projects"
	workItems "github.com/aacebo/ms-cli/cmd/ado/work_items"
)

func New(client *ado.Client) cli.Command {
	return cli.NewCommand("ado").
		WithDescription("Azure Devops").
		WithCommand(projects.New(client)).
		WithCommand(workItems.New(client))
}
