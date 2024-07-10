package version

import (
	"fmt"

	"github.com/aacebo/ms-cli/cli"
)

func version(args cli.Args) error {
	fmt.Println(args)
	fmt.Println("0.0.0")
	return nil
}
