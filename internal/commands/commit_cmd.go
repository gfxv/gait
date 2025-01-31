package cmds

import (
	"fmt"
	"gait/internal/model"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func CommitCmd(model *model.Model) *cli.Command {
	return &cli.Command{
		Name:    "commit",
		Aliases: []string{"c"},
		Usage:   "creates a commit with suggested message if acceptable",
		Action: func(cCtx *cli.Context) error {
			dir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current directory: %w", err)
			}

			if validateGitRepo(dir); err != nil {
				return fmt.Errorf("commit error: %w", err)
			}

			diff, err := exec.Command("git", "diff", "--staged").Output()
			if err != nil {
				return fmt.Errorf("failed to get diff, %w", err)
			}

			if len(diff) == 0 {
				return fmt.Errorf("no staged files found")
			}

			response, err := (*model).GetSuggestion(string(diff))
			if err != nil {
				return fmt.Errorf("failed to generate suggestion, %w", err)
			}

			fmt.Println(response)
			fmt.Println("Is this commit message acceptable? (y/n)")
			var input string
			if _, err = fmt.Scanf("%s", &input); err != nil {
				return fmt.Errorf("Failed to read input: ", err)
			}

			if strings.TrimSpace(strings.ToLower(input)) != "y" {
				return fmt.Errorf("Commit canceled.")
			}

			command := exec.Command("git", "commit", "-em", response)
			command.Stdin = os.Stdin
			command.Stdout = os.Stdout

			if err = command.Run(); err != nil {
				return fmt.Errorf("failed to start proc:", err)
			}

			return nil
		},
	}
}
