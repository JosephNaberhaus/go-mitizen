package prompt



type YesNo struct {
	base

	Description string

	editor *editor
}

func (y *YesNo) Show() error {
	y.base.Show()

	y.editor = &editor{NumCols: y.output.numCols}
	y.editor.Init()
	y.render()

	return loopUntilFinished(y)
}

func (y *YesNo) handleInput(input Key) {
	if y.finished {
		return
	}

	if input == ControlEnter {
		y.Finish()
		return
	} else {
		y.editor.handleInput(input)
	}

	y.render()
}

func (y *YesNo) render() {
	y.output.clear()

	y.output.writeColor("? ", ColorGreen)
	y.output.write(y.Description)
	y.output.writeColor(" (y/N) ", ColorGreen)

	offsetX := y.output.cursorX
	offsetY := y.output.cursorY

	if y.finished {
		if y.Response() {
			y.output.writeColorLn("Yes", ColorCyan)
		} else {
			y.output.writeColorLn("No", ColorCyan)
		}
	} else {
		y.output.writeLn(y.editor.curLine())
	}

	y.output.setCursor(y.editor.getRealCursorPosition(offsetX, offsetY))

	y.output.flush()
}

func (y *YesNo) Finish() {
	y.base.Finish()

	y.render()
	y.output.commit()
}

func (y *YesNo) Response() bool {
	return y.editor.curLineLength() > 0 && y.editor.curLine()[0] == 'y'
}
