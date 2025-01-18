package cmds

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func NewHandler(commands ...*cli.Command) *cli.App {
	return &cli.App{
		EnableBashCompletion: true,

		Name:  "gait",
		Usage: "An AI powered tool for generating commit messages",
		Action: func(*cli.Context) error {
			// TODO: print proper help (probably add some info like description, version, etc.)
			fmt.Println("Type 'gait help' to see all available commands")
			return nil
		},
		Commands: commands,
	}
}
