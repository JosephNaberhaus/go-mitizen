package prompt

import (
	"fmt"
	"log"
)

type SingleLine struct {
	base

	Name string
	Description string

	MaxCharacters int
	Required bool

	invalidMessage string // An invalid input message to display below the input text

	editor *editor
}

func (s *SingleLine) Show() error {
	s.base.Show()

	s.editor = newEditor(s.output.numCols)
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
		} else if s.MaxCharacters > 0 && s.editor.numCharacters() > s.MaxCharacters {
			s.invalidMessage = fmt.Sprintf("%s length must be less than or equal to %d characters. Current length is %d characters.", s.Name, s.MaxCharacters, s.editor.numCharacters())
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
	if s.MaxCharacters > 0 {
		s.output.write(fmt.Sprintf(" (max %d characters)", s.MaxCharacters))
	}
	s.output.write(": ")

	textColor := ColorWhite

	if !s.Required {
		s.output.write("(press enter to skip) ")
	}

	if s.MaxCharacters > 0 {
		s.output.nextLine()

		numCharacters := s.editor.numCharacters()
		if numCharacters > s.MaxCharacters {
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

	s.output.writeColorLn(s.editor.curLine(), textColor)

	if len(s.invalidMessage) != 0 {
		s.output.writeColor(">> ", ColorRed)
		s.output.write(s.invalidMessage)
	}

	s.output.setCursor(s.editor.getRealCursorPosition(offsetX, offsetY))
	s.output.flush()
}

func (s *SingleLine) Finish() {
	s.base.Finish()

	s.render()
	s.output.commit()
}

func (s *SingleLine) Response() string {
	return s.editor.curLine()
}
