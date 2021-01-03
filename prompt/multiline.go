package prompt

import (
	"log"
)

type Multiline struct {
	Description string

	editor *editor
	output *output

	showing bool
	finished bool
}

func (m *Multiline) Show() {
	if m.showing {
		panic("cannot call show multiple times")
	}

	m.output = newOutput()
	m.editor = newEditor(m.output.numCols)
	m.showing = true

	m.render()
}

func (m *Multiline) HandleInput(input Key) {
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

	for i, line := range m.editor.lines {
		if i + 1 == len(m.editor.lines) {
			m.output.write(line)
		} else {
			m.output.writeLn(line)
		}
	}

	m.output.setCursor(m.editor.getRealCursorPosition(offsetX, offsetY))
	m.output.flush()
}

func (m *Multiline) Showing() bool {
	return m.showing
}

func (m *Multiline) Finish() {
	m.finished = true
	m.render()
	m.output.commit()
}

func (m *Multiline) Finished() bool {
	return m.finished
}