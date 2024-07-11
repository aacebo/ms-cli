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
					"project": cli.String().
						Description("project id/name").
						Required(),
				}).
				WithHandler(search(client)),
		).
		WithCommand(
			cli.NewCommand("sync-github").
				WithDescription("sync work items with github issues").
				WithParams(cli.Params{
					"project": cli.String().
						Description("project id/name").
						Required(),
					"gh-org": cli.String().
						Description("github organization name").
						Required(),
					"gh-repo": cli.String().
						Description("github repository name").
						Required(),
				}).
				WithHandler(syncGithub(client)),
		)
}
