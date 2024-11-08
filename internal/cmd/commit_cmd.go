package cmd

import (
	"fmt"
	"gait/internal"
	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func CommitCmd(model *internal.Model) *cli.Command {
	return &cli.Command{
		Name:    "commit",
		Aliases: []string{"c"},
		Usage:   "creates a commit suggested message if acceptable",
		Action: func(cCtx *cli.Context) error {

			dir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current directory: %w", err)
			}

			repo, err := git.PlainOpen(dir)
			if err != nil {
				return fmt.Errorf("failed to open repository: %w", err)
			}

			wt, err := repo.Worktree()
			if err != nil {
				return fmt.Errorf("failed to get worktree: %w", err)
			}

			status, err := wt.Status()
			if err != nil {
				return fmt.Errorf("failed to get git status, %w", err)
			}

			if status.IsClean() {
				return fmt.Errorf("git worktree is clean")
			}

			diff, err := exec.Command("git", "diff", "--staged").Output()
			if err != nil {
				return fmt.Errorf("failed to get diff, %w", err)
			}

			response, err := (*model).GetSuggestion(string(diff))
			if err != nil {
				return fmt.Errorf("failed to generate suggestion, %w", err)
			}

			fmt.Println(response)
			return nil
		},
	}
}
