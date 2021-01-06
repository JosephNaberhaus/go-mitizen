package commit

import (
	"github.com/JosephNaberhaus/go-mitizen/util"
	"strings"
)

type info struct {
	CommitType      string
	Scope           string
	subject         string
	body            []string
	breakingChanges []string
	issueReference  []string
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

	if i.breakingChanges != nil {
		messageBuilder.WriteString("\n\n")
		messageBuilder.WriteString("BREAKING CHANGE: ")

		// Because of the prefix, we must re-wrap the lines so that the max line length isn't exceeded
		reWrappedLines := make([]string, 0)

		joined := strings.Join(i.breakingChanges, "")
		firstLineLength := util.Min(config.MaxLineLength-len("BREAKING CHANGE: "), len(joined))
		reWrappedLines = append(reWrappedLines, joined[:firstLineLength])
		reWrappedLines = append(reWrappedLines, util.SplitStringIntoChunks(joined[firstLineLength:], config.MaxLineLength)...)

		messageBuilder.WriteString(strings.Join(reWrappedLines, "\n"))
	}

	if i.issueReference != nil {
		messageBuilder.WriteString("\n\n")
		messageBuilder.WriteString(strings.Join(i.issueReference, "\n"))
	}

	return messageBuilder.String()
}
