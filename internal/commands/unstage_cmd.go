package cmds

import (
	"fmt"
	"gait/internal/model"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func UnstageCmd(model *model.Model) *cli.Command {
	return &cli.Command{
		Name: "unstage",
		// Aliases: []string{""},
		Usage: "TODO: write description and usage",
		Action: func(cCtx *cli.Context) error {

			dir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current directory: %w", err)
			}

			if err := validateGitRepo(dir); err != nil {
				return fmt.Errorf("unstage error: %w", err)
			}

			if !cCtx.Args().Present() {
				return fmt.Errorf("no args provided")
			}

			files := strings.Join(cCtx.Args().Slice(), "")

			command := exec.Command("git", "reset", "--", files)
			command.Stdin = os.Stdin
			command.Stdout = os.Stdout

			if err = command.Run(); err != nil {
				return fmt.Errorf("failed to start proc: %w", err)
			}

			return nil
		},
	}
}
