package commit

import (
	"fmt"
	"github.com/JosephNaberhaus/go-mitizen/form"
	"github.com/eiannone/keyboard"
	"log"
)

func showForm() (commit *info, err error) {
	commit = new(info)

	err = keyboard.Open()
	if err != nil {
		return nil, fmt.Errorf("can't listen to keyboard: %w", err)
	}
	defer keyboard.Close()

	f := form.NewForm()
	for !f.IsFinished() {
		curPrompt := f.CurPrompt()

		result, err := curPrompt.Show(f)
		if err != nil {
			return nil, err
		}

		log.Printf("Result is %v", result)

		if result == form.ResultSubmit {
			f.Next()
		} else {
			f.Previous()
		}
	}

	return &info{
		CommitType:      f.CommitType.StringValue(),
		Scope:           f.Scope.StringValue(),
		subject:         f.Subject.StringValue(),
		body:            f.Body.StringValue(),
		breakingChanges: f.BreakingChanges.StringValue(),
		issueReference:  f.IssueReferences.StringValue(),
	}, nil
}
