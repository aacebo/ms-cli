package version

import "github.com/aacebo/ms-cli/cli"

func New() cli.Command {
	return cli.NewCommand("version").
		WithDescription("Get Package Version").
		WithParams(cli.Params{
			"a": cli.String().
				Description("a string param").
				Required(),
			"b": cli.String().
				Description("a string param with min length").
				MinLength(3),
		}).
		WithHandler(version)
}
