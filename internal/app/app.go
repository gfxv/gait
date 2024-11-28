package app

import (
	"gait/internal/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type Model interface {
	GetSuggestion(string) (string, error)
}

type App struct {
	model Model
	cli   *cli.App
}

func NewApp(model Model) *App {
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
