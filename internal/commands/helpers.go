package cmds

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func validateGitRepo(dir string) error {

	repo, err := git.PlainOpen(dir)
	if err != nil {
		return fmt.Errorf("failed to open repository, %w", err)
	}

	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree, %w", err)
	}

	status, err := wt.Status()
	if err != nil {
		return fmt.Errorf("failed to get git status, %w", err)
	}

	if status.IsClean() {
		return fmt.Errorf("git worktree is clean")
	}

	return nil
}
