package commit

import (
	"github.com/JosephNaberhaus/go-mitizen/git"
)

func Commit(dryRun bool) error {
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
