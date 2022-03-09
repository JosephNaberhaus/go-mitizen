package form

import (
	"github.com/JosephNaberhaus/prompt"
)

type Form struct {
	CommitType        Question
	Scope             Question
	Subject           Question
	Body              Question
	AreBreakingChange Question
	BreakingChanges   Question
	AreIssueReference Question
	IssueReferences   Question

	curPromptIndex int
}

func NewForm() *Form {
	return &Form{
		CommitType:        newCommitTypeQuestion(),
		Scope:             newScopeQuestion(),
		Subject:           newSubjectQuestion(),
		Body:              newBodyQuestion(),
		AreBreakingChange: newAreBreakingChangesQuestion(),
		BreakingChanges:   newBreakingChangeQuestion(),
		AreIssueReference: newAreIssueReferenceQuestion(),
		IssueReferences:   newIssueReferencesQuestion(),
	}
}

func (q *Form) IsFinished() bool {
	return q.curPromptIndex == len(q.AllPrompts())
}

func (q *Form) AllPrompts() []Question {
	questions := make([]Question, 0)
	questions = append(questions, q.CommitType)
	questions = append(questions, q.Scope)
	questions = append(questions, q.Subject)
	questions = append(questions, q.Body)
	questions = append(questions, q.AreBreakingChange)
	if q.AreBreakingChange.prompt.State() == prompt.Finished && q.AreBreakingChange.StringValue() == "true" {
		questions = append(questions, q.BreakingChanges)
	}
	questions = append(questions, q.AreIssueReference)
	if q.AreIssueReference.prompt.State() == prompt.Finished && q.AreIssueReference.StringValue() == "true" {
		questions = append(questions, q.IssueReferences)
	}
	return questions
}

func (q *Form) Previous() {
	q.curPromptIndex--
}

func (q *Form) Next() {
	q.curPromptIndex++
}

func (q *Form) CurPrompt() Question {
	allPrompts := q.AllPrompts()

	return allPrompts[q.curPromptIndex]
}
