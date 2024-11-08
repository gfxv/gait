package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Name: "gait",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!!!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
