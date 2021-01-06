package util

import (
	"strings"
)

func WrapString(s string, maxLineSize int) []string {
	words := strings.Split(s, " ")

	lines := make([]string, 0)
	curLine := strings.Builder{}

	addLine := func() {
		lines = append(lines, curLine.String())
		curLine.Reset()
	}

	for i := 0; i < len(words); i++ {
		if len(words[i]) > maxLineSize {
			characterToAdd := maxLineSize - curLine.Len()
			curLine.WriteString(words[i][:characterToAdd])
			addLine()

			words[i] = words[i][characterToAdd:]
			i--
		} else {
			if (len(words[i]) + 1) > (maxLineSize - curLine.Len()) {
				addLine()
			}

			curLine.WriteString(words[i])
			if i != len(words)-1 {
				curLine.WriteRune(' ')
			}
		}
	}

	if curLine.Len() != 0 {
		addLine()
	}

	return lines
}
