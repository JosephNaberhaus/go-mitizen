package prompt

import (
	"github.com/JosephNaberhaus/go-mitizen/util"
	escapes "github.com/snugfox/ansi-escapes"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type output struct {
	numCols int

	cursorX int
	cursorY int

	numOutputLines int

	printBuffer *strings.Builder
}

func newOutput() *output {
	o := new(output)

	o.cursorX = 0
	o.cursorY = 0

	o.numOutputLines = 1

	dimensions, err := escapes.GetConsoleSize(os.Stdout.Fd())
	if err != nil {
		// Default to a standard sized terminal
		o.numCols = 80
	} else {
		o.numCols = dimensions.Cols
	}

	o.printBuffer = new(strings.Builder)

	return o
}

func (o *output) writeEscapeSequence(content string) {
	log.Printf("Writing escape sequence")
	o.printBuffer.WriteString(content)
}

func (o *output) write(content string) {
	runeCount := utf8.RuneCountInString(content)
	log.Printf("Writing %d characters with a rune count of %d", len(content), runeCount)

	o.printBuffer.WriteString(content)
	o.cursorX += runeCount
	o.wrapCursor()
}

func (o *output) writeColor(content string, color Color) {
	o.writeEscapeSequence(color.ToTextEscapes())
	o.write(content)
	o.writeEscapeSequence(ColorWhite.ToTextEscapes())
}

func (o *output) writeLn(content string) {
	o.write(content)
	o.nextLine()
}

func (o *output) writeColorLn(content string, color Color) {
	o.writeColor(content, color)
	o.nextLine()
}

// wraps internally stored position of the cursor when it goes beyond the column border so that it matches the real one
func (o *output) wrapCursor() {
	o.cursorY += (o.cursorX - 1) / o.numCols
	o.cursorX = (o.cursorX-1)%o.numCols + 1

	o.numOutputLines = util.Max(o.numOutputLines, o.cursorY+1)
}

func (o *output) moveCursor(dx, dy int) {
	if o.cursorX+dx < 0 {
		dx = -o.cursorX
	} else if o.cursorX+dx > o.numCols {
		dx = o.numCols - o.cursorX
	}

	if o.cursorY+dy < 0 {
		dy = -o.cursorY
	} else if o.cursorY+dy > o.numOutputLines {
		dy = o.numOutputLines - o.cursorY
	}

	if dx == 0 && dy == 0 {
		return
	}

	log.Printf("Shifting cursor: %d %d", dx, dy)

	o.cursorX += dx
	o.cursorY += dy

	o.printBuffer.WriteString(escapes.CursorMove(dx, dy))
}

func (o *output) setCursor(x, y int) {
	log.Printf("Moving cursor from: %d %d to: %d %d", o.cursorX, o.cursorY, x, y)
	if o.cursorX == x && o.cursorY == y {
		return
	}

	o.moveCursor(x-o.cursorX, y-o.cursorY)
}

func (o *output) hideCursor() {
	log.Println("Hiding the cursor")
	o.writeEscapeSequence(escapes.CursorHide)
}

func (o *output) showCursor() {
	log.Println("Showing the cursor")
	o.writeEscapeSequence(escapes.CursorShow)
}

func (o *output) nextLine() {
	log.Printf("Moving to next line with cursor at %d %d and %d output lines", o.cursorX, o.cursorY, o.numOutputLines)
	if o.cursorY+1 == o.numOutputLines {
		log.Println("Inserting a newline")
		o.printBuffer.WriteString("\n")
		o.numOutputLines++

		o.cursorY++
		o.cursorX = 0

		o.printBuffer.WriteRune(' ')
		o.cursorX++
		o.moveCursor(-1, 0)
	} else {
		o.setCursor(0, o.cursorY+1)
	}
}

func (o *output) clear() {
	log.Println("Clearing output")
	for y := 0; y < o.numOutputLines; y++ {
		o.setCursor(0, y)
		o.writeEscapeSequence(escapes.EraseLine)
	}
	o.setCursor(0, 0)
}

func (o *output) flush() {
	log.Println("Flushing output buffer")
	print(o.printBuffer.String())
	o.printBuffer.Reset()
}

func (o *output) commit() {
	log.Println("Committing output")

	if o.cursorX != 0 {
		o.nextLine()
	}

	o.flush()
}
