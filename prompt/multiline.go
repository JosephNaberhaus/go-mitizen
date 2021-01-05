package prompt

import (
	"log"
)

type Multiline struct {
	base

	Description     string

	AllowBlankLines bool // Whether to remove all blank lines from the input after it is submitted (doesn't effect input)
	WrapLineLength  int  // Number of columns to wrap the response by after it is submitted (doesn't effect input). Values <= represent no max length

	editor *editor
}

func (m *Multiline) Show() error {
	m.base.Show()

	m.editor = &editor{NumCols: m.output.numCols}
	m.editor.Init()

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

		if m.editor.isLastLine(m.editor.cursorY) && m.editor.numLines() >= 2 && m.editor.lineFromLast(0) == "" && m.editor.lineFromLast(1) == "" {
			numBlankLinesRemoved := 0
			for i := m.editor.numLines() - 1; i >= 0; i-- {
				if m.editor.lines[i] == "" {
					m.editor.removeLine(i)
					numBlankLinesRemoved++

					if m.AllowBlankLines && numBlankLinesRemoved == 2 {
						break
					}
				}
			}

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

	if m.WrapLineLength > 0 {
		m.editor.wrapLines(m.WrapLineLength)
	}

	m.render()
	m.output.commit()
}

func (m *Multiline) Response() []string {
	if m.editor.empty() {
		return nil
	}

	return m.editor.lines
}