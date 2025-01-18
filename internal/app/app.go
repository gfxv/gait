package app

import (
	"gait/internal/commands"
	"gait/internal/model"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type App struct {
	model model.Model
	cli   *cli.App
}

func NewApp(model model.Model) *App {
	return &App{
		model: model,
		cli: cmds.NewHandler(
			cmds.CommitCmd(&model),
			cmds.UnstageCmd(&model),
		),
	}
}

// MustRun runs cli and panics if any error occurred
func (a *App) MustRun() {
	if err := a.cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
