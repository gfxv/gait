package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func NewHandler(commands ...*cli.Command) *cli.App {
	return &cli.App{
		EnableBashCompletion: true,

		Name:  "gait",
		Usage: "",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!!!")
			return nil
		},
		Commands: commands,
	}
}
