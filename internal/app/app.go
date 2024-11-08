package app

import (
	"gait/internal/cmd"
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
		cli:   cmd.NewHandler(cmd.CommitCmd(&model)),
	}
}

// MustRun runs cli and panics if any error occurred
func (a *App) MustRun() {
	if err := a.cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
