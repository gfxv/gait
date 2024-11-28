package cmds

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func NewHandler(commands ...*cli.Command) *cli.App {
	return &cli.App{
		EnableBashCompletion: true,

		Name:  "gait",
		Usage: "TODO: write usage part",
		Action: func(*cli.Context) error {
			fmt.Println("hmmm...")
			return nil
		},
		Commands: commands,
	}
}
