package prompt

import (
	"log"
)

type Multiline struct {
	base

	Description string

	editor *editor
}

func (m *Multiline) Show() error {
	m.base.Show()

	m.editor = newEditor(m.output.numCols)
	m.render()

	return loopUntilFinished(m)
}

func (m *Multiline) handleInput(input Key) {
	if !m.showing || m.finished {
		return
	}

	if input == ControlEnter {
		if m.editor.empty() {
			m.Finish()
			return
		}

		if m.editor.numLines() >= 2 && m.editor.lineFromLast(0) == "" && m.editor.lineFromLast(1) == "" {
			m.editor.removeLine(m.editor.numLines() - 1)
			m.editor.removeLine(m.editor.numLines() - 1)

			m.Finish()
			return
		}
	}

	m.editor.handleInput(input)
	m.render()
}

func (m *Multiline) render() {
	log.Printf("Rendering multiline prompt")
	m.output.clear()

	m.output.writeColor("? ", ColorGreen)
	m.output.write(m.Description)

	if m.finished {
		m.output.writeLn(":")
	} else if m.editor.empty() {
		m.output.writeLn(": (press enter to skip)")
	} else {
		m.output.writeLn(": (enter two empty lines to submit)")
	}

	offsetX := m.output.cursorX
	offsetY := m.output.cursorY

	textColor := ColorWhite
	if m.finished {
		textColor = ColorCyan
	}

	for i, line := range m.editor.lines {
		if i + 1 == len(m.editor.lines) {
			m.output.writeColor(line, textColor)
		} else {
			m.output.writeColorLn(line, textColor)
		}
	}

	m.output.setCursor(m.editor.getRealCursorPosition(offsetX, offsetY))
	m.output.flush()
}

func (m *Multiline) Finish() {
	m.base.Finish()

	m.render()
	m.output.commit()
}

func (m *Multiline) Response() []string {
	return m.editor.lines
}