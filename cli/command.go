package cli

import (
	"errors"
	"fmt"
	"slices"
)

type Command struct {
	Name        string
	Description string
	Params      Params
	Commands    []Command
	Handler     func(args Args) error
}

func NewCommand(name string) Command {
	return Command{
		Name:     name,
		Params:   Params{},
		Commands: []Command{},
	}
}

func (self Command) WithDescription(description string) Command {
	self.Description = description
	return self
}

func (self Command) WithParams(params Params) Command {
	self.Params = params
	return self
}

func (self Command) WithCommand(cmd Command) Command {
	self.Commands = append(self.Commands, cmd)
	return self
}

func (self Command) WithHandler(handler func(args Args) error) Command {
	self.Handler = handler
	return self
}

func (self Command) Run(args ...string) error {
	if len(args) > 0 {
		if args[0] == "help" {
			self.help()
			return nil
		}

		i := slices.IndexFunc(self.Commands, func(cmd Command) bool {
			return cmd.Name == args[0]
		})

		if i > -1 {
			cmd := self.Commands[i]
			return cmd.Run(args[1:]...)
		}
	}

	errs := []error{}
	parsedArgs, err := self.Params.Validate(args...)

	if err != nil {
		errs = append(errs, err)
		return errors.Join(errs...)
	}

	if self.Handler != nil {
		return self.Handler(parsedArgs)
	}

	return errors.Join(errs...)
}

func (self Command) help() {
	if self.Description != "" {
		fmt.Println(self.Description)
	}

	if len(self.Params) > 0 {
		fmt.Println("\nParameters:")

		for key, param := range self.Params {
			fmt.Printf(
				"\t%s: %s\n",
				key,
				param.GetDescription(),
			)
		}
	}

	if len(self.Commands) > 0 {
		fmt.Println("\nSub Commands:")

		for _, cmd := range self.Commands {
			fmt.Printf("- %s: %s\n", cmd.Name, cmd.Description)
		}
	}

	return
}
