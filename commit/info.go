package commit

import "strings"

type info struct {
	CommitType string
	Scope string
	subject string
	body []string
	breakingChanges string
	issueReference string
}

func (i *info) toCommitMessage() string {
	messageBuilder := strings.Builder{}

	messageBuilder.WriteString(i.CommitType)

	if i.Scope != "" {
		messageBuilder.WriteRune('(')
		messageBuilder.WriteString(i.Scope)
		messageBuilder.WriteRune(')')
	}

	messageBuilder.WriteString(": ")
	messageBuilder.WriteString(i.subject)

	if i.body != nil {
		messageBuilder.WriteString("\n\n")
		messageBuilder.WriteString(strings.Join(i.body, "\n"))
	}

	if i.breakingChanges != "" {
		messageBuilder.WriteString("\n\n")
		messageBuilder.WriteString("BREAKING CHANGE: ")
		messageBuilder.WriteString(i.breakingChanges)
	}

	if i.issueReference != "" {
		messageBuilder.WriteString("\n\n")
		messageBuilder.WriteString(i.issueReference)
	}

	return messageBuilder.String()
}