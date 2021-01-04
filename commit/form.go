package commit

import (
	"fmt"
	"github.com/JosephNaberhaus/go-mitizen/prompt"
	"github.com/eiannone/keyboard"
)

func showForm(config *Config) (commit *info, err error) {
	commit = new(info)

	err = keyboard.Open()
	if err != nil {
		return nil, fmt.Errorf("can't listen to keyboard: %w", err)
	}
	defer keyboard.Close()

	// Type
	commitTypePrompt := prompt.Select{
		Description: "Select the type of change that you're committing",
		Options:     config.Types,
	}
	err = commitTypePrompt.Show()
	if err != nil {
		return nil, fmt.Errorf("error while showing commit type prompt: %w", err)
	}
	commit.CommitType = commitTypePrompt.Response().Name

	// Scope
	scopePrompt := prompt.SingleLine{
		Name:          "scope",
		Description:   "What is the scope of this change (e.g. component or file name)",
		MaxCharacters: config.MaxHeaderLength - len(commit.CommitType) - 5,
		Required:      false,
	}
	err = scopePrompt.Show()
	if err != nil {
		return nil, fmt.Errorf("error while showing scope prompt: %w", err)
	}
	commit.Scope = scopePrompt.Response()

	// Subject
	var maxSubjectLength int
	if commit.Scope != "" {
		maxSubjectLength = config.MaxHeaderLength - len(commit.CommitType) - 4 - len(commit.Scope)
	} else {
		maxSubjectLength = config.MaxHeaderLength - len(commit.CommitType) - 2
	}

	subjectPrompt := prompt.SingleLine{
		Name:          "subject",
		Description:   "Write a short, imperative tense description of the change",
		MaxCharacters: maxSubjectLength,
		Required:      true,
	}
	err = subjectPrompt.Show()
	if err != nil {
		return nil, fmt.Errorf("error while showing description prompt: %w", err)
	}
	commit.subject = subjectPrompt.Response()

	// Body
	bodyPrompt := prompt.Multiline{
		Description: "Provide a longer description of the change: (press enter to skip)",
	}
	err = bodyPrompt.Show()
	if err != nil {
		return nil, fmt.Errorf("error while showing body prompt: %w", err)
	}
	commit.body = bodyPrompt.Response()

	// Breaking changes
	areBreakingChangesPrompt := prompt.YesNo{
		Description: "Are there any breaking changes?",
	}
	err = areBreakingChangesPrompt.Show()
	if err != nil {
		return nil, fmt.Errorf("error while showing breaking changes prompt: %w", err)
	}

	if areBreakingChangesPrompt.Response() {
		breakingChangesPrompt := prompt.SingleLine{
			Description: "Describe the breaking changes",
		}
		err = breakingChangesPrompt.Show()
		if err != nil {
			return nil, fmt.Errorf("error while showing breaking changes description prompt: %w", err)
		}
		commit.breakingChanges = breakingChangesPrompt.Response()
	}

	areIssueReferencesPrompt := prompt.YesNo{
		Description: "Does this change affect any open issues?",
	}
	err = areIssueReferencesPrompt.Show()
	if err != nil {
		return nil, fmt.Errorf("error while showing are issue references prompt: %w", err)
	}

	if areIssueReferencesPrompt.Response() {
		issueReferencesPrompt := prompt.SingleLine{
			Description:   "Add issue references (e.g. \"fix #123\", \"re #123\".)",
		}
		err = issueReferencesPrompt.Show()
		if err != nil {
			return nil, fmt.Errorf("error while showing issue reference prompt: %w", err)
		}
		commit.issueReference = issueReferencesPrompt.Response()
	}

	return commit, nil
}