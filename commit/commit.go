package commit

import (
	"errors"
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

	err := overrideConfig()
	if err != nil {
		return err
	}

	commit, err := showForm()
	if err != nil {
		return err
	}

	if dryRun {
		println("Commit Message:")
		println(commit.toCommitMessage())
		return nil
	}

	println()
	return git.Commit(commit.toCommitMessage())
}
