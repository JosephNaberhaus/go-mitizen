package prompt

import (
	"errors"
	"fmt"
	"log"
)

type SingleLine struct {
	base

	Name        string
	Description string

	MaxLength      int  // Maximum number of characters that can be submitted
	WrapLineLength int  // Number of columns to wrap the response by after it is submitted (doesn't effect input). Values <= represent no max length
	Required       bool // Whether a blank input is allowed
	ForceLowercase bool

	invalidMessage string // An invalid input message to display below the input text

	editor *editor
}

func (s *SingleLine) Show() error {
	s.base.Show()

	s.editor = &editor{
		NumCols:        s.output.numCols,
		ForceLowercase: s.ForceLowercase,
	}
	s.editor.Init()

	s.render()

	return loopUntilFinished(s)
}

func (s *SingleLine) handleInput(input Key) {
	if !s.showing || s.finished {
		return
	}

	s.invalidMessage = ""

	if input == ControlEnter {
		if s.Required && s.editor.empty() {
			s.invalidMessage = s.Name + " is required"
		} else if s.MaxLength > 0 && s.editor.numCharacters() > s.MaxLength {
			s.invalidMessage = fmt.Sprintf("%s length must be less than or equal to %d characters. Current length is %d characters.", s.Name, s.MaxLength, s.editor.numCharacters())
		} else {
			s.Finish()
			return
		}
	} else {
		s.editor.handleInput(input)
	}

	s.render()
}

func (s *SingleLine) render() {
	log.Printf("Rendering single-line prompt")
	s.output.clear()

	s.output.writeColor("? ", ColorGreen)
	s.output.write(s.Description)
	if s.MaxLength > 0 {
		s.output.write(fmt.Sprintf(" (max %d characters)", s.MaxLength))
	}
	s.output.write(": ")

	textColor := ColorWhite

	if !s.Required && s.editor.empty() {
		s.output.write("(press enter to skip) ")
	}

	if s.MaxLength > 0 {
		s.output.nextLine()

		numCharacters := s.editor.numCharacters()
		if numCharacters > s.MaxLength {
			textColor = ColorRed
			s.output.writeColor(fmt.Sprintf(" (%d) ", numCharacters), ColorRed)
		} else {
			s.output.writeColor(fmt.Sprintf(" (%d) ", numCharacters), ColorGreen)
		}
	}

	offsetX := s.output.cursorX
	offsetY := s.output.cursorY

	if s.finished {
		textColor = ColorCyan
	}

	if s.editor.numLines() > 1 {
		s.output.nextLine()
	}

	for _, line := range s.editor.lines {
		s.output.writeColorLn(line, textColor)
	}

	if len(s.invalidMessage) != 0 {
		s.output.writeColor(">> ", ColorRed)
		s.output.write(s.invalidMessage)
	}

	s.output.setCursor(s.editor.getRealCursorPosition(offsetX, offsetY))
	s.output.flush()
}

func (s *SingleLine) Finish() {
	s.base.Finish()

	if s.WrapLineLength > 0 {
		s.editor.wrapLines(s.WrapLineLength)
	}

	s.render()
	s.output.commit()
}

// Returns the response asserting that it will occupy only one line. If the response has been wrapped due to its length
// exceeding the WrapLineLength property then this method will panic
func (s *SingleLine) ResponseSingle() string {
	if s.editor.numLines() > 1 {
		panic(errors.New("multiple lines in editor when only one line was asserted"))
	}

	return s.editor.curLine()
}

func (s *SingleLine) Response() []string {
	if s.editor.empty() {
		return nil
	}

	return s.editor.lines
}
