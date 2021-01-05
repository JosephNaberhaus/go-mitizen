package prompt

import (
	"fmt"
	"github.com/JosephNaberhaus/go-mitizen/util"
	"log"
	"strings"
)

const maxLines = 7

type SelectionOption struct {
	Name string
	Description string
}

type Select struct {
	base

	Description string
	Options []*SelectionOption

	optionLines []string
	optionIndexToLine []int

	cursor int
}

func (s *Select) Show() error {
	s.base.Show()

	s.computeOptionLines()

	s.output.hideCursor()
	s.render()

	return loopUntilFinished(s)
}

func (s *Select) handleInput(input Key) {
	if input == ControlUp {
		s.cursor--
		if s.cursor < 0 {
			s.cursor = len(s.Options) - 1
		}
	} else if input == ControlDown {
		s.cursor++
		if s.cursor == len(s.Options) {
			s.cursor = 0
		}
	} else if input == ControlEnter {
		s.Finish()
		return
	}

	s.render()
}

func (s *Select) render() {
	log.Printf("Rendering select prompt")
	s.output.clear()

	s.output.writeColor("? ", ColorGreen)
	s.output.write(s.Description)
	s.output.write(": ")
	if s.finished {
		s.output.writeColor(fmt.Sprintf(" %s: %s", s.selected().Name, s.selected().Description), ColorCyan)
		return
	} else {
		s.output.writeColor("(Use arrow keys)", ColorGreen)
	}
	s.output.nextLine()

	startLine := s.optionIndexToLine[s.cursor] - (util.Min(maxLines - 1, len(s.optionLines)) / 2)
	endLine := startLine + util.Min(maxLines - 1, len(s.optionLines))
	for line := startLine; line <= endLine; line++ {
		wrappedLine := s.wrapLine(line)

		option := s.optionLines[wrappedLine]

		if wrappedLine == s.optionIndexToLine[s.cursor] {
			s.output.writeColor("> ", ColorCyan)
		} else {
			s.output.write("  ")
		}

		if s.lineIsSelected(wrappedLine) {
			s.output.writeColorLn(option, ColorCyan)
		} else {
			s.output.writeLn(option)
		}
	}

	if len(s.optionLines) > maxLines {
		s.output.writeColor("(Move up and down to reveal more choices)", ColorGreen)
	}

	s.output.flush()
}

func (s *Select) Finish() {
	s.base.Finish()

	s.output.showCursor()
	s.render()
	s.output.commit()
}

func (s *Select) wrapLine(line int) int {
	return (line % len(s.optionLines) + len(s.optionLines)) % len(s.optionLines)
}

func (s *Select) selected() *SelectionOption {
	return s.Options[s.cursor]
}

func (s *Select) lineIsSelected(line int) bool {
	if s.cursor + 1 == len(s.Options) {
		return s.optionIndexToLine[s.cursor] <= line
	}

	return s.optionIndexToLine[s.cursor] <= line && line < s.optionIndexToLine[s.cursor + 1]
}

func (s *Select) computeOptionLines() {
	s.optionLines = make([]string, 0)

	longestName := s.longestName()
	longestNamePadding := strings.Repeat(" ", longestName + 1)
	for _, option := range s.Options {
		optionString := fmt.Sprintf("%s:%s %s", option.Name, strings.Repeat(" ", longestName - len(option.Name)), option.Description)

		s.optionIndexToLine = append(s.optionIndexToLine, len(s.optionLines))
		s.optionLines = append(s.optionLines, optionString[:util.Min(len(optionString), s.output.numCols - 2)])

		for i := s.output.numCols - 2; i < len(optionString); i += s.output.numCols - 2 - longestName {
			line := optionString[i:util.Min(len(optionString), i + s.output.numCols)]

			s.optionLines = append(s.optionLines, longestNamePadding + line)
		}
	}
}

func (s *Select) longestName() int {
	longestName := 0

	for _, option := range s.Options {
		if len(option.Name) > longestName {
			longestName = len(option.Name)
		}
	}

	return longestName
}

func (s *Select) Response() *SelectionOption {
	return s.selected()
}