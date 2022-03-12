package commit

import (
	"github.com/JosephNaberhaus/go-mitizen/config"
	"github.com/JosephNaberhaus/go-mitizen/util"
	"regexp"
	"strings"
)

type info struct {
	CommitType      string
	Scope           string
	subject         string
	body            string
	breakingChanges string
	issueReference  string
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

	if i.body != "" {
		messageBuilder.WriteString("\n")

		body := i.body
		if !config.AllowBlankLinesInBody {
			body = regexp.MustCompile("\\n\\n+").ReplaceAllString(body, "\n")
		}

		messageBuilder.WriteString(body)
	}

	if i.breakingChanges != "" {
		messageBuilder.WriteString("\n\n")
		messageBuilder.WriteString("BREAKING CHANGE: ")

		// Because of the prefix, we must re-wrap the lines so that the max line length isn't exceeded
		reWrappedLines := make([]string, 0)

		firstLineLength := util.Min(config.MaxLineLength-len("BREAKING CHANGE: "), len(i.breakingChanges))
		reWrappedLines = append(reWrappedLines, i.breakingChanges[:firstLineLength])
		reWrappedLines = append(reWrappedLines, util.WrapString(i.breakingChanges[firstLineLength:], config.MaxLineLength)...)

		messageBuilder.WriteString(strings.Join(reWrappedLines, "\n"))
	}

	if i.issueReference != "" {
		messageBuilder.WriteString("\n\n")
		messageBuilder.WriteString(i.issueReference)
	}

	return messageBuilder.String()
}
