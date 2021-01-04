package commit

import (
	"github.com/JosephNaberhaus/go-mitizen/git"
)

func Commit() error {
	config := loadDefaultConfig()

	commit, err := showForm(config)
	if err != nil {
		return err
	}

	println()
	return git.Commit(commit.toCommitMessage())
}
