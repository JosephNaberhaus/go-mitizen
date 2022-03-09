package form

import (
	"github.com/JosephNaberhaus/go-mitizen/config"
	"github.com/JosephNaberhaus/prompt"
)

func newCommitTypeQuestion() Question {
	p := &prompt.Select{
		Question: "Select the type of change that you're committing",
		Options:  config.Types,
	}

	return Question{
		prompt: p,
	}
}

func newScopeQuestion() Question {
	p := &prompt.Text{
		Question:                 "What is the Scope of this change (e.g. component or file name)",
		IsSingleLine:             true,
		ShouldShowCharacterCount: true,
		ShouldForceLowercase:     config.ForceScopeLowerCase,
		OnKeyFunc:                goBackHandler,
	}

	return Question{
		prompt: p,
		beforeShow: func(form *Form) {
			maxLength := config.MaxHeaderLength

			// Remove the length of the commit type
			maxLength = maxLength - len(form.CommitType.StringValue()) - 2

			// Remove the length of the parentheses plus require at least one character for the Subject
			maxLength = maxLength - 3

			p.ValidatorFunc = maxLengthValidator("Scope", maxLength)
		},
	}
}

func newSubjectQuestion() Question {
	p := &prompt.Text{
		Question:                 "Write a short, imperative tense description of the change",
		IsSingleLine:             true,
		ShouldShowCharacterCount: true,
		ShouldForceLowercase:     config.ForceSubjectLowerCase,
		OnKeyFunc:                goBackHandler,
	}

	return Question{
		prompt: p,
		beforeShow: func(form *Form) {
			maxLength := config.MaxHeaderLength

			// Remove the length of the commit type
			maxLength = maxLength - len(form.CommitType.StringValue()) - 2

			// Remove the length of the Scope
			if form.Scope.StringValue() != "" {
				maxLength = maxLength - len(form.Scope.StringValue()) - 2
			}

			p.ValidatorFunc = combineValidators(
				requiredValidator("Subject"),
				maxLengthValidator("Subject", maxLength),
			)
		},
	}
}

func newBodyQuestion() Question {
	p := &prompt.Text{
		Question:              "Provide a longer description of the change: (press enter to skip)",
		OnSubmitMaxLineLength: config.MaxLineLength,
		OnKeyFunc:             goBackHandler,
	}

	return Question{
		prompt: p,
	}
}

func newAreBreakingChangesQuestion() Question {
	p := &prompt.Boolean{
		Question:  "Are there any breaking changes?",
		OnKeyFunc: goBackHandler,
	}

	return Question{
		prompt: p,
	}
}

func newBreakingChangeQuestion() Question {
	p := &prompt.Text{
		Question:              "Describe the breaking changes",
		IsSingleLine:          true,
		OnSubmitMaxLineLength: config.MaxLineLength,
		OnKeyFunc:             goBackHandler,
	}

	return Question{
		prompt: p,
	}
}

func newAreIssueReferenceQuestion() Question {
	p := &prompt.Boolean{
		Question:  "Does this change affect any open issues?",
		OnKeyFunc: goBackHandler,
	}

	return Question{
		prompt: p,
	}
}

func newIssueReferencesQuestion() Question {
	p := &prompt.Text{
		Question:              "Add issue references (e.g. \"fix #123\", \"re #123\".)",
		IsSingleLine:          true,
		OnSubmitMaxLineLength: config.MaxLineLength,
		OnKeyFunc:             goBackHandler,
	}

	return Question{
		prompt: p,
	}
}
