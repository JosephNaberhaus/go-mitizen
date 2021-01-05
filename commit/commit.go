package commit

import (
	"github.com/JosephNaberhaus/go-mitizen/git"
)

func Commit() error {
	commit, err := showForm()
	if err != nil {
		return err
	}

	println()
	return git.Commit(commit.toCommitMessage())
}
