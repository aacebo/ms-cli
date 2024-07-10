package version

import "github.com/aacebo/ms-cli/cli"

func New() cli.Command {
	return cli.NewCommand("version").
		WithDescription("Get Package Version").
		WithHandler(version)
}
