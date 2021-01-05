package prompt

import (
	"github.com/JosephNaberhaus/go-mitizen/util"
	"log"
	"unicode"
)

// Implements a standard editor.
type editor struct {
	ForceLowercase bool
	NumCols        int

	lines []string

	cursorX int
	cursorY int

	preferredX int
}

func (e *editor) Init() {
	e.lines = []string{""}
}

func (e *editor) left() {
	e.cursorX--

	if e.cursorX < 0 {
		if e.cursorY > 0 {
			e.cursorY--
			e.cursorX = e.curLineLength()
		} else {
			e.cursorX = 0
		}
	}

	e.preferredX = e.cursorX
}

func (e *editor) right() {
	e.cursorX++

	if e.cursorX > e.curLineLength() {
		if e.cursorY + 1 < e.numLines() {
			e.cursorY++
			e.cursorX = 0
		} else {
			e.cursorX = e.curLineLength()
		}
	}

	e.preferredX = e.cursorX
}

func (e *editor) up() {
	if e.cursorX >= e.NumCols {
		e.cursorX -= e.NumCols
	} else if e.cursorY != 0 {
		e.cursorY--

		e.cursorX = e.NumCols* (e.curLineLength() / e.NumCols) + util.Min(e.curLineLength(), e.preferredX)
	}
}

func (e *editor) down() {
	if e.curLineLength() - e.cursorX >= e.NumCols {
		e.cursorX += e.NumCols
	} else if !e.isLastLine(e.cursorY) {
		e.cursorY++
		e.cursorX = util.Min(e.curLineLength(), e.preferredX)
	}
}

func (e *editor) newline() {
	leftOfCursor := e.curLine()[:e.cursorX]
	rightOfCursor := e.curLine()[e.cursorX:]

	if e.isLastLine(e.cursorY) {
		e.lines[e.cursorY] = leftOfCursor
		e.lines = append(e.lines, rightOfCursor)
	} else {
		e.lines = append(e.lines[:e.cursorY+1], e.lines[e.cursorY:]...)
		e.lines[e.cursorY] = leftOfCursor
		e.lines[e.cursorY+1] = rightOfCursor
	}

	e.cursorX = 0
	e.cursorY++
}

func (e *editor) backspace() {
	if e.cursorX == 0 {
		if e.cursorY != 0 {
			deletedLine := e.removeLine(e.cursorY)

			e.lines[e.cursorY] = e.curLine() + deletedLine
		}
	} else {
		e.lines[e.cursorY] = e.curLine()[:e.cursorX-1] + e.curLine()[e.cursorX:]
		e.cursorX--
	}
}

func (e *editor) write(input rune) {
	if e.ForceLowercase {
		input = unicode.ToLower(input)
	}

	if e.cursorX == e.curLineLength() {
		e.lines[e.cursorY] = e.lines[e.cursorY] + string(input)
	} else {
		leftOfCursor := e.curLine()[:e.cursorX]
		rightOfCursor := e.curLine()[e.cursorX:]

		e.lines[e.cursorY] = leftOfCursor + string(input) + rightOfCursor
	}

	// Unicode characters can take up more than one byte. This at least shows this to the user but is not a whole solution
	for i := 0; i < len(string(input)); i++ {
		e.right()
	}
}

func (e *editor) removeLine(y int) string {
	removedLine := e.lines[y]
	e.lines = append(e.lines[:y], e.lines[y+1:]...)

	if y < e.cursorY || e.cursorY == e.numLines() {
		e.cursorY--
		e.cursorX = e.curLineLength()
		log.Printf("shifted cursor to %d %d", e.cursorX, e.cursorY)
	}

	return removedLine
}

func (e *editor) isLastLine(y int) bool {
	return y + 1 == e.numLines()
}

func (e *editor) numLines() int {
	return len(e.lines)
}

func (e *editor) numCharacters() int {
	numChars := 0

	for _, line := range e.lines {
		numChars += len(line)
	}

	return numChars
}

func (e *editor) curLine() string {
	return e.lines[e.cursorY]
}

func (e *editor) curLineLength() int {
	return e.lineLength(e.cursorY)
}

func (e *editor) lineLength(y int) int {
	return len(e.lines[y])
}

func (e *editor) empty() bool {
	return len(e.lines) == 1 && len(e.lines[0]) == 0
}

func (e *editor) lineFromLast(n int) string {
	return e.lines[e.numLines() - 1 - n]
}

// Compute what the cursor position will be if all lines are soft wrapped to fit within a certain number of columns
func (e *editor) getRealCursorPosition(offsetX, offsetY int) (cursorX int, cursorY int) {
	realX := (e.cursorX + offsetX) % e.NumCols
	realY := offsetY

	for y := 0; y < e.cursorY; y++ {
		realY += (e.lineLength(y) / e.NumCols) + 1
	}

	if ((offsetX + e.lineLength(0)) / e.NumCols) > (e.lineLength(0) / e.NumCols) {
		realY++
	}

	realY += e.cursorX / e.NumCols

	log.Printf("Computed real cursor postion of %d %d from virtual position of %d %d with offset %d %d", realX, realY, e.cursorX, e.cursorY, offsetX, offsetY)
	return realX, realY
}

func (e *editor) handleInput(input Key) {
	switch v := input.(type) {
	case RuneKey:
		if unicode.IsPrint(rune(v)) {
			e.write(rune(v))
		}
	case ControlKey:
		switch v {
		case ControlLeft: e.left()
		case ControlRight: e.right()
		case ControlUp: e.up()
		case ControlDown: e.down()
		case ControlEnter: e.newline()
		case ControlBackspace: e.backspace()
		case ControlSpace: e.write(' ')
		}
	}
}
