package commit

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/go-mitizen/git"
)

func Commit(dryRun bool) error {
	if !git.IsInGitRepository() {
		println("Error: not in Git repository")
		return errors.New("not in git repository")
	}

	if !git.AreStagedFiles() {
		println("Error: no files are staged to commit")
		return errors.New("no staged files")
	}

	commit, err := showForm()
	if err != nil {
		return err
	}

	if dryRun {
		fmt.Println("Commit Message:")
		fmt.Println(commit.toCommitMessage())
		return nil
	}

	fmt.Println()
	return git.Commit(commit.toCommitMessage())
}
