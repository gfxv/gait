package cmds

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func NewHandler(commands ...*cli.Command) *cli.App {
	return &cli.App{
		EnableBashCompletion: true,

		Name:  "gait",
		Usage: "Use 'gait c' or 'gait commit' to generate commit messsage based on staged files",
		Action: func(*cli.Context) error {
			// TODO: print proper help
			fmt.Println("Use 'gait c' or 'gait commit' to generate commit messsage based on staged files")
			return nil
		},
		Commands: commands,
	}
}
